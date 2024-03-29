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
    change := make([]rune, 6)
    //
    if(strings.ToLower(palabra) == "martes") {
        change = []rune("jueves")
        } else {
            change = []rune("martes")
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
    matchMartes := strings.Index(strings.ToLower(frase), "martes")
    matchJueves := strings.Index(strings.ToLower(frase), "jueves")

    for matchMartes != -1 || matchJueves != -1 {
        if matchMartes != -1 && (matchMartes < matchJueves || matchJueves == -1) {
            newFrase += frase[0:matchMartes]
            inicioStr = matchMartes + 6
            newFrase += procesarPalabra(frase[matchMartes : matchMartes+6])
            frase = frase[inicioStr:] // voy guardando el "RESTITO" desde la posición donde terminó mi anterior match en adelante
            matchMartes = strings.Index(strings.ToLower(frase), "martes") // matcheo nuevamente para entrar al loop
            matchJueves = strings.Index(strings.ToLower(frase), "jueves") // prevencion a errores con e lindex ( ya me paso)
        } else if matchJueves != -1 && (matchJueves < matchMartes || matchMartes == -1) {
            newFrase += frase[0:matchJueves]
            inicioStr = matchJueves + 6
            newFrase += procesarPalabra(frase[matchJueves : matchJueves+6])
            frase = frase[inicioStr:] // voy guardando el "RESTITO" desde la posición donde terminó mi anterior match en adelante
            matchJueves = strings.Index(strings.ToLower(frase), "jueves") // matcheo nuevamente para entrar al loop
            matchMartes = strings.Index(strings.ToLower(frase), "martes") // matcheo martes por posibles errores en el manejo de index
        }
    }

    newFrase += frase // Agregar cualquier resto de la frase que quede

    fmt.Println("Frase filtrada por martes y jueves:", newFrase)

}
