package main

import "fmt"

//the value of a constant can't be changed, that is it is immutable
//however it can be shadowed by assigning a new value or type of same variable
//inside a block, different from the one at the package level.
const (

	// value of iota starts from 0 and auto incremeants, that is 0,1
	_  = iota // ignore first value by assigning blank identifer
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 4000000000. //you can add a period(.) at the end of untyped value to make it a float
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
