package main

import "fmt"

func main() {

	//var person map[string]string= map[string]string{}
	/*
		ini versi panjang tapi yaudahlah cukup tau aja!!!
	*/

	person := map[string]string{
		"name":    "Dito",
		"address": "Tangerang",
	}

	person["title"] = "Software Developer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	fmt.Println("----------------------")

	var book map[string]string = make(map[string]string)
	book["Title"] = "Ganteng Doang"
	book["author"] = "Dito"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
