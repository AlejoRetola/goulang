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
	for i, r := range runas {
		if unicode.IsUpper(r) {
			runas[i] = unicode.ToLower(runas[i]) // Convertir a minúscula si es mayúscula
		} else {
			runas[i] = unicode.ToUpper(runas[i]) // Convertir a mayúscula si es minúscula
		}
	}
	return string(runas) // Convertir la slice de runes de nuevo a una cadena
} 
func main() {
    var frase , palabra string
    fmt.Println("Escriba palabra a buscar")
    fmt.Scanln(&palabra)

    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Escriba una frase")
    frase , _ = reader.ReadString('\n')
    newFrase := ""
    inicioStr := 0

    match := strings.Index(strings.ToLower(frase) , palabra)

    for match != -1 {
            newFrase += frase[0:match] // corto desde el inicio al match de la frase y lo guardo sin cambiar
            inicioStr = match + len(palabra) // guardo la pos desde el match + la longitud de la palabra
            newFrase += procesarPalabra(frase[match: inicioStr]) // corto desde el match hasta donde termina la palabra y agrego la palabra cambiada
            frase = frase[inicioStr:] // voy guardando el "RESTITO" desde la posición donde terminó mi anterior match en adelante
            match = strings.Index(strings.ToLower(frase), palabra) // matcheo nuevamente para entrar al loop
    }

    newFrase += frase // Agregar cualquier resto de la frase que quede

    fmt.Println("Frase filtrada por ", palabra , "/n", newFrase)

}
