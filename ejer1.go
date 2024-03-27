package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)


func procesarMartes(s string)  string {
    res := []rune("jueves")
    par := []rune(s)
    for i , l := range par {
            if unicode.IsUpper(l) {
               res[i] = unicode.ToUpper(res[i])
            } else {
               res[i] = unicode.ToLower(res[i])
            }
        }
    return string(res)
}
func procesarJueves(s string)  string {
    res := []rune("martes")
    par := []rune(s)
    for i , l := range par {
            if unicode.IsUpper(l) {
               res[i] = unicode.ToUpper(res[i])
            } else {
               res[i] = unicode.ToLower(res[i])
            }
        }
    return string(res)
}

/*
func procesarPalabra(palabra string) string {
	runas := []rune(palabra) // Convertir la palabra en una slice de runes
    fmt.Println(string(runas))
	for i, r := range runas {
		if unicode.IsUpper(r) {
			runas[i] = unicode.ToLower(r) // Convertir a minúscula si es mayúscula
		} else {
			runas[i] = unicode.ToUpper(r) // Convertir a mayúscula si es minúscula
		}
	}
	return string(runas) // Convertir la slice de runes de nuevo a una cadena
} */
func main() {
    var frase string
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Escriba una frase")
    frase , _ = reader.ReadString('\n')
    var match [] int
    res := strings.Fields(frase)
    for i , s := range res {
        if (strings.ToLower(s) == "martes") || (strings.ToLower(s) == "jueves") {
            match = append(match,i )
        }
    }
    for _ , s := range match {
        cambio := ""
       if strings.ToLower(res[s]) == "martes" {
            cambio = procesarMartes(res[s])
        } else {
           cambio = procesarJueves(res[s])
        } 
        res[s] = cambio
    }
    fmt.Println(strings.Join(res, " "))
}