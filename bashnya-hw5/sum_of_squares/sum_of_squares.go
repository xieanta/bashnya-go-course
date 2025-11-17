package main

import (
	"fmt"
	"sync"
)

var sumOfSquare int
var mu sync.Mutex

func square(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	sumOfSquare += x * x
	mu.Unlock()

}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	seq := [5]int{2, 4, 6, 8, 10}
	for _, num := range seq {
		go square(num, wg)

	}
	wg.Wait()
	fmt.Println(sumOfSquare)

}
