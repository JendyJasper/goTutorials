package main

import (
	"fmt"
	"sync"
)

//instantiate the waitgroup
var wg sync.WaitGroup

//when two functions share same resource like a variable, it's neccessary to lock the resource when one goroutine is using it
//it is acheived using mutex(mutually exclusive) it is in the sync package
var mt sync.RWMutex

var balance float64

func init() {
	balance = 0
}
func deposit(amount float64, wg *sync.WaitGroup, mt *sync.RWMutex) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		//use the lock method to loc the variable untill the go routine is done
		mt.Lock()
		balance = balance + amount
		fmt.Printf("%v successfully deposited. Balance = %v\n", amount, balance)
		mt.Unlock()
	}
	return
}

func withdraw(amount float64, wg *sync.WaitGroup, mt *sync.RWMutex) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		//use the lock method to loc the variable untill the go routine is done
		mt.Lock()
		balance = balance - amount
		fmt.Printf("%v successfully withdrawn. Balance = %v\n", amount, balance)
		mt.Unlock()
	}
	return
}

func balanceChecker(wg *sync.WaitGroup, mt *sync.RWMutex) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		mt.RLock()
		fmt.Printf("Balance = %v\n", balance)
		mt.RUnlock()
	}
}

func sayHello(name string, wg *sync.WaitGroup) {
	//defer the done method on the waitgroup. Done decreases the number of goroutines
	defer wg.Done()
	fmt.Println("Hello,", name)
}

func sayAge(age int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("My age is, ", age)
}

func main() {
	//sayHello("Jendy", &wg)
	//to make the function a goroutine, add the keyword, go, in front of the function
	//go sayHello("Jasper")

	//you will observe that the function above with the go keyword won't run before the main function exits.
	//you can use waitgroups to block the main function untill the goroutines is completed
	//use Add method to decleare the number of go routines you are running
	wg.Add(7)
	go sayHello("Nkechi", &wg)
	go sayAge(34, &wg)
	go sayAge(67, &wg)
	go sayHello("Iyke", &wg)
	go deposit(500, &wg, &mt)
	go withdraw(200, &wg, &mt)
	go balanceChecker(&wg, &mt)
	wg.Wait()
	//balanceChecker()
	fmt.Println("End of Execution")
}
