// Code generated from lispie.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // lispie

import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type lispieParser struct {
	*antlr.BaseParser
}

var lispieParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func lispieParserInit() {
  staticData := &lispieParserStaticData
  staticData.literalNames = []string{
    "", "'('", "')'", "'self'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "", "OP", "KEYWORD", "INT", "FLOAT", "STRING", "IDENTIFIER", 
    "TYPE", "COMMENT", "WS",
  }
  staticData.ruleNames = []string{
    "program", "sexpr", "operator", "keyword", "args", "atom", "self",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 12, 55, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0, 
	1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 26, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 
	1, 32, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 5, 4, 40, 8, 4, 10, 4, 
	12, 4, 43, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 51, 8, 5, 1, 
	6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 57, 0, 17, 1, 0, 0, 
	0, 2, 31, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 35, 1, 0, 0, 0, 8, 41, 1, 0, 
	0, 0, 10, 50, 1, 0, 0, 0, 12, 52, 1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 14, 
	1, 0, 0, 0, 16, 19, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 
	18, 20, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 21, 5, 0, 0, 1, 21, 1, 1, 0, 
	0, 0, 22, 25, 5, 1, 0, 0, 23, 26, 3, 4, 2, 0, 24, 26, 3, 6, 3, 0, 25, 23, 
	1, 0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28, 3, 8, 4, 0, 
	28, 29, 5, 2, 0, 0, 29, 32, 1, 0, 0, 0, 30, 32, 3, 10, 5, 0, 31, 22, 1, 
	0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 3, 1, 0, 0, 0, 33, 34, 5, 4, 0, 0, 34, 
	5, 1, 0, 0, 0, 35, 36, 5, 5, 0, 0, 36, 7, 1, 0, 0, 0, 37, 40, 3, 10, 5, 
	0, 38, 40, 3, 2, 1, 0, 39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 43, 
	1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 9, 1, 0, 0, 0, 
	43, 41, 1, 0, 0, 0, 44, 51, 5, 6, 0, 0, 45, 51, 5, 7, 0, 0, 46, 51, 5, 
	8, 0, 0, 47, 51, 5, 9, 0, 0, 48, 51, 5, 10, 0, 0, 49, 51, 3, 12, 6, 0, 
	50, 44, 1, 0, 0, 0, 50, 45, 1, 0, 0, 0, 50, 46, 1, 0, 0, 0, 50, 47, 1, 
	0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 49, 1, 0, 0, 0, 51, 11, 1, 0, 0, 0, 52, 
	53, 5, 3, 0, 0, 53, 13, 1, 0, 0, 0, 6, 17, 25, 31, 39, 41, 50,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// lispieParserInit initializes any static state used to implement lispieParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewlispieParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LispieParserInit() {
  staticData := &lispieParserStaticData
  staticData.once.Do(lispieParserInit)
}

// NewlispieParser produces a new parser instance for the optional input antlr.TokenStream.
func NewlispieParser(input antlr.TokenStream) *lispieParser {
	LispieParserInit()
	this := new(lispieParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &lispieParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "lispie.g4"

	return this
}


// lispieParser tokens.
const (
	lispieParserEOF = antlr.TokenEOF
	lispieParserT__0 = 1
	lispieParserT__1 = 2
	lispieParserT__2 = 3
	lispieParserOP = 4
	lispieParserKEYWORD = 5
	lispieParserINT = 6
	lispieParserFLOAT = 7
	lispieParserSTRING = 8
	lispieParserIDENTIFIER = 9
	lispieParserTYPE = 10
	lispieParserCOMMENT = 11
	lispieParserWS = 12
)

// lispieParser rules.
const (
	lispieParserRULE_program = 0
	lispieParserRULE_sexpr = 1
	lispieParserRULE_operator = 2
	lispieParserRULE_keyword = 3
	lispieParserRULE_args = 4
	lispieParserRULE_atom = 5
	lispieParserRULE_self = 6
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllSexpr() []ISexprContext
	Sexpr(i int) ISexprContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lispieParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lispieParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(lispieParserEOF, 0)
}

func (s *ProgramContext) AllSexpr() []ISexprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISexprContext); ok {
			len++
		}
	}

	tst := make([]ISexprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISexprContext); ok {
			tst[i] = t.(ISexprContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Sexpr(i int) ISexprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISexprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISexprContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.ExitProgram(s)
	}
}




func (p *lispieParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lispieParserRULE_program)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 1994) != 0) {
		{
			p.SetState(14)
			p.Sexpr()
		}


		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
		p.Match(lispieParserEOF)
	}



	return localctx
}


// ISexprContext is an interface to support dynamic dispatch.
type ISexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSexprContext differentiates from other interfaces.
	IsSexprContext()
}

type SexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySexprContext() *SexprContext {
	var p = new(SexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lispieParserRULE_sexpr
	return p
}

func (*SexprContext) IsSexprContext() {}

func NewSexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexprContext {
	var p = new(SexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lispieParserRULE_sexpr

	return p
}

func (s *SexprContext) GetParser() antlr.Parser { return s.parser }

func (s *SexprContext) CopyFrom(ctx *SexprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type SexprOpContext struct {
	*SexprContext
}

func NewSexprOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SexprOpContext {
	var p = new(SexprOpContext)

	p.SexprContext = NewEmptySexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SexprContext))

	return p
}

func (s *SexprOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexprOpContext) Args() IArgsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *SexprOpContext) Operator() IOperatorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *SexprOpContext) Keyword() IKeywordContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}


