package parser

import (
	"bufio"
	"errors"
	"io"
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

	return &Parser{
		texts: rowTexts,
	}
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
		return 0, errors.New("存在しないコマンドです")
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

func (p Parser) Dest() (string, error) {
	isExistDest := strings.Contains(p.command, "=")
	if !isExistDest {
		return "000", nil
	}

	switch p.command {
	case "M":
		return "001", nil
	case "D":
		return "010", nil
	case "MD":
		return "011", nil
	case "A":
		return "100", nil
	case "AM":
		return "101", nil
	case "AD":
		return "110", nil
	case "AMD":
		return "111", nil
	default:
		return "", errors.New("このコマンドに対応することができません")
	}
}

func (p Parser) Comp() (string, error) {
	var mnemonic string
	isExistDest := strings.Contains(p.command, "=")
	if isExistDest {
		mnemonic = p.command[strings.Index(p.command, "="):]
	} else {
		mnemonic = p.command[:strings.Index(p.command, ";")]
	}

	switch mnemonic {
	case "0":
		return "0101010", nil
	case "1":
		return "0111111", nil
	case "-1":
		return "0111010", nil
	case "D":
		return "0001100", nil
	case "A":
		return "0110000", nil
	case "!D":
		return "0001101", nil
	case "!A":
		return "0110001", nil
	case "-D":
		return "0001111", nil
	case "-A":
		return "0110011", nil
	case "D+1":
		return "0011111", nil
	case "A+1":
		return "0110111", nil
	case "D-1":
		return "0001110", nil
	case "A-1":
		return "0110010", nil
	case "D+A":
		return "0000010", nil
	case "D-A":
		return "0010011", nil
	case "A-D":
		return "0000111", nil
	case "D&A":
		return "0000000", nil
	case "D|A":
		return "0010101", nil
	case "M":
		return "1110000", nil
	case "!M":
		return "1110000", nil
	case "-M":
		return "1110011", nil
	case "M+1":
		return "1110111", nil
	case "M-1":
		return "1110010", nil
	case "D+M":
		return "1000010", nil
	case "D-M":
		return "1010011", nil
	case "M-D":
		return "1000111", nil
	case "D&M":
		return "1000000", nil
	case "D|M":
		return "1010101", nil
	default:
		return "", errors.New("このコマンドに対応することができません")
	}
}

func (p Parser) Jump() (string, error) {
	isExistJump := strings.Contains(p.command, ";")
	if !isExistJump {
		return "000", nil
	}

	mnemonic := p.command[strings.Index(p.command, ";"):]

	switch mnemonic {
	case "JGT":
		return "001", nil
	case "JEQ":
		return "010", nil
	case "JGE":
		return "011", nil
	case "JLT":
		return "100", nil
	case "JNE":
		return "101", nil
	case "JLE":
		return "110", nil
	case "JMP":
		return "111", nil
	default:
		return "", errors.New("このコマンドに対応することができません")
	}
}
