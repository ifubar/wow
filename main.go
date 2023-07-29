package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ifubar/wow/server"
	"github.com/ifubar/wow/storage"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()

	wServer := server.NewServer("0.0.0.0:4444", storage.NewWisdom(), storage.NewTask())
	if err := wServer.Serve(ctx); err != nil {
		log.Default().Fatal(err)
	}
}
