package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

var wg sync.WaitGroup
var done = make(chan bool)
 var sliceFinal = make(chan []int)

func dividirPrimos(inicio int , fin int , ch chan int) {
	defer wg.Done()

	for inicio != fin  {
		if(big.NewInt(int64(inicio)).ProbablyPrime(0) ) { //checkeo si es primo
			ch <- inicio // lo mando por mi canal canales[i]
		}
		inicio++
	}
	done <- true
	close(ch)
}

func procesarPrimos(canales[4] chan int) {
	defer wg.Done() 
	boca := make([]int , 0)
	cerrarCanales := 0
	for cerrarCanales != len(canales) {
		select {
		case val , ok := <-canales[0] : if ok { boca = append(boca, val) }   // aca tenia un else
		case val , ok := <-canales[1] :	if ok { boca = append(boca, val) }	// de else { cerrarCanales++}
		case val , ok := <-canales[2] :	if ok { boca = append(boca, val) }	
		case val , ok := <-canales[3] :	if ok { boca = append(boca, val)}	
		case <-done: cerrarCanales++
	}
 	}
	sliceFinal <- boca
}


func unicaRutina(desde int , hasta int) {
	/*a) Realice el programa con una única goroutine que muestre en pantalla la lista
		de números primos encontrados.
	*/
	defer wg.Done()
	for desde != hasta {
		if big.NewInt(int64(desde)).ProbablyPrime(0) {
			fmt.Println("Es primo el " , desde)
		}
		desde++
	}
}

func primesFunc(ch chan int, ini int, fin int) {
	defer wg.Done()
	for ini != fin {
		if big.NewInt(int64(ini)).ProbablyPrime(0) {
			ch <-ini
		}
		ini++
	}
	done <- true
}

func main() {
	var num int //  mi numero a calcular
	var ch = make(chan int , 4)
	var primeList = make([]int , 0)

	cantRutinas := 4 

	fmt.Scanln(&num) // escaneo el numero
	timeStart := time.Now() // comienz ami programa
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for aux := range ch {
			primeList = append(primeList, aux)
		
		}
	}()

	// Goroutine para cerrar el canal después de que todas las primesFuncs hayan terminado
	go func() {
		for range cantRutinas {
			<-done
		}
		close(ch)
	}()

	rango := num / cantRutinas 
	for i := 0; i < cantRutinas; i++ {
			inicio := i * rango // que porcion le toca, 3 * 25 seria del 75 al 100 por ejemplo
			fin :=  (i + 1) * rango
			
			if fin > num {
				fin = num // si me pase, lo limito al numero
			}
			
			wg.Add(1)
			go primesFunc(ch,inicio,fin) // inicio un procesamiento de los primos de ese cacho
		}
	wg.Wait()
	fmt.Println("Termine la rutina" , time.Since(timeStart))
	// fmt.Println(primeList)
	}


	// go unicaRutina(1 , num + 1) ESTE ES EL PUNTO a)


	/*
	var canales [4]chan int // creo 1 canal para cada rutina
	for i := range cantRutinas {
		canales[i] = make(chan int) // los inicio
	}	

	
	rango := num / cantRutinas // calculo de manera equitativa cuando le tengo que enviar a cada rutina 100/ 4 = 25 , 25 a cada rutina
	
	wg.Add(1) // Espero la func que acabo de crear anonima
	go func() {
		defer wg.Done()	
		wg.Add(1) // agrego el Add para mi procesarPrimos
		go procesarPrimos(canales) // envio el slice, y canales
		time.Sleep(5*time.Second)
		for i := 0; i < cantRutinas; i++ {
			inicio := i * rango // que porcion le toca, 3 * 25 seria del 75 al 100 por ejemplo
			fin :=  (i + 1) * rango
			
			if fin > num {
				fin = num // si me pase, lo limito al numero
			}
			
			wg.Add(1)
			go dividirPrimos(inicio,fin,canales[i]) // inicio un procesamiento de los primos de ese cacho
		}
		
		} ()
		
		
		listaPrimos := <- sliceFinal // se queda colgaod post wait, xq no deja hacer el defer al procesarPrimos
		wg.Wait()	
		
		
		
		fmt.Println("Fin del codigo " , time.Since(timeStart)) // ESTO TARDO CON 4 RUTINAS
		
		
	if len(listaPrimos) > 0 {
		fmt.Println(listaPrimos)
	} */