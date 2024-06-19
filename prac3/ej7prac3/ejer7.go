package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			ch1 <- rand.Intn(100)
			time.Sleep(1*time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			ch2 <- rand.Intn(100)
			time.Sleep(2*time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1;i <=6; i++ {
			select {
			case val := <-ch1: 
			fmt.Printf("Recibi %d , voy a esperar 5 segs \n" , val)
			//time.Sleep(1*time.Second)
			case  val := <-ch2: 
			fmt.Printf("Recibi %d , voy a esperar 10 segs \n" , val)
			//time.Sleep(2*time.Second)
			}
		}
	} ()

	wg.Wait()
}