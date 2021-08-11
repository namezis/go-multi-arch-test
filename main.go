package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("Platform OS is: %s\n", runtime.GOOS)
}
