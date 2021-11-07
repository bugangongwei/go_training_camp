package week03

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

var (
	stop = make(chan struct{}, 2)
)

type appHandler struct {}

func (ah *appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Debug")
}

func serveApp(ctx context.Context) error {
	s:=http.Server{
		Addr: "localhost:8080",
		Handler: &appHandler{},
	}

	go func() {
		<-stop
		fmt.Println("shut down server: ", s.Addr)
		s.Shutdown(ctx)
	}()

	return s.ListenAndServe()
}

func serveDebug(ctx context.Context) error {
	s:=http.Server{
		Addr: "localhost:8899",
		Handler: http.NewServeMux(),
	}

	go func() {
		<-stop
		fmt.Println("shut down server: ", s.Addr)
		s.Shutdown(ctx)
	}()

	return s.ListenAndServe()
}

func watchSignal() error {
	quit := make(chan os.Signal, 1)
	signal.Reset(os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	fmt.Println("received signal: ", <-quit)

	for i:=0;i<len(stop);i++{
		fmt.Println(i)
		stop<- struct{}{}
	}
	close(stop)

	return nil
}

func LifeCycleManage() {
	ctx := context.Background()
	
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return serveDebug(egCtx)
	})

	eg.Go(func() error {
		return serveApp(egCtx)
	})
	
	eg.Go(func() error {
		return watchSignal()
	})

	err := eg.Wait()

	fmt.Println(err)

	if err != nil {
		fmt.Println("should quit while err happen: ", err)
	}
}
