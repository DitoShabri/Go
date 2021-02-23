package main

import "fmt"

func main() {

	name := "Dito"

	if name == "Firli" {
		fmt.Println("ini Firli ")
	} else if name == "Enggar" {
		fmt.Println("Enggar yak ?")
	} else {
		fmt.Println("Intro aja dulu")
	}

	if length := len(name); length > 5 {
		fmt.Println("Panjang Juga nama lu ?")
	} else {
		fmt.Println("Nahh ini baru")
	}
}
