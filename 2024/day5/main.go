package main

import (
	"log"
	"os"

	"aoc/2024/day5/hashtable"
	"aoc/2024/day5/toposort"
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
