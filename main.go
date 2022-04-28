package main

import "fmt"

var version = "dev"

func main() {

	fmt.Printf("Version: %s\n", version)
	fmt.Print("HI push test")
	fmt.Println(hello())
}

func hello() string {
	return "Hello Glang"
}
