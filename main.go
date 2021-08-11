package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("-----------------------------")
	fmt.Printf("\nGOOS:%s", runtime.GOOS)
	fmt.Printf("\nGOARCH:%s", runtime.GOARCH)
	fmt.Printf("\nGOROOT:%s", runtime.GOROOT())
	fmt.Printf("\nCompiler:%s", runtime.Compiler)
	fmt.Printf("\nNo. of CPU:%d", runtime.NumCPU())	
}
