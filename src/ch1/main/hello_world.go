package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) >1 {
		fmt.Println("Hello, world", os.Args[0])
	}

	os.Exit(-1)
}