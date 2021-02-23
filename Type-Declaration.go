package main

import "fmt"

func main() {

	type lahir string
	type Sibukga bool

	var dimanalahir = lahir("Tangerang")
	var lagisibukga = Sibukga(true)

	fmt.Println(dimanalahir)
	fmt.Println(lagisibukga)
}
