package main

import (
	"flag"
	"fmt"
	"hack_assembler"
	"os"
	"strconv"
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
	table := hack_assembler_go.NewSymbolTable()
	c := hack_assembler_go.NewCode()

	if err := DefineSymbol(p, table); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p.ClearLoad()

	if err := WriteBinaries(p, c, table); err != nil {
		fmt.Println(err)
		os.Exit(1)
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
		file.Write(binary)
		file.Write([]byte("\n"))
	}

	fmt.Println("success compile!!")
}

func DefineSymbol(p *hack_assembler_go.Parser, t *hack_assembler_go.SymbolTable) error {
	currentFileLine := 0
	for i := 0; i < len(p.Commands()); i++ {
		if !p.HasMoreCommands() {
			break
		}

		p.Advance()

		commandType, err := p.CommandType()
		if err != nil {
			return err
		}

		if commandType != hack_assembler_go.L_COMMAND {
			currentFileLine++
		}

		switch commandType {
		case hack_assembler_go.A_COMMAND, hack_assembler_go.L_COMMAND:
			if t.IsNonNumericSymbol(p.Symbol()) {
				t.AddEntry(currentFileLine, p.Symbol(), commandType)
			}
		}
	}

	return nil
}

func WriteBinaries(p *hack_assembler_go.Parser, c *hack_assembler_go.Code, t *hack_assembler_go.SymbolTable) error {
	for i := 0; i < len(p.Commands()); i++ {
		if !p.HasMoreCommands() {
			break
		}

		p.Advance()

		commandType, err := p.CommandType()
		if err != nil {
			return err
		}

		switch commandType {
		case hack_assembler_go.A_COMMAND:
			symbol := p.Symbol()
			if t.IsNonNumericSymbol(symbol) && t.Contains(symbol) {
				address := strconv.Itoa(t.GetAddress(symbol))
				symbol = address
			}

			if err := c.WriteACommand(symbol); err != nil {
				return err
			}

		case hack_assembler_go.C_COMMAND:
			if err := c.WriteCCommand(p.Dest(), p.Comp(), p.Jump()); err != nil {
				return err
			}
		}
	}
	return nil
}
