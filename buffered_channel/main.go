package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sendCh(ch chan int) {
	ch <- 45
	ch <- 21
	ch <- 12
	close(ch) // close the sending channel so that the receive function will now when to stop receiving
	wg.Done()
}

func recvCh(ch chan int) {
	// fmt.Println(<-ch) //the below 3 print statements will work fine because there are just 3 items sent to the channel
	// fmt.Println(<-ch) receive more than 3 items and it will panick
	// fmt.Println(<-ch) you can use close function to close the send channel
	//a smarter way to print them all
	for i := range ch {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	//ch := make(chan int)//this will panic with deadlock. Make it a buffered channel  by passing
	//in another parameter to stop the panic, even tho it eon't print all the messages
	ch := make(chan int, 50)
	wg.Add(2)
	go sendCh(ch)
	go recvCh(ch)
	wg.Wait()
}
