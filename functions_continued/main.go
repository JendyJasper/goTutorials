package main

import "fmt"

//functions as types/variables
var subtract = func(a, b float64) {
	fmt.Println(a - b)
}

//working with methods
//first, decleare a struct
type greeter struct {
	greeting string
	name     string
}

//secondly, create a function and use the struct name as data type
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

//methods with pointers
type obodo struct {
	state   string
	country string
}

//the ponter reference added
func (g *obodo) ebe() {
	g.state = "Imo State"
	g.country = "It is in Nigeria"
	fmt.Println(g.state, g.country)
}

func main() {
	//anonymous functions
	//these are invoked at the end of the function
	func() {
		fmt.Println("I am the first anonymous function")
	}()

	//it still accepts parameters and arguments
	func(i string) {
		fmt.Println(i)
	}("Being anonymous, i can still accept a parameter and arguments")

	//functions as variables
	adder := func(a, b float64) {
		fmt.Println(a + b)
	}
	adder(3, 7)

	var minus func() = func() {
		fmt.Println("Hello")
	}
	minus()

	subtract(9, 3)

	//you can also decide and pass in only the datatypes and later assign the parameters
	var printer func(string, float64) float64
	printer = func(a string, b float64) float64 {
		fmt.Println(a, b)
		return b
	}
	ch := printer("Hello", 65)

	fmt.Println(ch)

	//now initialise the struct
	sayIt := greeter{
		greeting: "How are you,",
		name:     "Jendy Jasper",
	}

	//call on the method
	sayIt.greet()

	//working with pointers in methods. Observe how the name printed will be different
	mama := obodo{
		state:   "I am from Umuolulo Atta,",
		country: "It is in Imo state",
	}
	//none of the values decleared above will work, instead it is the one declread inside the method as it is a pointer
	mama.ebe()
}
