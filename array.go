package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Firli"
	names[1] = "Suryawan"
	names[2] = "Sabri"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		85,
	}

	fmt.Println(values)
	fmt.Println(values[0])

	fmt.Println(len(names))
}
