package main

import "fmt"

func imprimir() {
	fmt.Println("Texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "hola" + s
}

func suma(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("A y B son iguales")
	} else {
		fmt.Println("A y B no son iguales")
	}
}

func main() {
	//fmt.Println("Texto en la funcion main")
	//imprimir()
	// fmt.Println(hola_string(" Camilo"))
	// fmt.Println(suma(3, 5))
	comparar_bool()
}
