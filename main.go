package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("Platform is: %s\n", runtime.GOOS)
}
