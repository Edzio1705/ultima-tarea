package main

import "fmt"

func main() {
	var palabra string
	fmt.Println("por favor ingrese una Palabra para comprobar si es un palindrome")
	fmt.Scanln(&palabra)
	fmt.Printf("Tu palabra es: '%s' veamos si tu palabra es un palindrome: %t\n", palabra, isPalindrome(palabra))
}

func isPalindrome(palabra string) bool {
	palogrande := len(palabra) - 1
	for i := 0; i < palogrande/2 && i < (palogrande-i); i++ {
		if palabra[i] != palabra[palogrande-i] {
			return false
		}
	}
	return true
}
