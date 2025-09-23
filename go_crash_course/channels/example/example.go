package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	fmt.Println("Real World Channels Example")
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrice(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofu_price = rand.Float32() * 20
		if tofu_price < MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func checkChickenPrice(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chicken_price = rand.Float32() * 20
		if chicken_price < MAX_CHICKEN_PRICE {
			c <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found deal on chicken at %v.", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent: Found deal on tofu at %v.", website)
	}
}
