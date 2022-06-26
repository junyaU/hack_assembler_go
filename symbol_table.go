package hack_assembler_go

import (
	"strconv"
)

const _variableBaseAddress = 16

//　Predefined addresses
const (
	SP     = 0
	LCL    = 1
	ARG    = 2
	THIS   = 3
	THAT   = 4
	R0     = 0
	R1     = 1
	R2     = 2
	R3     = 3
	R4     = 4
	R5     = 5
	R6     = 6
	R7     = 7
	R8     = 8
	R9     = 9
	R10    = 10
	R11    = 11
	R12    = 12
	R13    = 13
	R14    = 14
	R15    = 15
	SCREEN = 16384
	KBD    = 24576
)

type SymbolTable struct {
	table         map[string]int
	variableCount int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		table:         make(map[string]int),
		variableCount: 0,
	}
}

func (t *SymbolTable) AddEntry(address int, symbol string, commandType CommandType) {
	switch commandType {
	case L_COMMAND:
		if _, ok := t.table[symbol]; ok {
			t.table[symbol] = address
			return
		}
	case A_COMMAND:
		address = t.assignACommandAddress(symbol)
		if _, ok := t.table[symbol]; ok {
			return
		}
	}

	t.table[symbol] = address
}

func (t SymbolTable) Contains(symbol string) bool {
	_, ok := t.table[symbol]
	return ok
}

func (t SymbolTable) GetAddress(symbol string) int {
	return t.table[symbol]
}

func (t SymbolTable) Table() map[string]int {
	return t.table
}

func (t SymbolTable) IsNonNumericSymbol(symbol string) bool {
	_, err := strconv.Atoi(symbol)
	return err != nil
}

func (t *SymbolTable) assignACommandAddress(symbol string) int {
	switch symbol {
	case "SP":
		return SP
	case "LCL":
		return LCL
	case "ARG":
		return ARG
	case "THIS":
		return THIS
	case "THAT":
		return THAT
	case "R0":
		return R0
	case "R1":
		return R1
	case "R2":
		return R2
	case "R3":
		return R3
	case "R4":
		return R4
	case "R5":
		return R5
	case "R6":
		return R6
	case "R7":
		return R7
	case "R8":
		return R8
	case "R9":
		return R9
	case "R10":
		return R10
	case "R11":
		return R11
	case "R12":
		return R12
	case "R13":
		return R13
	case "R14":
		return R14
	case "R15":
		return R15
	case "SCREEN":
		return SCREEN
	case "KBD":
		return KBD

	default:
		t.variableCount++
		variableAddress := t.variableCount - 2

		return variableAddress + _variableBaseAddress
	}
}
