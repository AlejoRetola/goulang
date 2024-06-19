package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
	
var wg  sync.WaitGroup


func enviar(canal chan int) {
	for i := 1; i < 6; i++ {
		canal <- rand.Intn(100)
		time.Sleep(5)
	}
	wg.Done()
}

func main() {

	cha1 := make(chan int)
	cha2 := make(chan int)
	cha3 := make(chan int)

	wg.Add(4)

	go enviar(cha1)
	go enviar(cha2)
	go enviar(cha3)

	go func() {
		total1,total2,total3 := 0,0,0
		for i := 1; i < 16; i++ {
			select {
			case val := <-cha1:
				fmt.Println(val , " De canal 1")
				total1+= val
			case val := <-cha2:
				fmt.Println(val , " de canal 2")
				total2+= val
			case val := <-cha3:
				fmt.Println(val, " de canal 3")
				total3+= val
			}
		}

		fmt.Println("Total 1 = " , total1)
		fmt.Println("Total 2 = " , total2)
		fmt.Println("Total 3 = " , total3)
		wg.Done()
	}()
	wg.Wait()
}