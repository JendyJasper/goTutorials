package main

import (
	"fmt"
)

func main() {
	var dType bool = true
	fmt.Printf("%v %T\n", dType, dType)

	var (
		a int       = 10
		b int       = 3
		c int       = 13
		d complex64 = 3 + 2i
		e complex64 = 2i
	)

	fmt.Println(a * b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(float32(a) / float32(b))

	// bitshift
	fmt.Println(c << 3)
	fmt.Println(c >> 2)

	// complex numbers
	fmt.Println(d + e)
	fmt.Println(d - e)
	fmt.Println(d * e)
	fmt.Println(d / e)

	// pulling out the imaginary part and real part from the complex numbers
	fmt.Println(real(d))
	fmt.Println(real(e))
	fmt.Println(imag(d))
	fmt.Println(imag(e))

	// making a complex number
	fmt.Println(complex(3, 7))

	// byte collection
	f := "This is Jendy"
	g := []byte(f)
	fmt.Println(f)
	fmt.Println(g)

	c = 73
	fmt.Println(c << 2) //left bitshift
	fmt.Println(c >> 4) //right bitshift

}
