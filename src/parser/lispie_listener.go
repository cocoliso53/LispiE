// Code generated from lispie.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // lispie

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// lispieListener is a complete listener for a parse tree produced by lispieParser.
type lispieListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSexprOp is called when entering the sexprOp production.
	EnterSexprOp(c *SexprOpContext)

	// EnterSexprAtom is called when entering the sexprAtom production.
	EnterSexprAtom(c *SexprAtomContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterSelf is called when entering the self production.
	EnterSelf(c *SelfContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSexprOp is called when exiting the sexprOp production.
	ExitSexprOp(c *SexprOpContext)

	// ExitSexprAtom is called when exiting the sexprAtom production.
	ExitSexprAtom(c *SexprAtomContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitSelf is called when exiting the self production.
	ExitSelf(c *SelfContext)
}
