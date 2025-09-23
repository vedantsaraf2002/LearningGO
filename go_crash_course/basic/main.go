package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hey Vedant")
	var num1 int = 3
	var num2 int = 0
	fmt.Println(intDivision(num1, num2))
	var result, remainder, err = intDivision(num1, num2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	s1, s2 := 1, 2
	fmt.Println(s1 + s2)
	const pi float32 = 3.1415 // value never changes that's why const.
	fmt.Println(pi)
	// Array
	var intArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Print(intArr)

	intArr2 := [...]int32{1, 2, 3, 4}
	fmt.Print(intArr2)

	// Slices -- array of undefined length
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Print(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	// Maps
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)

	var age, ok = myMap2["Adam"]
	/* default return by GO, returns true if value found else false,
	if value is not there still
	returns the default value of that datatype in case of int, 0 */

	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// Loops
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age:%v \n", name, age)
	}

}

// func intDivision(num1 int, num2 int) int {
// 	return num1 / num2
// }

// returning multiple values

func intDivision(num1 int, num2 int) (int, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return num1 / num2, num1 % num2, err
}
