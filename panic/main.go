package main

import (
	"fmt"
	//"os"
)

func main() {
	//when you introduce a panick keyword, it won't run all other codes after the panick keyword
	//also observe that the defer statement printed before the panic and the statement after the panic din't print
	//remove the comment from the line of code after the panick keyword to see what i mean
	fmt.Println(" ")
	defer fmt.Println("Still print this defer statement before hitting the panic")
	fmt.Println("This will print print before the panic")
	panic("I don't want to print the statemnt below me")
	//fmt.Println("This will not print print before the panic")

}
