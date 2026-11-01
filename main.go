package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

type Job struct {
	ID string `json:"id"`
}

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Minute * 5)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg := sync.WaitGroup{}

	cores := runtime.NumCPU()

	jobQueue := make(chan Job, cores*50)

	for i := 0; i <= cores; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			log.Printf("worker %d started", id)

			for {
				select {
				case <-ctx.Done():
					log.Printf("worker %d stopping", id)
					return
				case job, ok := <-jobQueue:
					if !ok || ctx.Err() != nil {
						return
					}
					time.Sleep(time.Second)
					log.Printf("worker %d processing %v", id, job.ID)
					continue
				}
			}
		}(i)
	}

	for i := 0; i <= 40; i++ {
		jobQueue <- Job{
			ID: fmt.Sprintf("job-%d", i),
		}
	}

	wg.Wait()

	log.Print("gorountines stopped")
}
