package main

import (
	"fmt"
	"reflect"
)

func main() {
	// maps - first method of assigning variable when you are declearing and assigning values in the same line
	var StatesAndCapitals = map[string]string{ //the first data type explains the key and the other datatype explains the value
		"Imo":       "Owerri",
		"Abia":      "Umuahia",
		"Adamawa":   "Yola",
		"Akwa Ibom": "Uyo",
		"Anambara":  "Awka",
	}
	fmt.Println(StatesAndCapitals)
	//you can add more items to the existing map
	StatesAndCapitals["Nigeria"] = "Abuja"
	StatesAndCapitals["Ghana"] = "Accra"
	fmt.Println(StatesAndCapitals)
	//you can delete items from the map
	delete(StatesAndCapitals, "Abia")
	fmt.Println(StatesAndCapitals)
	//print values of keys that are not in the map
	fmt.Println(StatesAndCapitals["Taraba"]) //prints an empty map
	//check if a map exists first with the comma and ok. it prints false if it doesnt exist
	_, ok := StatesAndCapitals["Jos"]
	fmt.Println(ok)

	//Another method of assigning values to variables when everything is not decleared in the same line. Notice the Make Keyword
	var StatesAndCapital = make(map[string]string)
	StatesAndCapital["Enugu"] = "Enugu"
	StatesAndCapital["Imo"] = "Owerri"

	fmt.Println(StatesAndCapital)

	//structs
	type Country struct {
		NumOfStates int
		Capital     string
		President   string
	}

	// initialising the struct and assigning it to Nigeria
	var Nigeria = Country{
		NumOfStates: 36,
		Capital:     "Abuja",
		President:   "GMB Buhari",
	}
	fmt.Println(Nigeria)           //prints the entire contents
	fmt.Println(Nigeria.Capital)   //prints just the capital
	fmt.Println(Nigeria.President) //prints just the president

	var d Country      //used to initialise a struct and assign a value to it same time
	d.NumOfStates = 27 //use the initialised variable to assign values to the structs
	d.Capital = "Abu Dhabi"
	d.President = "I don't know"
	fmt.Println(d)
	fmt.Println(d.Capital)

	//embedding structs
	type Family struct {
		NumOfTypes int
		Breed      string
	}

	type Animal struct {
		Family //struct called
		Name   string
		Colour string
	}

	var hawk = Animal{}
	hawk.NumOfTypes = 5
	hawk.Breed = "Asian"
	hawk.Name = "Hawk"
	hawk.Colour = "Black"
	fmt.Println(hawk)
	fmt.Println(hawk.NumOfTypes)

	//Tags in structs
	type Car struct {
		Make string //`required max:"100"` //The additional is the tag
	}

	t := reflect.TypeOf(Car{})        //using the reflect function to get the tag value
	field, _ := t.FieldByName("Make") //to know the particuar tag for this field
	var c Car
	c.Make = "Toyota"
	fmt.Println(field.Tag)

}
