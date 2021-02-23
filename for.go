package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"Firli", "Dito", "Sabri", "Adit", "Enggar"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	/*
		"didalam Go semua variable harus digunakan, maka dari itu anda bisa menggunakan
			underscore "_" untuk menyatakan kepada Golang bahwa variable ini
			belum di perlukan"
	*/
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Dito"
	person["Title"] = "Programer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
