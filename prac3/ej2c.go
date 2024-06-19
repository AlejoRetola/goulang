package main

import (
	"fmt"
	"sync"
	"time"
)

type caja struct {
	id   int
	cola chan string
}

func atenderCola(cola chan string) {
	for range cola {
		cliente := <-cola
		fmt.Println("Atendiendo al cliente ", cliente)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	var wg sync.WaitGroup
	cajas := make([]caja, 5)
	nombres := []string{
		"Ana", "Carlos", "María", "José", "Luis",
		"Sofía", "Miguel", "Lucía", "Juan", "Laura",
		"Andrés", "Elena", "Javier", "Isabel", "David",
		"Paula", "Sergio", "Marta", "Pedro",
	}

	wg.Add(1)
	//go func() {
	for i, caja := range cajas {
		caja.id = i + 1
		caja.cola = make(chan string, 10)
		go atenderCola(caja.cola)
		fmt.Println("CREE LA CAJA ", caja)
	}
	//}()
	//time.Sleep(5 * time.Second)
	//go func() {
	for _, persona := range nombres {
		colaCorta := 9999
		enviar := 0
		for i, caja := range cajas {
			if len(caja.cola) < colaCorta {
				colaCorta = len(caja.cola)
				enviar = i
			}
		}
		fmt.Println("me rompo :((( ", cajas)
		cajas[enviar].cola <- persona
	}
	for _, caja := range cajas {
		close(caja.cola)
	}
	wg.Done()
	//}()

	wg.Wait()
	fmt.Println("Aguante boca jiji jaja")
}
