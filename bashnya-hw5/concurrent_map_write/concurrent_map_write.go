package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

type MagicBallAnswer struct {
	ID      int
	Details string
}

var mu sync.Mutex

func worker(ch chan MagicBallAnswer, mapData *map[int]string, ctx context.Context, numOfWorker int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case res, ok := <-ch:
			if !ok {
				fmt.Printf("Воркер-%d: канал закрыт\n", numOfWorker)
				return
			}
			mu.Lock()
			(*mapData)[res.ID] = res.Details
			mu.Unlock()
			size := len(*mapData)
			fmt.Printf("Привет из воркера %d\nТекущий размер map: %d записей\n", numOfWorker, size)
		case <-ctx.Done():
			fmt.Printf("Воркер-%d завершен\n", numOfWorker)
			return
		}
	}
}
func main() {
	var workersCount = flag.Int("w", 1, "nums of workers")
	flag.Parse()
	dataChan := make(chan MagicBallAnswer, *workersCount)
	mapData := make(map[int]string)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(*workersCount)
	for i := 0; i < *workersCount; i++ {
		go worker(dataChan, &mapData, ctx, i, &wg)
	}

	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	go func() {
		for i := 0; ; i++ {
			answer := answers[rand.Intn(len(answers))]
			dataChan <- MagicBallAnswer{
				ID:      i,
				Details: answer,
			}
			time.Sleep(time.Second)
		}
	}()

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt)
	<-signChan
	fmt.Printf("Итоговый словарь: %v\n", mapData)
	cancel()
	close(dataChan)
	wg.Wait()

}
