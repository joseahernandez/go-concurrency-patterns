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
	fmt.Println("Stopping generator after generate 3 numbers\n")

	stop := make(chan bool)
	c := generate(stop)

	for i := 0; i < 3; i++ {
		number := <-c
		fmt.Println(fmt.Sprintf("The number generated is %d", number))
	}

	// Sending stop message
	stop <- true

	// Waiting for answer in the stop channel
	fmt.Println("Generator has stopped", <-stop)
}

func generate(stop chan bool) chan int {
	c := make(chan int)

	go func() {
		for {
			select {
			case c <- rand.Intn(maxValueToGenerate):
				time.Sleep(time.Duration(rand.Intn(maxSleepTime)) * time.Second)

			case <-stop:
				fmt.Println("Stopping generator")
				// Sending again the stop message to report the main function
				stop <- true
				return
			}
		}
	}()

	return c
}
