package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ch chan string, ctx context.Context, numOfWorker int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case res, ok := <-ch:
			if !ok {
				fmt.Printf("Воркер-%d: канал закрыт\n", numOfWorker)
				return
			}
			fmt.Printf("Привет из воркера %d: %s\n", numOfWorker, res)
		case <-ctx.Done():
			fmt.Printf("Воркер-%d завершен\n", numOfWorker)
			return
		}
	}

}

func main() {
	var workersCount = flag.Int("w", 1, "nums of workers")
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(*workersCount)

	dataChan := make(chan string, *workersCount)

	for i := 0; i < *workersCount; i++ {
		go worker(dataChan, ctx, i, &wg)
	}
	go func() {
		for i := 0; ; i++ {
			dataChan <- fmt.Sprintf("data-%d", i)
			time.Sleep(time.Second)
		}
	}()

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt)
	<-signChan
	cancel()
	close(dataChan)
	wg.Wait()

}
