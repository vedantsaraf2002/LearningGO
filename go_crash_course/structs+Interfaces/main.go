package main

import (
	"fmt"
)

// type gasEngine struct {
// 	mpg       uint8
// 	gallons   uint8
// 	ownerInfo owner
// }

// can also do directly
type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricalEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricalEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	fmt.Println("Hello Again Vedant!")
	var myEngine gasEngine = gasEngine{25, 15, owner{"Vedant"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Printf("Miles left %v", myEngine.milesLeft())
	var myEngine2 electricalEngine = electricalEngine{25, 15}
	canMakeIt(myEngine, 50)
	canMakeIt(myEngine2, 50)
}
