package main

import (
	"fmt"
	"strconv"
)

var (
	ada    string  = "a girl"
	obi    string  = "a boy"
	age    int     = 32
	height float32 = 6.7
)

func main() {

	fmt.Printf("%s and %s are married. %s is %v years old and %v feet in heigh\n", ada, obi, obi, age, height)

	ada := "a woman"
	fmt.Printf("%s and %s are married. %s is %v years old and %v feet in heigh\n", ada, obi, obi, age, height)
	fmt.Printf("%v %T\n", height, height)
	age = int(height)
	fmt.Printf("%v %T\n", age, age)

	var amount string
	amount = strconv.Itoa(age)
	fmt.Printf("%v %T\n", amount, amount)

	var money int
	money, _ = strconv.Atoi(amount)
	fmt.Printf("%v %T\n", money, money)

}
