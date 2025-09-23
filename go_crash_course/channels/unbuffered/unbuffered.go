package main

import (
	"fmt"
)

// func main(){
// 	fmt.Println("SO far, SO good!")
// 	var c = make(chan int)
// 	c <- 1
// 	var i = <- c
// 	fmt.Println(i)
// } creates deadlock, solution below

// unbuffered channel
/*
func main(){
	var c = make(chan int)
	go process(c)
	fmt.Println(<-c)
}

func process(c chan int){
	 c <- 123
}
*/

// loop method with defer keyword
func main() {
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
