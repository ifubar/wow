package server

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ifubar/wow/entities"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type GetTaskResponse struct {
	Token string
	Task  entities.Task
}

func (s *WisdomServer) getTask(w http.ResponseWriter, req *http.Request) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task := entities.NewTask()
	claims := Claims{
		Task: task,
		IP:   ip,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	raw, err := json.Marshal(GetTaskResponse{
		Token: token,
		Task:  task,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("task issued: %+v\n", task)
	w.Write(raw)
}

func (s *WisdomServer) getWisdom(w http.ResponseWriter, req *http.Request) {
	solution, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	task, ok := entities.Task{}.FromCtx(req.Context())
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !task.SolutionValid(solution) {
		log.Printf("solution invalid, task: %+v, solution %s\n", task, solution)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("solution invalid"))
		return
	}
	err = s.taskStorage.MarkSolved(task, expire)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("task solved, task: %+v, solution %s\n", task, solution)
	w.Write([]byte(s.wisdomStorage.Get()))
}
