package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// HAY QUE REHACERLO CON .INDEX LRPM

func procesarPalabra(palabra string) string {
	runas := []rune(palabra) // Convertir la palabra en una slice de runes
    change := make([]rune,9)
    //
    if(strings.ToLower(palabra) == "miercoles") {
        change = []rune("automovil")
        } else {
            change = []rune("miercoles")
        }
    //  
	for i, r := range runas {
		if unicode.IsUpper(r) {
			change[i] = unicode.ToLower(change[i]) // Convertir a minúscula si es mayúscula
		} else {
			change[i] = unicode.ToUpper(change[i]) // Convertir a mayúscula si es minúscula
		}
	}
	return string(change) // Convertir la slice de runes de nuevo a una cadena
} 
func main() {
    var frase string
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Escriba una frase")
    frase , _ = reader.ReadString('\n')
    newFrase := ""
    inicioStr := 0
    /*
    // MATCHEO POR MARTES
    match := strings.Index(strings.ToLower(frase), "martes")
    for match != -1 {
        newFrase += frase[0:match] // desde el inicio de lstring, hasta el match, lo guardo en el newString
        inicioStr = match + 6 // no se suma, se iguala para saber la pos donde cortar el string de parametro donde trabajamos
        newFrase +=  procesarPalabra(frase[match:match+6]) // proceso desde el match, mas la cantidad de letras que buscamos
        frase = frase[inicioStr:] // voy guardando el "RESTITO" desde la posiciond donde termino mi anterior match en adelante
        match = strings.Index(frase, "martes") // matcheo nuevamente para entrar al loop
    }
   fmt.Println("Frase filtrada por martes" ,newFrase)

    DESCARTADO PORQUE NO FUNCIONA PARA LOS DOS, TENGO QUE USAR CONDICIONALES COMO ABAJ OY BLABLALBLABLAL

    */
// MATCHEO POR MARTES Y JUEVES
    matchMiercoles := strings.Index(strings.ToLower(frase), "miercoles")
    matchAutomovil := strings.Index(strings.ToLower(frase), "automovil")

    for matchMiercoles != -1 || matchAutomovil != -1 {
        if matchMiercoles != -1 && (matchMiercoles < matchAutomovil || matchAutomovil == -1) {
            newFrase += frase[0:matchMiercoles]
            inicioStr = matchMiercoles + 9
            newFrase += procesarPalabra(frase[matchMiercoles : matchMiercoles+9])
            frase = frase[inicioStr:] // voy guardando el "RESTITO" desde la posición donde terminó mi anterior match en adelante
            matchMiercoles = strings.Index(strings.ToLower(frase), "miercoles") // matcheo nuevamente para entrar al loop
            matchAutomovil = strings.Index(strings.ToLower(frase), "automovil") // prevencion a errores con e lindex ( ya me paso)
        } else if matchAutomovil != -1 && (matchAutomovil < matchMiercoles || matchMiercoles == -1) {
            newFrase += frase[0:matchAutomovil]
            inicioStr = matchAutomovil + 9
            newFrase += procesarPalabra(frase[matchAutomovil : matchAutomovil+9])
            frase = frase[inicioStr:] // voy guardando el "RESTITO" desde la posición donde terminó mi anterior match en adelante
            matchAutomovil = strings.Index(strings.ToLower(frase), "automovil") // matcheo nuevamente para entrar al loop
            matchMiercoles = strings.Index(strings.ToLower(frase), "miercoles") // matcheo martes por posibles errores en el manejo de index
        }
    }

    newFrase += frase // Agregar cualquier resto de la frase que quede

    fmt.Println("Frase filtrada por miercoles y automovil:", newFrase)

}
