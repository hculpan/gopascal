package ast

const (
	STRING  = iota
	INTEGER = iota
	FLOAT   = iota
)

type Argument struct {
	ArgumentType int
	StringValue  string
	IntValue     int
	FloatValue   float32
}
