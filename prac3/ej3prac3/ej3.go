package main

import (
	"fmt"
)

func hello(c chan string) {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
			c <- "Hello World"
		}
	fmt.Println("Termina Goroutine de hello")
}


func main() {
	c := make(chan  string)
	fmt.Println("Inicia Goroutine del main")
	go hello(c)
	for i := 0; i < 3; i ++ {
		fmt.Println(<-c)
	}
	fmt.Println("Termina Goroutine del main")
}



/*
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("Inicia Goroutine del main")
	go hello(&wg)
	wg.Wait()
	fmt.Println("Termina Goroutine del main")

	ASI ES CON WAITGROUP , sino tenes que calcular un time.Sleep() al final del ejer
	*/