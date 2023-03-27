// Code generated from lispie.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // lispie

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaselispieListener is a complete listener for a parse tree produced by lispieParser.
type BaselispieListener struct{}

var _ lispieListener = &BaselispieListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselispieListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselispieListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselispieListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselispieListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaselispieListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaselispieListener) ExitProgram(ctx *ProgramContext) {}

// EnterSexprOp is called when production sexprOp is entered.
func (s *BaselispieListener) EnterSexprOp(ctx *SexprOpContext) {}

// ExitSexprOp is called when production sexprOp is exited.
func (s *BaselispieListener) ExitSexprOp(ctx *SexprOpContext) {}

// EnterSexprAtom is called when production sexprAtom is entered.
func (s *BaselispieListener) EnterSexprAtom(ctx *SexprAtomContext) {}

// ExitSexprAtom is called when production sexprAtom is exited.
func (s *BaselispieListener) ExitSexprAtom(ctx *SexprAtomContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaselispieListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaselispieListener) ExitOperator(ctx *OperatorContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaselispieListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaselispieListener) ExitKeyword(ctx *KeywordContext) {}

// EnterArgs is called when production args is entered.
func (s *BaselispieListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaselispieListener) ExitArgs(ctx *ArgsContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaselispieListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaselispieListener) ExitAtom(ctx *AtomContext) {}

// EnterSelf is called when production self is entered.
func (s *BaselispieListener) EnterSelf(ctx *SelfContext) {}

// ExitSelf is called when production self is exited.
func (s *BaselispieListener) ExitSelf(ctx *SelfContext) {}
