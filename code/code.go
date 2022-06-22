package code

import "errors"

type Code struct {
	mnemonic string
}

func (c Code) OutPutDestBinary() (string, error) {
	switch c.mnemonic {
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

func (c Code) OutputCompBinary() (string, error) {
	switch c.mnemonic {
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

func (c Code) OutputJumpBinary() (string, error) {
	switch c.mnemonic {
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
