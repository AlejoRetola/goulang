package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func procesarPalabra(s string ) string {
	res := []rune(s)
	for i , l := range res {
		if unicode.IsUpper(l) {
			res[i] = unicode.ToLower(res[i])
		} else {
			res[i] = unicode.ToUpper(res[i])
		}
	}
	return string(res)
}
func main() {
    var frase string
	var palabra string
    var match [] int
	//Leo frase y palabra
	fmt.Println("Escriba una palabra a buscar")
	fmt.Scanln(&palabra)
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Escriba una frase")
    frase , _ = reader.ReadString('\n')
	// separo en palabras en un slice
    res := strings.Fields(frase)
	// itero sobre el slice de las palabras
    for i , s := range res {
        if (strings.ToLower(s) == strings.ToLower(palabra))  {
            match = append(match,i )
        }
    }
    for _ , s := range match {
		res[s] = procesarPalabra(res[s])
    }
    fmt.Println(strings.Join(res, " "))
}
