package main

import (
	"fmt"
	//"os"
)

func main() {
	defer fmt.Println("This will be the last defer statemnt to print") //it will print last because of the defer keyword, even the it is the first statement
	fmt.Println("This will print first")
	fmt.Println("This will print second")

	fmt.Println(" ")
	//defer uses a last in first out approach, hence the last defer call is the first to print
	fmt.Println("This will print third")
	defer fmt.Println("This will be the fitst defer statement to print")
	fmt.Println("This will print  fourth")

	//differ doesn't change the arguments value even tho it is printed last
}
