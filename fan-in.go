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
	fmt.Println(fmt.Sprintf("2 generators will generate %d numbers each one\n", numbersToGenerate))

	c := merge(generate(numbersToGenerate), generate(numbersToGenerate))

	for i := 0; i < numbersToGenerate*2; i++ {
		number := <-c
		fmt.Println(fmt.Sprintf("The number generated is %d", number))
	}

	fmt.Println("\nNumbers have been generated. Bye!")
}

func merge(input1, input2 <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
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
