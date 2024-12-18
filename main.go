package main

import (
	"os"
	"github.com/bhat-abhishek/cubg/cmd"
)

func main() { 
	err := cmd.Execute()
	if err != nil { 
		os.Exit(1)
	}
}

