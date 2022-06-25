package hack_assembler_go

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Parser struct {
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

	return &Parser{
		texts: rowTexts,
	}
}

func (p Parser) Commands() []string {
	return p.texts
}

func (p Parser) HasMoreCommands() bool {
	return len(p.texts) > p.currentLine
}

func (p *Parser) Advance() {
	p.command = p.texts[p.currentLine]
	p.currentLine += 1
}

func (p Parser) CommandType() (CommandType, error) {
	isACommand := strings.Index(p.command, "@") == 0
	isLCommand := strings.HasPrefix(p.command, "(") && strings.HasSuffix(p.command, ")")
	isCCommand := strings.Contains(p.command, "=") || strings.Contains(p.command, ";")

	switch {
	case isACommand:
		return A_COMMAND, nil
	case isLCommand:
		return L_COMMAND, nil
	case isCCommand:
		return C_COMMAND, nil
	default:
		return 0, errors.New("this command does not exist")
	}
}

func (p Parser) Symbol() string {
	isACommand := strings.Index(p.command, "@") == 0
	if isACommand {
		return p.command[1:]
	}

	symbol := strings.TrimLeft(p.command, "(")
	return strings.TrimRight(symbol, ")")
}

func (p Parser) Dest() string {
	isExistDest := strings.Contains(p.command, "=")
	if !isExistDest {
		return ""
	}

	return p.command[:strings.Index(p.command, "=")]
}

func (p Parser) Comp() string {
	isExistDest := strings.Contains(p.command, "=")
	if isExistDest {
		return p.command[strings.Index(p.command, "=")+1:]
	}

	return p.command[:strings.Index(p.command, ";")]
}

func (p Parser) Jump() string {
	isExistJump := strings.Contains(p.command, ";")
	if !isExistJump {
		return ""
	}

	return p.command[strings.Index(p.command, ";")+1:]
}
