package main

import "fmt"

func main() {
	x := "ma"

	if x == "th" {
		fmt.Println("Thailand")
	} else if x == "jp" {
		fmt.Println("Japan")
	} else if x == "ma" {
		fmt.Println("M<alaysia")
	} else {
		fmt.Println("ERROR")
	}
}
