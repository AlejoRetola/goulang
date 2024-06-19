package main

import (
	"fmt"
)
func pxng(m chan string, str string) {
		m <- str
}


func main() {
	messages := make(chan string , 10)
	for i := 0; i < 5; i++ {
		go pxng(messages, "PING")
		go pxng(messages, "PONG")
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-messages)
	}
}
