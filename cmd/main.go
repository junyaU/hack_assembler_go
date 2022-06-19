package main

import (
	"flag"
	"fmt"
	"hack_assembler/parser"
	"log"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("specify asm file as an argument")
		os.Exit(1)
	}

	f, err := os.Open("testdata/" + flag.Args()[0])
	if err != nil {
		log.Println("")
	}

	defer f.Close()

	parser.NewParser(f)

}
