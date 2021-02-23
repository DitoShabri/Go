package main

import "fmt"

func main() {

	var name = "Firli"

	switch name {
	case "Adit":
		fmt.Println("ini Adit")
	case "Enggar":
		fmt.Println("Ini Enggar")
	default:
		fmt.Println("Udahlah siapa aja bantu dong")

	}

	//switch length := len(name);length > 5{
	//case true:
	//	fmt.Println("Nama nya Panjang Masa")
	//case false:
	//	fmt.Println("Nahh ini mantu idaman")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Antum Panjang")
	case length > 5:
		fmt.Println("Nama Cukup panjang")
	default:
		fmt.Println("Nama Udah bagus ko")
	}

}
