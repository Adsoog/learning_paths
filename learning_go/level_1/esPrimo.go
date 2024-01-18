package main

import "fmt"

func esPrimo(numero int) bool {
	// Verifica si 'numero' es primo
	if numero <= 1 {
		return false // 0 y 1 no son primos
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false // Si es divisible por algún número, no es primo
		}
	}

	return true
}

func main() {
	// Encuentra y muestra los primeros 10 números primos
	contador := 0
	numero := 2 // Comienza con el primer número primo

	for contador < 10 {
		if esPrimo(numero) {
			fmt.Printf("%d es primo\n", numero)
			contador++
		}
		numero++
	}
}