func (s *SexprOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.EnterSexprOp(s)
	}
}

func (s *SexprOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.ExitSexprOp(s)
	}
}


type SexprAtomContext struct {
	*SexprContext
}

func NewSexprAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SexprAtomContext {
	var p = new(SexprAtomContext)

	p.SexprContext = NewEmptySexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SexprContext))

	return p
}

func (s *SexprAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexprAtomContext) Atom() IAtomContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}


func (s *SexprAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.EnterSexprAtom(s)
	}
}

func (s *SexprAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.ExitSexprAtom(s)
	}
}



func (p *lispieParser) Sexpr() (localctx ISexprContext) {
	this := p
	_ = this

	localctx = NewSexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lispieParserRULE_sexpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case lispieParserT__0:
		localctx = NewSexprOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(lispieParserT__0)
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case lispieParserOP:
			{
				p.SetState(23)
				p.Operator()
			}


		case lispieParserKEYWORD:
			{
				p.SetState(24)
				p.Keyword()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(27)
			p.Args()
		}
		{
			p.SetState(28)
			p.Match(lispieParserT__1)
		}


	case lispieParserT__2, lispieParserINT, lispieParserFLOAT, lispieParserSTRING, lispieParserIDENTIFIER, lispieParserTYPE:
		localctx = NewSexprAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Atom()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP() antlr.TerminalNode

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lispieParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lispieParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) OP() antlr.TerminalNode {
	return s.GetToken(lispieParserOP, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.ExitOperator(s)
	}
}




func (p *lispieParser) Operator() (localctx IOperatorContext) {
	this := p
	_ = this

	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lispieParserRULE_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(lispieParserOP)
	}



	return localctx
}


// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KEYWORD() antlr.TerminalNode

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lispieParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lispieParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(lispieParserKEYWORD, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.ExitKeyword(s)
	}
}




func (p *lispieParser) Keyword() (localctx IKeywordContext) {
	this := p
	_ = this

	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lispieParserRULE_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(lispieParserKEYWORD)
	}



	return localctx
}


// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAtom() []IAtomContext
	Atom(i int) IAtomContext
	AllSexpr() []ISexprContext
	Sexpr(i int) ISexprContext

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lispieParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lispieParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllAtom() []IAtomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtomContext); ok {
			len++
		}
	}

	tst := make([]IAtomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtomContext); ok {
			tst[i] = t.(IAtomContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Atom(i int) IAtomContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ArgsContext) AllSexpr() []ISexprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISexprContext); ok {
			len++
		}
	}

	tst := make([]ISexprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISexprContext); ok {
			tst[i] = t.(ISexprContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Sexpr(i int) ISexprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISexprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISexprContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.ExitArgs(s)
	}
}




func (p *lispieParser) Args() (localctx IArgsContext) {
	this := p
	_ = this

	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lispieParserRULE_args)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 1994) != 0) {
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(37)
				p.Atom()
			}


		case 2:
			{
				p.SetState(38)
				p.Sexpr()
			}

		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	Self() ISelfContext

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lispieParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lispieParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) INT() antlr.TerminalNode {
	return s.GetToken(lispieParserINT, 0)
}

func (s *AtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(lispieParserFLOAT, 0)
}

func (s *AtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(lispieParserSTRING, 0)
}

func (s *AtomContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(lispieParserIDENTIFIER, 0)
}

func (s *AtomContext) TYPE() antlr.TerminalNode {
	return s.GetToken(lispieParserTYPE, 0)
}

func (s *AtomContext) Self() ISelfContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelfContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelfContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.ExitAtom(s)
	}
}




func (p *lispieParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, lispieParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case lispieParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Match(lispieParserINT)
		}


	case lispieParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(lispieParserFLOAT)
		}


	case lispieParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.Match(lispieParserSTRING)
		}


	case lispieParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.Match(lispieParserIDENTIFIER)
		}


	case lispieParserTYPE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(48)
			p.Match(lispieParserTYPE)
		}


	case lispieParserT__2:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(49)
			p.Self()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ISelfContext is an interface to support dynamic dispatch.
type ISelfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSelfContext differentiates from other interfaces.
	IsSelfContext()
}

type SelfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelfContext() *SelfContext {
	var p = new(SelfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lispieParserRULE_self
	return p
}

func (*SelfContext) IsSelfContext() {}

func NewSelfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelfContext {
	var p = new(SelfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lispieParserRULE_self

	return p
}

func (s *SelfContext) GetParser() antlr.Parser { return s.parser }
func (s *SelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.EnterSelf(s)
	}
}

func (s *SelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lispieListener); ok {
		listenerT.ExitSelf(s)
	}
}




func (p *lispieParser) Self() (localctx ISelfContext) {
	this := p
	_ = this

	localctx = NewSelfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lispieParserRULE_self)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(lispieParserT__2)
	}



	return localctx
}


