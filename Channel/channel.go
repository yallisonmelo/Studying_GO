package main

import (
	"fmt"
	"time"
)

func task(done chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println("Task done!")
	done <- true

}

func sendMessages(c chan<- string) {
	c <- "Hello"
}

func receiveMessages(c chan string, done chan bool) {
	msg := <-c
	fmt.Println(msg)
	done <- true
}

func testChannel(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Printf("Write in channel:%d\n", i)
		time.Sleep(time.Millisecond * 90)
	}
	fmt.Println("Done Write!")
	fmt.Println("Close channel")
	close(c)
}

func main() {

	//define size channel in 3
	cn := make(chan int, 3)
	go testChannel(cn)

	for v := range cn {
		fmt.Printf("Read from channel: %d\n", v)
		time.Sleep(time.Millisecond * 190)
	}
	done := make(chan bool)
	go task(done)
	<-done

	msg := make(chan string)
	finished := make(chan bool)
	go sendMessages(msg)
	go receiveMessages(msg, finished)
	<-finished
	fmt.Println("Done!")

}
