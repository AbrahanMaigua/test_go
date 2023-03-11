package main

import "fmt"

// las funciones debes declar cual sera la salida
func greting(name string) string {
	if name == "" {
		return "ola, Mundo"
	}
	return "ola, " + name
}

// function main
func main() {
	fmt.Println(greting(""))
	fmt.Println(greting("Wold"))

}
