package main

import "fmt"

var version = "dev"

func main() {

	fmt.Printf("Version: %s\n", version)
	fmt.Print("HI  test vignesh")
	fmt.Println(hello())
}

func hello() string {
	return "Hello Golang"
}
