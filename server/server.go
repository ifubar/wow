package server

import (
	"context"
	"net/http"
	"time"

	"github.com/ifubar/wow/entities"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

type WisdomStorage interface {
	Get() string
}

type TaskStorage interface {
	MarkSolved(task entities.Task, duration time.Duration) error
}

type WisdomServer struct {
	http          *http.Server
	wisdomStorage WisdomStorage
	taskStorage   TaskStorage
}

func NewServer(addr string, wisdom WisdomStorage, task TaskStorage) *WisdomServer {
	server := &WisdomServer{
		http:          nil,
		wisdomStorage: wisdom,
		taskStorage:   task,
	}

	rateLimiterMiddleware := stdlib.NewMiddleware(limiter.New(memory.NewStore(), limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  100,
	}))

	mux := http.NewServeMux()
	mux.Handle("/task", rateLimiterMiddleware.Handler(http.HandlerFunc(server.getTask)))
	mux.Handle("/wisdom", rateLimiterMiddleware.Handler(checkJWT(server.getWisdom)))

	server.http = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return server
}

func (s *WisdomServer) Serve(ctx context.Context) error {
	done := make(chan error, 1)
	go func() {
		done <- s.http.ListenAndServe()
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		// http serv has 3s to shut down
		shCtx, shCancel := context.WithTimeout(context.Background(), time.Second*3)
		defer shCancel()
		return s.http.Shutdown(shCtx)
	}
}
