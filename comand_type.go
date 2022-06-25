package hack_assembler_go

type CommandType int

const (
	A_COMMAND CommandType = iota + 1
	C_COMMAND
	L_COMMAND
)

func (c CommandType) String() string {
	v := [...]string{"", "A_COMMAND", "C_COMMAND", "L_COMMAND"}
	return v[c]
}
