package main

import "fmt"

//Geometry interface
type Geometry interface {
	Area() float64
	Perim() float64
	Volume() float64
}

//Rectangle shape struct which defines the 3 sides of the shape
type Rectangle struct {
	lenght float64
	width  float64
	height float64
}

//Square shape struct which defines just one side since they are of equal sizes
type Square struct {
	width float64
}

//Area function which checks the Area of the rectangle
func (r Rectangle) Area() float64 {
	area := r.lenght * r.width
	return area
}

//Perim function which checks the perimeter of a rectangle
func (r Rectangle) Perim() float64 {
	perim := r.lenght + r.lenght + r.width + r.width
	return perim
}

//Volume function which checks the volume of a rectangle
func (r Rectangle) Volume() float64 {
	volume := r.lenght * r.height * r.width
	return volume
}

//Area of a square
//You can use pointer receiver as well
func (s *Square) Area() float64 {
	area := (*s).width * (*s).width
	return area
}

//Perim of a square
func (s *Square) Perim() float64 {
	perim := s.width + s.width + s.width + s.width
	return perim
}

//Volume of a square
func (s *Square) Volume() float64 {
	volume := s.width * s.width * s.width
	return volume
}

//Measure function which checks for the different measurements
func Measure(g map[string]Geometry) {
	for i, shape := range g {
		fmt.Printf("The Area of %s is %v\n", i, shape.Area())
		fmt.Printf("The Perimeter of %s is %v\n", i, shape.Perim())
		fmt.Printf("The Volume of %s is %v\n", i, shape.Volume())
	}

}

//you can also use empty interface in a function like
func hey(i interface{}) {
	fmt.Println(i)
}

//you can as well use interfaces in a struct
type anotherStruct struct {
	name interface{}
}

func main() {
	var rect = Rectangle{lenght: 45, width: 20, height: 14}
	var squ = Square{width: 23}
	var rect2 = Rectangle{lenght: 2, width: 4, height: 7}
	var shape2 = Square{width: 4}
	var rect3 = Rectangle{lenght: 2, width: 3.4, height: 2.9}
	var square2 = Square{width: 7.9}

	//pass in the pointer value to the slice on the square shape since it is a pointer receiver
	allShapes := map[string]Geometry{
		"rect":    rect,
		"squ":     &squ,
		"rect2":   rect2,
		"shape2":  &shape2,
		"rect3":   rect3,
		"square2": &square2,
	}
	Measure(allShapes)

	//another way of using the interface
	var card Geometry
	hello := Square{width: 2}
	card = &hello
	fmt.Println(card.Area())

	//yet another way to use the interface
	var ball Geometry = &(Square{width: 5})
	fmt.Println(ball.Area())

	//another way too
	var pen Geometry
	pen = &(Square{width: 4})
	fmt.Println(pen.Area())

	//empty interfaces. notice that it implements any type assigned to it
	var noType interface{}
	noType = "Jendy"
	fmt.Printf("%v is of type %T\n", noType, noType)
	noType = 34
	fmt.Printf("%v is of type %T\n", noType, noType)
	noType = 78.6
	fmt.Printf("%v is of type %T\n", noType, noType)
	noType = []int{23, 78}
	fmt.Printf("%v is of type %T\n", noType, noType)

	fmt.Println("")
	//pass in different types to the hey fuunction and observe that it works
	hey("Jendy")
	hey(35)
	hey(56.8)

	fmt.Println("")
	//pass in different types to the struct and observe how it works
	firstOne := anotherStruct{name: "Jendy"}
	fmt.Println(firstOne)
	secondOne := anotherStruct{name: 34}
	fmt.Println(secondOne)
	thirdOne := anotherStruct{name: 45.89}
	fmt.Println(thirdOne)

	//type assertions
	//you can check the type the interface implements
	var TypeTwo interface{}
	TypeTwo = 67.8

	fmt.Println("")
	checkType := TypeTwo.(float64)
	fmt.Println(checkType)

	//when you pass in a value that is not of float54 type, the program is gonna panic
	//you can use the comma ok syntax to handle the panic
	var TypeThree interface{}
	TypeThree = "I will have to ocme over"
	//checkTypeAgain := TypeThree.(int)
	//the second variable, ok returns a bolean value.
	if checkTypeAgain, ok := TypeThree.(string); ok {
		fmt.Println(checkTypeAgain, ok)
	} else {
		fmt.Println("It is not a string")
	}
}
