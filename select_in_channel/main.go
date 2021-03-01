package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sendCha1(cs chan int, s []int) {
	for _, i := range s {
		cs <- i
	}
	close(cs)
	wg.Done()
}

func sendCha2(cc chan string, c map[string]string) {
	for k, v := range c {
		cc <- k
		cc <- v
	}
	close(cc)
	wg.Done()
}

func main() {
	cs := make(chan int, 100)
	cc := make(chan string, 100)
	s := []int{12, 21, 76, 90, 43, 909}
	c := map[string]string{"Abia": "Umuahia", "Imo": "Owerri", "Ebonyi": "Abakiliki"}
	wg.Add(2)
	go sendCha1(cs, s)
	go sendCha2(cc, c)

	//use for loop to make it none blocking
	for i := 0; i < 100; i++ {
		select {
		case <-cs:
			for css := range cs {
				fmt.Println(css)
			}
		case <-cc:
			for ccc := range cc {
				fmt.Println(ccc)
			}
		//use the default statement to prevent it from exiting
		default:
			fmt.Println("Loading #>#>#>")
			//sleep it for a second to allow other go routines to process
			time.Sleep(1 * time.Second)
		}
	}

	wg.Wait()

}
