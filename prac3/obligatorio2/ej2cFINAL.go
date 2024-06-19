package main

import (
	"fmt"
	"sync"
	"time"
)

var done = make(chan bool)

func atenderCaja(cola chan string, i int) {
	for cliente := range cola {
		fmt.Println("Atendiendo al cliente ", cliente, " en cola ", i, " longitud ", len(cola))
		time.Sleep(1 * time.Second)
		done <- true // dicho por los ayudantes que tenia uqe avisar por cada cliente que terminaba
	}

}

func main() {
	var wg sync.WaitGroup
	cajas := make([]chan string, 5)
	nombres := []string{
		"Ana", "Carlos", "María", "José", "Luis",
		"Sofía", "Miguel", "Lucía", "Juan", "Laura",
		"Andrés", "Elena", "Javier", "Isabel", "David",
		"Paula", "Sergio", "Marta", "Pedro",
	}

	for i := range cajas { // ESTE RANGE FUNCA, RANGE SI PONGO _ , CAJA := RANGE CAJAS NO
		cajas[i] = make(chan string, 100)
		go atenderCaja(cajas[i], i)
	}
	wg.Add(1)
	go func() {
		for _, persona := range nombres {
			colaCorta := 9999
			enviar := 0
			for i := range cajas {
				if len(cajas[i]) < colaCorta {
					colaCorta = len(cajas[i])
					enviar = i
				}
			}
			//fmt.Println("Envie a ", persona, " a la cola ", enviar, " de len ", len(cajas[enviar]))
			cajas[enviar] <- persona
		}
	}()

	go func() {
		defer wg.Done()
		for range nombres {
			<-done // y cuando todos los clientes terminasen, cerra canales
		}
		for i := range cajas {
			close(cajas[i])
		}
	}()

	wg.Wait()
}
