package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch:   make(chan string, 10),
		stop: make(chan struct{}),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t Tracker) Run() {
	for data := range t.ch {
		time.Sleep(time.Second)
		log.Println(data)
	}
	fmt.Printf("%s\n", "data chan closed")
	t.stop <- struct{}{}
}

func (t Tracker) Shutdown(ctx context.Context) {
	fmt.Printf("%s\n", "close data chan")
	close(t.ch)
	select {
	case <-t.stop:
		fmt.Printf("%s\n", "stopped")
	case <-ctx.Done():
		fmt.Printf("%s\n", "timeout")
	}
}

func main() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "test1")
	_ = tr.Event(context.Background(), "test2")
	_ = tr.Event(context.Background(), "test3")
	time.Sleep(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}
