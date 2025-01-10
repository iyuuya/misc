package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: choise <arg1> <arg2> ... <argN>")
		os.Exit(1)
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(args))
	selectedArg := args[randomIndex]

	fmt.Println(selectedArg)
}
