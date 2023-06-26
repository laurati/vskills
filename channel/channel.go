package main

import "fmt"

func main() {

	// exemplo 1
	c := make(chan bool)
	go waitAndSay(c, "world")
	fmt.Println("Hello")

	// we send a signal to c in order to allow waitAndSay to continue
	c <- true
	// we wait to receive another signal on c before we exit
	<-c

	// exemplo 2
	// channels
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// exemplo 3
	a := make(chan string)
	go sayHelloMultipleTimes(a, 5)
	for s := range a {
		fmt.Println(s)
	}

}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}

func sayHelloMultipleTimes(a chan string, n int) {
	for i := 1; i <= n; i++ {
		a <- "Hello"
	}
	close(a)
}
