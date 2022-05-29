package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func hello(w http.ResponseWriter, r *http.Request) {
	text := fmt.Sprintf("%s %s hello", r.Host, r.Method)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(text))
}

func ping(w http.ResponseWriter, r *http.Request) {
	text := fmt.Sprintf("%s %s pong", r.Host, r.Method)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(text))
}

func main() {
	rootCtx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(rootCtx)

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case s := <-signalChan:
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					log.Printf("Received signal: %v, shutdown ...", s)
					cancel()
				default:
					log.Println("Undefined signal:", s)
				}
			}
		}
	})

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/hello", hello)
	server1 := &http.Server{
		Addr: ":8080",
		Handler: mux1,
	}

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/ping", ping)
	server2 := &http.Server{
		Addr: ":8081",
		Handler: mux2,
	}

	eg.Go(func() error {
		log.Println("start server1")
		return server1.ListenAndServe()
	})

	eg.Go(func() error {
		<- ctx.Done()
		log.Println("close server1")
		return server1.Shutdown(ctx)
	})

	eg.Go(func() error {
		log.Println("start server2")
		return server2.ListenAndServe()
	})

	eg.Go(func() error {
		<- ctx.Done()
		log.Println("close server2")
		return server2.Shutdown(ctx)
	})

	if err := eg.Wait(); err != nil {
		log.Println("group error: ", err)
	}
	log.Println("all group done!")
}
