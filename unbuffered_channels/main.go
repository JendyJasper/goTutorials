package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func chanAllSend(ch chan<- int, msg []int) {
	//i used chan<- to create a send only function. It means the function can't receive
	//keep sending each value in the slice into the channel
	for _, i := range msg {
		ch <- i
	}
	//close the channel when you are done sending so that it won't enter deadlock
	close(ch)
	wg.Done()

}

func chanAllRcv(ch <-chan int) {
	//i used <-chan to create a receive only function. It means the function can't send
	//loop through the entire channel and print each item in the channel
	for i := range ch {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	//make the channel
	ch := make(chan int)
	msg := []int{34, 32, 12, 9, 23, 90, 206}
	wg.Add(2)
	go chanAllSend(ch, msg)
	go chanAllRcv(ch)

	wg.Wait()
}
