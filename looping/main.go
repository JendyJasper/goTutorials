package main

import (
	"fmt"
	//	"strings"
)

var spacer = " " //strings.Repeat(" ", 1)

func main() {
	//for loops
	//this is a very basic example
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//another way to declear the initialiser, expression and increamenter
	//observe the variable decleared above and a semic colon used
	// variables decleared outside of the for block becomes available to the entire function
	j := 0
	for ; j < 4; j++ {
		fmt.Println(j)
	}

	//yet another way to declear it. This is more like a do while loop in python
	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}

	//another way to increment the value of the variiable.
	//Uncomment any of the incrementing types to observer
	l := 0
	for l < 5 {
		fmt.Printf("Number is %v\n", l)
		//l += 1
		//l++
		l = l + 1
	}

	//you can declear more than one variable in a for statement
	for i, j := 0, 2; i < 5; i, j = i+1, j*2 {
		fmt.Println(i, j)
	}

	//you can use if statements inside a for block and nest it
	//expression to check even and odd numbers
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is Even number\n", i)
		} else {
			fmt.Printf("%v is an Odd Number\n", i)
		}
	}

	//you can also nest for statements inside for statements and use if statements too
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(i * j)
			if i*j >= 20 { //stop the calculation the moment the answer is more than or eqaul to 20 using the break statement
				break
			}
		}

	}

	//use the loop keyword called a label to break out of the entire for loop and not just the nested
loop:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(i * j)
			if i*j >= 15 { //stop the calculation the moment the answer is more than or eqaul to 20 using the break statement
				break loop
			}

		}
	}

	//you can use the continue keyword to break out of the current iteration
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Printf("%v is an odd number\n", i)
		}
	}

	//you can as well use for loop to iterate through collections like arrays, slices, maps, strings
	var testmap = map[string]string{
		"Ada":       "1st daughter",
		"Opara":     "1st son",
		"Olu":       "2nd daughter",
		"Olu Opara": "2nd son",
	}

	var testslice = []int{23, 56, 45, 89, 9, 23}

	var testarray = [4]float64{34.6, 56.98, 23.0902, 45.786}

	var teststring = "This is Jendy"

	//looping collections produces two results, the values and the keys or the index value and the value
	// you can deside to use the underscore _ keyword if you need to print just the values
	for k, v := range testmap {
		println(k, v)
	}
	for k := range testmap {
		println(k)
	}
	for _, v := range testmap { //printing just values
		println(v)
	}
	for k, v := range testslice {
		println(k, v)
	}
	for k, v := range testarray {
		println(k, v)
	}
	for k, v := range teststring {
		println(k, string(v))
	}

	//space out
	fmt.Println(spacer)

	//check for prime number of a given value
	isprime := 4
	primebool := true
	if isprime == 0 || isprime == 1 {
		fmt.Printf("%v is not a prime number\n", isprime)
	} else {
		for i := 2; i <= isprime/2; i++ {
			if isprime%i == 0 {
				fmt.Printf("%v is not prime\n", isprime)
				primebool = false
				break
			}

		}
		if primebool == true {
			fmt.Printf("%v is prime\n", isprime)
		}
	}

	//space out
	fmt.Println(spacer)

	//check for prime numbers in a range
	for m := 0; m <= 100; m++ {
		if m == 0 || m == 1 {
			fmt.Printf("%v is not prime\n", m)
		} else {
			n := m / 2
			isright := true
			for k := 2; k <= n; k++ {
				if m%k == 0 {
					fmt.Printf("%v is not prime\n", m)
					isright = false
					break
				}
			}
			if isright == true {
				fmt.Printf("%v is prime\n", m)
			}
		}
	}

	//space out
	fmt.Println(spacer)

	//check for how many prime numbers there are in a range, and also print out the prime numbers in a slice
	m := 0
	hi := 100
	lo := 0
	lenprime := []int{}
	for ; m <= hi; m++ {
		if m == 0 || m == 1 {
			//fmt.Printf("%v is not prime\n", m)
		} else {
			n := m / 2
			isright := true

			for k := 2; k <= n; k++ {
				if m%k == 0 {
					// fmt.Printf("%v is not prime\n", m)
					isright = false
					break
				}
			}
			if isright == true {
				//fmt.Printf("%v is prime\n", m)
				lenprime = append(lenprime, m)

			}
		}
	}
	fmt.Printf("There are %v prime numbers between %v and %v\n", len(lenprime), lo, hi)
	fmt.Printf("The %v prime numbers are %v\n", len(lenprime), lenprime)

	//space out
	fmt.Println(spacer)

	//calculate the number of digits in an integer that evenly divide it
	isInteger := 105
	isDivisor := 1
	isDigit := []int{}

	for ; isDivisor <= isInteger; isDivisor++ {
		if isInteger%isDivisor == 0 {
			//fmt.Printf("%v can divide %v\n", isDivisor, isInteger)
			isDigit = append(isDigit, isDivisor)
		}
	}
	fmt.Printf("All the didgits that can divide the integer %v are %v\n", isInteger, isDigit)
	fmt.Printf("They are %v in number\n", len(isDigit))
}
