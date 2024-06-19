package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)




func main() {
	var wg  sync.WaitGroup
	ch1 := make(chan int)
	wg.Add(4)	
	go func() {
		defer wg.Done()
		for i := 1;i <= 3; i++ {
			//waitTime := rand.Intn(1)
			time.Sleep(1)
			ch1 <- rand.Intn(100)
		}
	} ()
	go func() {
		defer wg.Done()
		for i := 1;i <= 3; i++ {
			//waitTime := rand.Intn(1)
			time.Sleep(1)
			ch1 <- rand.Intn(100)
		}
	} ()
	
	go func() {
		defer wg.Done()

		for i := 1; i <= 3; i++ {
			fmt.Println("Consumido por el UNO , " ,<-ch1)
		}
	} ()
	go func() {
		defer wg.Done()

		for i := 1; i <= 3; i++ {
			fmt.Println("Consumido por el DOS , " ,<-ch1)
		}
	} ()

	wg.Wait()
}