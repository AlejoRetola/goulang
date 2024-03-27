package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func procesarPalabra(s string) string {
	par := []rune(s)
	res := make([]rune, 9)
	if strings.ToLower(s) == "miercoles" {
		res = []rune("automovil")
	} else {
		res = []rune("miercoles")
	}
	for i , l := range par {
		if unicode.IsUpper(l) {
			res[i] = unicode.ToUpper(res[i])
		} else {
			res[i] = unicode.ToLower(res[i])
		}
	}
	return string(res)
}
func main() {
    var frase string
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Escriba una frase")
    frase , _ = reader.ReadString('\n')
    var match [] int
    res := strings.Fields(frase)
    for i , s := range res {
        if (strings.ToLower(s) == "miercoles") || (strings.ToLower(s) == "automovil") {
            match = append(match,i )
        }
    }
    for _ , s := range match {
		res[s] = procesarPalabra(res[s]) 
    }
    fmt.Println(strings.Join(res, " "))
}
var2 := 0
loop

str1 = 
losmiercolesmiercolesautomovil
var := str1.index

 str2 = str1[var2:var]
 str2 *= procesar(var: var+9)
var + 9
var2 := var + 9
condicion si -1
 break
end;


str2 += lo que queda 

str2 = "losMIERCOLESMIERCOLES"
