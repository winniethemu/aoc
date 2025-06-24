package main

import (
	day5 "2024/day5/internal"
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
		day5.RunHashtable()
	case "toposort":
		day5.RunToposort()
	default:
		log.Fatal("invalid argument name")
	}
}
