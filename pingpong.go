package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func icmp(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	fmt.Println("Pressione ENTER para finalizar")

	c := make(chan string)

	go ping(c)
	go pong(c)
	go icmp(c)

	var dig string
	fmt.Scanln(&dig)
}
