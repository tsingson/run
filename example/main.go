package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tsingson/run"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var g run.Group
	{
		ctx, cancel := context.WithCancel(ctx) // note: shadowed
		g.Add(func() error {
			return runUntilCanceled(ctx)
		}, func(error) {
			cancel()
		})
	}

	go func() {
		time.Sleep(15 * time.Second)
		cancel()
		return
	}()
	g.Run()

}

func runUntilCanceled(ctx context.Context) error {
	fmt.Println("cron job goroutine in running *****************************", time.Now())
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("cron job goroutine in running *****************************", time.Now())
			time.Sleep(1 * time.Second)
		}

	}

}
