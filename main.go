package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func server(addr string, handler http.Handler, ctx context.Context) error {
	s := http.Server{Addr: addr, Handler: handler}

	go func() {
		<-ctx.Done()
		s.Shutdown(ctx)
	}()

	return s.ListenAndServe()
}

type Handler struct {
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return server("0.0.0.0:8080", &Handler{}, ctx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			return errors.New("通过信号关闭http服务")
		case <-ctx.Done():
			return errors.New("http服务关闭")
		}

	})

	err := g.Wait()
	fmt.Printf("关闭服务来源：%v", err)

}
