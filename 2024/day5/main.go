package main

import (
	"2024/day5/hashtable"
	"2024/day5/toposort"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("invalid arguments")
	}
	switch args[0] {
	case "hashtable":
		hashtable.Run()
	case "toposort":
		toposort.Run()
	default:
		log.Fatal("invalid argument name")
	}
}
