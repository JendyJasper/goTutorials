package main

import "fmt"

func main() {
	//you use the pointers (*) beside a type(int in this case) to assing the address of 'a' to the value of 'b'
	//you use the refernce (&) to get the address value
	var a int = 56
	var b *int = &a
	//here now, a will print 56 and b will print the address value
	fmt.Println(a, b)
	//you can use the deference symbol(*) which is beside the variable name to print the actual value in the address
	fmt.Println(a, *b)
	//so essentially, an * beside a type means it points to an address.
	//an * beside a variable dereferences the value and gives you the value
	//an & beside a variable prints the address
	//when you use pointers, you can change the value assigned to a variable and it affects the original value
	*b = 7
	fmt.Println(a, *b)
	a = 29
	fmt.Println(a, *b)

	//another way to declear the vareiable is same line declearaton
	c := 87
	d := &c
	fmt.Println(c, d)
	fmt.Println(c, *d)

	//using pointers in a struct
	type myStruct struct {
		foo string
	}

	var ms *myStruct
	//the empty value of a struct is nil
	fmt.Println(ms)
	ms = &myStruct{foo: "Jendy"}
	fmt.Println(*ms)    //to print the actual value
	fmt.Println(ms)     // to print it and indicating that it is a pointer
	fmt.Println(ms.foo) // printing just the value oof foo, which go sees as (*ms).foo

	//the empty value of a struct inialised with a new keyword is empty
	type myStruct2 struct {
		bar string
	}
	var num = new(myStruct2)
	fmt.Println(num)

	//slices and maps behave same way in the sense that a variable assigned to a map or slice
	//actually points to the memory address instead of the underllying value, hence
	//when the value is changed in one, it affects the other
	j := map[string]string{
		"Omo": "Ohyea",
		"Ozo": "BBN",
	}
	v := j
	fmt.Println(j, v)
	//now change any of either  or v and see how it affects both
	v["Anyi"] = "Enete"
	fmt.Println(j, v)
	v["Skul"] = "Boys"
	fmt.Println(j, v)

	//same applies to a slice
	n := []int{56, 98, 43}
	q := n
	fmt.Println(n, q)
	q[1] = 90
	fmt.Println(n, q)
	n[0] = 12
	fmt.Println(n, q)
	fmt.Println(" ")

	//but this is not the case for an array, as arrays copies just the underlying value
	wa := [3]int{56, 98, 43}
	qa := wa
	fmt.Println(wa, qa)
	wa[1] = 90
	fmt.Println(wa, qa)
	qa[0] = 12
	fmt.Println(wa, qa)

}
