package main

import "fmt"

//you pass in a parameter and its data type. do this for every passed in parameter
func sayHi(a string, b string) {
	fmt.Println(a, b)
}

//if they are same datatype, you can pass in the parameters first and declear the data type last
func add(c, d int) {
	fmt.Println(c + d)
}

//let's try working with pointers in functions
func sayGreetings(e, f *string) {
	//look at the parameter using pass by pointer instead of by value
	//therefore, the change on the value of the parameter in the main function wont affect it
	*f = "Nkechi"
	fmt.Println(*f)
	fmt.Println(*e, *f)

}

//variadic functions
//you can use elipses and the data type to make a function accept multiple parameters, printed as a slice
//it can have one parameter
func variadic(num ...int) {
	fmt.Println("The numbers are", num)
}

//you can also pass the nessage in to the parameters
func variadicMsg(msg string, num ...int) {
	fmt.Println(msg, num)
}

//return values in a function
//the int data type after the parameters are passed in determines what is expected in the return value
func returnValue(msg string, numbers ...int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}

//returning more than one value in functions
func mathsOperations(a, b int) (int, int, int) {
	addIt := a + b
	multiplyIt := a * b
	subtractIt := a - b
	return addIt, multiplyIt, subtractIt

}

//you can as well pass the variable of the return paraeters into the return data type declearaaton
func mathsOp(a, b int) (addIt, multiplyIt, subtractIt int) {
	addIt = a + b
	multiplyIt = a * b
	subtractIt = a - b
	//you don't need to mention the reurn variable again
	return

}

//returning an error. lets expand the mathsOp function in a way that it could throw error
func divErr(a, b float64) (c, d, e, f float64, err error) {
	c = a + b
	d = a * b
	e = a - b
	f = a / b
	if b == 0 {
		return c, d, e, f, fmt.Errorf("Cannot divide by 0")
	}
	return
}

func main() {
	sayHi("I am", "Jendy Jasper")
	add(78, 2)
	saying := "Hello"
	name := "Jendy"
	//because the second parameter is passing by pointer, the jendy passed to the name variable won't print,
	//instead it is the one passed in the function itself
	sayGreetings(&saying, &name)
	variadic(1, 5, 7, 9, 76)
	variadicMsg("The nubers are", 1, 4, 6, 9)
	//when returning a value, you will have to assign the function to a variable and return the function
	suemer := returnValue("The sum is ", 10, 20, 30, 40, 50)
	fmt.Println(suemer)
	//returning multiple values. Assign number of variables that eqauls the number of return values
	plusIt, timesIt, subIt := mathsOperations(4, 5)
	fmt.Println(plusIt)
	fmt.Println(timesIt)
	fmt.Println(subIt)
	//still returning multiple statements but with the return statements cleaner
	mathsPIt, mathsTIt, mathsSIt := mathsOp(4, 6)
	fmt.Println(mathsPIt)
	fmt.Println(mathsTIt)
	fmt.Println(mathsSIt)
	//returning error
	fmt.Println(" ")
	q, r, s, t, err := divErr(15, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)

}
