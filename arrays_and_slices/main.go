package main

import (
	"fmt"
)

func main() {
	//ARRAYS ARE DECLEARED EITHER WITH ELIPSES OR THE CAPACITY
	// first way to declear an array. Notice the size of the array was specified. colon used in place of var
	//this accepts fixed lenght
	classResults := [3]int{0, 2, 5}
	pri5 := [5]string{"Jendy", "Silas"}
	fmt.Println(classResults)
	fmt.Println(pri5)

	//another way to declear an array.  Notice the size is not decleared. elipses(...) used instead/. Var used
	//this accepts variable lenght
	var classAge = [...]int{11, 12, 15}
	pri6 := [...]string{"Jendy", "Silas", "Mellisa", "Julia", "Wisdom", "Iyke"}
	fmt.Println(classAge)
	fmt.Println(pri6)
	fmt.Println(cap(pri6)) //print the capacity
	fmt.Println(len(pri6)) //print the lenght

	//decleared the varriables of the array in different lines
	var classAddr [3]int
	var pri4 [2]string
	classAddr[0] = 324
	classAddr[1] = 297
	classAddr[2] = 123
	pri4[0] = "John"
	pri4[1] = "Kennedy"
	fmt.Println(classAddr)
	fmt.Println(pri4)

	//this is how to declear a mattrixed array
	var mattrArr [2][2]int
	mattrArr[0][0] = 2
	mattrArr[0][1] = 3
	mattrArr[1][0] = 7
	mattrArr[1][1] = 8
	fmt.Println(mattrArr)

	//SLICES ARE DECLEARED WITHOUT MAKING THE CAPACITY KNOWN
	// same line declearation and assignment
	var classCaptains = []string{"chin", "Nemere", "Jazzy", "Robby"}
	fmt.Println(classCaptains)

	// different lines decleararion and assignent. Use append to add more values
	var classTeachers = []string{}
	var classPrefects = []string{"Ekene", "Mavin"}
	classTeachers = append(classTeachers, "Silas")
	fmt.Println(classTeachers)
	classTeachers = append(classTeachers, "Chinedu")
	fmt.Println(classTeachers)
	classTeachers = append(classTeachers, "Silas", "David")
	fmt.Println(classTeachers)

	// use elipses at the end of the variable you are concatenating to add to slices
	classTeachers = append(classTeachers, classPrefects...)
	fmt.Println(classTeachers)
	classTeachers = append(classTeachers, []string{"Adaora", "Oke"}...)
	fmt.Println(classTeachers)

	// you can make a slice, specifying the capacity and lenght with one argument
	schoolNames := make([]string, 2) //3 is lenght and capacity when one argument is specified
	schoolNames = append(schoolNames, "Davis", "Royal College")
	fmt.Println(cap(schoolNames))
	fmt.Println(schoolNames)

	// you can make a slice, specifying the capacity and lenght
	schoolCourses := make([]string, 3, 20) //3 is lenght and 20 is capacity when two argument are specified
	schoolCourses = append(schoolCourses, "Davis", "Royal College", "Queens college")
	fmt.Println(len(schoolCourses))
	fmt.Println(schoolCourses)
	fmt.Println(cap(schoolCourses))

	// i can use index to crop some portions of the slice and assign it to a new variable
	schoolCourses2 := classTeachers[2:] //slice items from the 2nd index or 3rd element
	fmt.Println(schoolCourses2)

	// assign the value of a to b in a way that changes made to b will affect a
	a := [4]int{2, 5, 7, 9}
	b := &a // used to assign the value of b to the a location
	// a and b the same thing
	fmt.Println(a)
	fmt.Println(b)
	// 2nd index of b changed and it affected a on the next print
	b[2] = 34
	fmt.Println(b)
	fmt.Println(a)

	// assign the value of c to d in a way that changes made to d won't affect c
	c := [4]int{12, 25, 74, 19}
	d := c // value of c assigned to d
	// c and d the same thing
	fmt.Println(c)
	fmt.Println(d)
	// 2nd index of d changed and it didn't affect c on the next print
	d[2] = 340
	fmt.Println(d)
	fmt.Println(c)
}
