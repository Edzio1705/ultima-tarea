package main

import (
	"fmt"
	"strings"
)

func main() {
	var Tex1 string
	var Tex2 string

	fmt.Println("por favor ingrese la oracion")
	fmt.Scanln(&Tex1)
	fmt.Println("ingrese la palabra a buscar dentro de la oracion.")
	fmt.Scanln(&Tex2)
	fmt.Println(strings.Contains(Tex1, Tex2))

}
