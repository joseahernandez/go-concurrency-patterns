package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numbersToGenerate = 5
const maxValueToGenerate = 10
const maxSleepTime = 2

func main() {
	fmt.Println("Generating numbers until ")

	c := merge(generate(numbersToGenerate), generate(numbersToGenerate))

	for {
		select {
		case number := <-c:
			fmt.Println(fmt.Sprintf("The number generated is %d", number))

		case <- time.After(2 * time.Second):
			fmt.Println("Too much time waiting. Stop here")
			return
		}
	}
}

func generate(numbersToGenerate int) chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < numbersToGenerate; i++ {
			c <- rand.Intn(maxValueToGenerate)
			time.Sleep(time.Duration(rand.Intn(maxSleepTime)) * time.Second)
		}
	}()

	return c
}
