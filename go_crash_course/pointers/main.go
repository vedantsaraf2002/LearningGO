package main

import "fmt"

func main() {
	fmt.Println("Hey there, Vedant!")
	// Pointers
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p point to is: %v", *p)
	fmt.Printf("\nThe value if i is %v", i)
	*p = 10
	p = &i // now p refs to the memory address of i, giving the same value in return.
	*p = 1 //changes the value of i also, because p directly changes the value stored in that memory, which i just refs.

	// Slices are based on pointers
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	// Both of them will print the same as updating any of them updates the pointer.
}
