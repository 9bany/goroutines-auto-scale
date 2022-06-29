package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asyncReceiver(c chan int, name string) {
	for {
		if shouldScale(c) < 0 {
			fmt.Println("End go rountine: " + name)
			return
		}
		<-c
		r := rand.Intn(100) + 300
		time.Sleep(time.Duration(r) * time.Millisecond) 
	}
}

func asyncSender(c chan int) {
	for {
		c <- 1
		r := rand.Intn(100) + 100
		time.Sleep(time.Duration(r) * time.Millisecond)
	}
}

func shouldScale(c chan int) int {
	capacity := cap(c)
	orders := len(c)

	if orders >= capacity-5 {
		return 1
	} else if orders <= 5 {
		return -1
	}
	return 0

}

func main() {
	receiveCounter := 0
	c := make(chan int, 20)
	go asyncReceiver(c, fmt.Sprintf("receier: %d", receiveCounter))
	go asyncSender(c)
	for {
		l := len(c)
		fmt.Println(l)
		if shouldScale(c) > 0 {
			name:= fmt.Sprintf("receier: %d", receiveCounter)
			fmt.Println("Run new go rountine with name: ", name)
			receiveCounter++
			go asyncReceiver(c, name)
		}
		time.Sleep(time.Second)
	}
}
