package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type  RR struct {
	nombre string
	turno int
}

var cerrado = make(chan bool)
var filaGlobal = make(chan string) // la fil maximo es de 10 xq si

func atenderCliente(wg *sync.WaitGroup) {
	defer wg.Done()
	for cliente := range filaGlobal {
		fmt.Println("HOLA CLIENTE, YO LO ATIENDO JIJIJA" , cliente)
		time.Sleep(time.Duration(rand.Int63n(1)) * time.Second)
	}
} 

func atenderRobin(proceso *RR, tiempo int , done chan RR) {
		if proceso.turno > tiempo {
			fmt.Println("Atiendo al cliente " +proceso.nombre+" todo bien, pero no alcanza el tiempo")
			time.Sleep(time.Duration(proceso.turno) * time.Millisecond)
			proceso.turno -= tiempo
		} else {
			fmt.Println("Atiendo al cliente "+proceso.nombre+" todo bien, tiempo suficiente")
			time.Sleep(time.Duration(proceso.turno) * time.Millisecond)
			proceso.turno = 0
			cerrado <- true
		}
	
	done <- *proceso
	fmt.Println(proceso.nombre , " vuelve")
}

func cerrarDone(done chan RR , cantCanales int) {
	canalesCerrados := 0
	for canalesCerrados != cantCanales {
		<- cerrado 
		canalesCerrados++
	}
	close(done)
}

func adminRobin(done chan RR , roundRobin []RR , mu *sync.Mutex) {
	for proceso := range done { // NO ES CON UN FOR
		fmt.Println("Recibo " , proceso.nombre , proceso.turno)
		if proceso.turno > 0 { // si le queda tiempo en su turno, no termino
				mu.Lock()
				roundRobin = append(roundRobin, proceso) // asi que lo agrego al final again
				mu.Unlock()
			}
		}	
}

func main() {
	var espera sync.WaitGroup

	nombres := []string{
        "Ana", "Carlos", "María", "José", "Luis",
        "Sofía", "Miguel", "Lucía", "Juan", "Laura",
        "Andrés", "Elena", "Javier", "Isabel", "David",
        "Paula", "Sergio", "Marta", "Pedro", "Natalia",
    }

	roundRobin := []RR {
		{"Ana" , 150},
		{"Carlos", 250},
		{"Maria", 170},
		{"jose", 350},
		{"Luis", 200},
	}

	fmt.Println("ALA MADRID " , nombres[1])
	/* PUNTO A
	espera.Add(1)
	go atenderCliente(&espera) // PRIMERO LLEGA EL CAJERO Y DPS LOS CLIENTES PA
	go func () {
		for i := range nombres {
			filaGlobal <- nombres[i]
		}
		close(filaGlobal)
		} ()
	*/
	espera.Add(1)
	done := make(chan RR)
	info := make(chan []RR)
	var mu sync.Mutex
	go func () {
		tiempo := 100
		defer espera.Done()
		go cerrarDone(done , len(roundRobin))
		adminRobin(done, roundRobin, &mu)

		for len(roundRobin) > 0 { // mientras haya procesos en la cola
			//proceso := roundRobin[0] // el primer proceso
			//roundRobin = roundRobin[1:] // lo que resta del arr
			for range len(roundRobin) {
				proceso := roundRobin[0]
				roundRobin = roundRobin[1:]
				fmt.Println("Mando " , proceso.nombre , " " , roundRobin)
				go atenderRobin(&proceso , tiempo , done ) // mando todos los procesos
			}
			roundRobin = <- info
		}
	} ()
	
	

	espera.Wait()
}