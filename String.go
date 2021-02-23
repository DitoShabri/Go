package main

import "fmt"

func main() {
	fmt.Println("Firli")
	fmt.Println("Dito")
	fmt.Println("Sabri")

	fmt.Println(len("Firli"))
	fmt.Println("Dito"[0])
	fmt.Println("Sabri"[1])

	var a = "Dito"
	var b = (a)[0]
	var c = string(b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var e = "Sabri"
	var f = (e)[1]
	var g = string(f)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
