# Go Concurrency Patterns

In this repo we can see some examples to work in Go with goroutines and channels.

Let's start with a function that generates numbers between 0 and `maxValueToGenerate` in a period of time between 0 and `maxSleepTime` seconds and send that values by a channel.


```go
const maxValueToGenerate = 10
const maxSleepTime = 3

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
```

## Waiting in a loop

The first example is the simplest, we generate a know number of values and wait in a loop until the channel returns that number of values.

```go
c := generate(3)

for i := 0; i < 3; i++ {
    number := <-c
    fmt.Println(fmt.Sprintf("The number generated is %d", number))
}
```

The file `generator.go` contains the full code to execute this example.

## Fan in

In this case we have two goroutines generating numbers each one in a different channel. Then we merge those numbers in another channel that is returned to the main function.

The `generate` function is the same that we saw before and the function that merge the result of the channels is the following:

```go
func merge(input1, input2 <- chan int) <- chan int {
    c := make(chan int)

    go func() { for { c <- <- input1 }}()
    go func() { for { c <- <- input2 }}()

    return c
}
```

All messages received from the `generate` function are listen by one of the goroutine in the `merge` function. Then these messages are send by the `c` channel that is the channel that is listen by the main function.

So the use of that `merge` function is as follows: 

```go
c := merge(generate(numbersToGenerate), generate(numbersToGenerate))

for i := 0; i < numbersToGenerate * 2; i++ {
    number := <-c
    fmt.Println(fmt.Sprintf("The number generated is %d", number))
}
```

The file `fan-in.go` has the code.

### Select statement

One improvement that we can do in the code of the `merge` function is use `select` instead the two goroutines listening the messages from `generate`. The `select` statement blocks until one of its cases can run.

```go
func merge(input1, input2 <-chan int) <-chan int {
    c := make(chan int)

    go func() {
        for {
            select {
            case s := <-input1:
                c <- s
            case s := <-input2:
                c <- s
            }
        }
    }()

    return c
}
```

File `fan-in-select.go` contains the code.

### Timeout

Inside the main function we can create a timeout to stop waiting messages if the time set in the timeout is exceeded.

```go
for {
    select {
    case number := <-c:
        fmt.Println(fmt.Sprintf("The number generated is %d", number))

    case <- time.After(2 * time.Second):
        fmt.Println("Too much time waiting. Stop here")
        return
    }
}
```

`fan-in-select-timeout.go` has the code.

## Stop a channel

In this case we have a `generate` function that generates numbers until it receives a message to stop. The function receives a channel `stop` and when receives a value stop generating numbers and end the goroutine.

```go
func generate(stop chan bool) chan int {
    c := make(chan int)

    go func() {
        for {
            select {
            case c <- rand.Intn(maxValueToGenerate):
                time.Sleep(time.Duration(rand.Intn(maxSleepTime)) * time.Second)

            case <- stop:
                fmt.Println("Stopping generator")
                return
            }
        }
    }()

    return c
}
```
In the main function we create a channel to send the stop message to `generate` function and after generate 3 numbers we send the message that makes the goroutine of `generate` stops.

```go
stop := make(chan bool)
c := generate(stop)

for i := 0; i < 3; i++ {
    number := <-c
        fmt.Println(fmt.Sprintf("The number generated is %d", number))
    }

stop <- true
```

File `topping-channel.go` contains de code.