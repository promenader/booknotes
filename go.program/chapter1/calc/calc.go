package main

import (
	"fmt"
	"os"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are :\n\t add \t Addition of two values. \n\t sqrt \t Square root of a non-negative value.")
}

func main()  {
	args := os.Args
	if args == nil ||len(args) <2 {
		Usage()
		return
	}

	switch args[0] {
	case "add":
		
	}
}
