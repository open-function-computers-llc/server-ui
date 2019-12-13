package main

import "fmt"

func main() {
	err := runApp()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
