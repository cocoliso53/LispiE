package parser

import (
	"fmt"
	"strings"
)

type LispiEYulListener struct {
	BaselispieListener
	yulCode []string
}

func NewLispiEYulListener() *LispiEYulListener {
	return &LispiEYulListener{
		yulCode: []string{},
	}
}

func (l *LispiEYulListener) EnterProgram(ctx *ProgramContext) {
	l.yulCode = append(l.yulCode, "object \"Main\" {")
	l.yulCode = append(l.yulCode, "code {")
}

func (l *LispiEYulListener) ExitProgram(ctx *ProgramContext) {
	l.yulCode = append(l.yulCode, "}")
	l.yulCode = append(l.yulCode, "}")
}

func (l *LispiEYulListener) EnterSexprOp(ctx *SexprOpContext) {
	operator := ctx.Operator().GetText()
	switch operator {
	case "contract":
		contractName := ctx.Sexpr().GetChild(1).GetText()
		l.yulCode = append(l.yulCode, fmt.Sprintf("object \"%s\" {", contractName))
		l.yulCode = append(l.yulCode, "code {")
	case "function":
		functionName := ctx.Sexpr().GetChild(1).GetText()
		l.yulCode = append(l.yulCode, fmt.Sprintf("function %s() {", functionName))
	case "storage":
		varName := ctx.Sexpr().GetChild(1).GetText()
		initValue := ctx.Sexpr().GetChild(2).GetText()
		l.yulCode = append(l.yulCode, fmt.Sprintf("let %s := %s", varName, initValue))
	}
}

func (l *LispiEYulListener) ExitSexprOp(ctx *SexprOpContext) {
	operator := ctx.Operator().GetText()
	switch operator {
	case "contract":
		l.yulCode = append(l.yulCode, "}")
		l.yulCode = append(l.yulCode, "}")
	case "function":
		l.yulCode = append(l.yulCode, "}")
	}
}

func (l *LispiEYulListener) GetYulCode() string {
	return strings.Join(l.yulCode, "\n")
}
