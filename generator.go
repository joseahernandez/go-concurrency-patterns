package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numbersToGenerate = 5
const maxValueToGenerate = 10
const maxSleepTime = 3

func main() {
	fmt.Println(fmt.Sprintf("%d numbers will be generate\n", numbersToGenerate))

	c := generate(numbersToGenerate)

	for i := 0; i < numbersToGenerate; i++ {
		number := <-c
		fmt.Println(fmt.Sprintf("The number generated is %d", number))
	}

	fmt.Println("\nNumbers have been generated. Bye!")
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
