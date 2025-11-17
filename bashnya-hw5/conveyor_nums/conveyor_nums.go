package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbersStr = flag.String("numbers", "1,2,3,4,5", "numbers for pipeline (comma-separated)")
	flag.Parse()
	arrayOfStrings := strings.Split(*numbersStr, ",")
	arrayOfNumbers := make([]int, 0, len(arrayOfStrings))
	for _, numStr := range arrayOfStrings {
		cleanedNum := strings.TrimSpace(numStr)
		numInt, err := strconv.Atoi(cleanedNum)
		if err != nil {
			fmt.Println("Ошибка преобразования в число")
			continue
		}
		arrayOfNumbers = append(arrayOfNumbers, numInt)

	}
	in := make(chan int, len(arrayOfNumbers))
	out := make(chan int, len(arrayOfNumbers))
	go func() {
		for _, value := range arrayOfNumbers {
			in <- value
			fmt.Printf("Число %d отправлено в канал in\n", value)
		}
		close(in)

	}()
	go func() {
		for {
			num, ok := <-in
			if !ok {
				fmt.Println("Канал in закрыт")
				close(out)
				fmt.Println("Канал out закрыт")
				return
			}
			out <- num * 2
			fmt.Printf("Число %d записано в канал out \n", num*2)

		}

	}()

	for result := range out {
		fmt.Printf("Число %d получено из канала out\n", result)
	}
}
