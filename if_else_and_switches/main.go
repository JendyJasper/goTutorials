package main

import (
	"fmt"
)

func main() {
	//if and else statement. It will print the statement within the if block is a is Jahdi
	//and print the statement within the else block if a is not Jahdi
	a := "Jahdi"
	if a == "Jahdi" {
		fmt.Println("It is Jahdi")
	} else {
		fmt.Println("It is not Jahdi")
	}

	//else if statement. It will print the statement within the if block is a is Jahdi
	//and print the statement within the first else if block if a is Olivia and print the
	//statement within the next else if block if it is Somtoo and print the statement
	//within the else block if it is neither of them
	d := "Jendy"
	if d == "Jahdi" {
		fmt.Println("It is Jahdi")
	} else if d == "Olivia" {
		fmt.Println("It is Olivia")
	} else if d == "Somtoo" {
		fmt.Println("This is Somtoo")
	} else {
		fmt.Println("It is neither Jahdi, Olivia nor Somtoo")
	}

	//you can also do comparison operations with the if else statemnts
	b := "Oke"
	c := "Adau"
	//it prints the if block if either of Jendy, Joy or Iyke s the value of b.
	//it prints the else if block if both the value of b is Oke and value of c is Ada
	//it prints the else block if none of the statements are true
	if b == "Jendy" || b == "Joy" || b == "Iyke" {
		fmt.Println("b is either Jendy, Joy or Iyke")
	} else if b == "Oke" && c == "Ada" {
		fmt.Println("c is either Oke or Ada")
	} else {
		fmt.Println("It is none of the above")
	}

	//Switches and Cases
	//this functions like an if else statement, where the defaults keyword is equivalent of the else statement
	//so any of the cases gets printed depending on the value passed to the switch, and
	//if none is passed, the default value gets printed
	switch 2 {
	case 2:
		fmt.Println("This is case 2")
		//the fallthrough keyword makes it move to the next switch case no matter the result
		//remove the comment from any of the fallthrough keywords and observe what happens
	//	fallthrough
	case 4:
		fmt.Println("This is case 4")
	//	fallthrough
	case 9:
		fmt.Println("This is case 9")
	//	fallthrough
	default:
		fmt.Println("None of the above")
	}

	//the values passed to the switch keywords are called tags.they can be variables or actual values of any data type
	//however, the switch and the case must be of same data type
	//you can also make tagless switches

	e := 500
	f := 5000
	switch {
	case e < f:
		fmt.Println("e is less than f")
	case e > f:
		fmt.Println("e is greater than f")
	case e == f:
		fmt.Println("e is eqaul to f")
	}

	//if you want to compare cases of different data types from the switch, use the interface keyword
	var g interface{} = true
	switch g.(type) {
	case int:
		fmt.Println("g is int")
	case float64:
		fmt.Println("g is float64")
	case string:
		fmt.Println("g is string")
	case [3]int:
		fmt.Println("g is an array of 3 capacity with int")
	case [10]int:
		fmt.Println("g is an array of 10 capacity with int")
	case []string:
		fmt.Println("g is a slice of string")
	case map[string]string:
		fmt.Println("g is a map")
	case bool:
		fmt.Println("g is a bool")
	// case struct:
	// 	fmt.Println("g is a struct")
	default:
		fmt.Println("None of the above")

	}

	//you can check multiple values of a case against a switch
	switch 9 {
	case 12, 3, 5, 7:
		fmt.Println("It is either 12, 3, 5 or 7")
	case 1, 4, 9:
		fmt.Println("This is either 1, 4 or 9")
	default:
		fmt.Println("None of the above")

	}
}
