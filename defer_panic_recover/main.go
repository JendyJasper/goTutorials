package main

import (
	"fmt"
	//"os"
)

func panicChecker() {
	if r := recover(); r != nil {
		fmt.Println("I just recovered from a panic named: ", r)
	}
}
func main() {
	//when you introduce a panick keyword, it won't run all other codes after the panick keyword
	//also observe that the defer statement printed before the panic
	//and the statement after the panic din't print
	//remove the comment from the line of code after the panick keyword to see what i mean
	//now panicchecker function was created to recover from the panic when there is indeed a panic
	//it is defered in the main function so that it will get printed only when
	//all other statemnts are printed and before the panic is hit
	//also note that the panicchecker function which has recover keyword will only print when it hits a panic
	//add a comment to the line of code with the panic keyword and observe that the panicchker function won't run
	fmt.Println(" ")
	defer panicChecker()
	fmt.Println("This will print print before the panic")
	panic("I don't want to print the statemnt below me")
	//fmt.Println("This will not print print before the panic")

}
