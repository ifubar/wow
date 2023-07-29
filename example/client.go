package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/ifubar/wow/client"
	"github.com/ifubar/wow/server"
)

func main() {
	cl := client.NewClient("http://server:4444")

	var err error
	var taskResp server.GetTaskResponse

	// wait for server start
	err = retrier.New(retrier.ConstantBackoff(10, time.Second), nil).Run(func() error {
		taskResp, err = cl.GetTask()
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	var solution string
	for {
		i++
		solution = taskResp.Task.InputPrefix + fmt.Sprintf(":try:%d", i)
		if taskResp.Task.SolutionValid([]byte(solution)) {
			break
		}
	}
	log.Printf("try=%d, solution=%s", i, solution)
	wisdom, err := cl.GetWisdom(solution, taskResp.Token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wisdom=%s", wisdom)
}
