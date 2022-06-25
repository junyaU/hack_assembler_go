package main

import (
	"flag"
	"fmt"
	"hack_assembler"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("specify asm file as an argument")
		os.Exit(1)
	}

	fileName := flag.Args()[0]
	dataDir := "testdata/"

	f, err := os.Open(dataDir + fileName)
	if err != nil {
		fmt.Println("the specified file does not exist")
		os.Exit(1)
	}

	defer f.Close()

	p := hack_assembler_go.NewParser(f)
	c := hack_assembler_go.NewCode()

	for i := 0; i < len(p.Commands()); i++ {
		if !p.HasMoreCommands() {
			break
		}

		p.Advance()

		commandType, err := p.CommandType()
		if err != nil {
			fmt.Println("hack syntax error")
			os.Exit(1)
		}

		switch commandType {
		case hack_assembler_go.A_COMMAND:
			if err := c.WriteACommand(p.Symbol()); err != nil {
				fmt.Println("hack syntax error1")
				os.Exit(1)
			}
		case hack_assembler_go.C_COMMAND:
			if err := c.WriteCCommand(p.Dest(), p.Comp(), p.Jump()); err != nil {
				fmt.Println("hack syntax error2")
				fmt.Println(err)
				os.Exit(1)
			}
		case hack_assembler_go.L_COMMAND:
		}
	}

	extensionIndex := strings.Index(fileName, ".asm")
	hackFile := fileName[:extensionIndex] + ".hack"

	file, err := os.Create(dataDir + hackFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	for _, binary := range c.BinaryResult() {
		if _, err := file.Write(binary); err != nil {
		}

		file.Write([]byte("\n"))
	}

	fmt.Println("success compile!!")

}
