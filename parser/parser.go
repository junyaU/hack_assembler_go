package parser

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type Parser struct {
	stream      io.Reader
	texts       []string
	currentLine int
	command     string
}

func NewParser(f io.Reader) *Parser {
	scanner := bufio.NewScanner(f)
	var rowTexts []string

	for scanner.Scan() {
		text := scanner.Text()

		commentOutIndex := strings.Index(text, "//")
		if commentOutIndex != -1 {
			text = text[:commentOutIndex]
		}

		if text == "" {
			continue
		}

		rowTexts = append(rowTexts, text)
	}

	log.Println(rowTexts)

	return &Parser{
		texts: rowTexts,
	}
}

func (p Parser) HasMoreCommands() bool {
	return len(p.texts) > p.currentLine
}

func (p *Parser) Advance() {

}
