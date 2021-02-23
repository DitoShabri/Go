package main

import "fmt"

func main() {

	var months = [...]string{
		"Januari",
		"febuari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[5] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0]= "MeiLagi"
	//fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Dito")
	fmt.Println(slice3)

	slice3[1] = "BackToDesember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newslice := make([]string, 2, 5)

	newslice[0] = "Dito"
	newslice[1] = "Sabri"

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	copySlice := make([]string, len(newslice), cap(newslice))
	copy(copySlice, newslice)

	fmt.Println(copySlice)
}
