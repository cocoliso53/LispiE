// Generated from /Users/cuau/Projects/LispiE/grammar/lispie.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class lispieParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, OP=5, KEYWORD=6, INT=7, ID=8, FLOAT=9, 
		STRING=10, IDENTIFIER=11, TYPE=12, COMMENT=13, WS=14;
	public static final int
		RULE_program = 0, RULE_contractName = 1, RULE_sexpr = 2, RULE_operator = 3, 
		RULE_keyword = 4, RULE_args = 5, RULE_atom = 6, RULE_self = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "contractName", "sexpr", "operator", "keyword", "args", "atom", 
			"self"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'#'", "'('", "')'", "'self'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "OP", "KEYWORD", "INT", "ID", "FLOAT", 
			"STRING", "IDENTIFIER", "TYPE", "COMMENT", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "lispie.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public lispieParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public ContractNameContext contractName() {
			return getRuleContext(ContractNameContext.class,0);
		}
		public TerminalNode EOF() { return getToken(lispieParser.EOF, 0); }
		public List<SexprContext> sexpr() {
			return getRuleContexts(SexprContext.class);
		}
		public SexprContext sexpr(int i) {
			return getRuleContext(SexprContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(16);
			contractName();
			setState(20);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__3) | (1L << INT) | (1L << ID) | (1L << FLOAT) | (1L << STRING) | (1L << IDENTIFIER) | (1L << TYPE))) != 0)) {
				{
				{
				setState(17);
				sexpr();
				}
				}
				setState(22);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(23);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ContractNameContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(lispieParser.ID, 0); }
		public ContractNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_contractName; }
	}

	public final ContractNameContext contractName() throws RecognitionException {
		ContractNameContext _localctx = new ContractNameContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_contractName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			match(T__0);
			setState(26);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SexprContext extends ParserRuleContext {
		public SexprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sexpr; }
	 
		public SexprContext() { }
		public void copyFrom(SexprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class SexprOpContext extends SexprContext {
		public ArgsContext args() {
			return getRuleContext(ArgsContext.class,0);
		}
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public KeywordContext keyword() {
			return getRuleContext(KeywordContext.class,0);
		}
		public SexprOpContext(SexprContext ctx) { copyFrom(ctx); }
	}
	public static class SexprAtomContext extends SexprContext {
		public AtomContext atom() {
			return getRuleContext(AtomContext.class,0);
		}
		public SexprAtomContext(SexprContext ctx) { copyFrom(ctx); }
	}

	public final SexprContext sexpr() throws RecognitionException {
		SexprContext _localctx = new SexprContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sexpr);
		try {
			setState(37);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
				_localctx = new SexprOpContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(28);
				match(T__1);
				setState(31);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case OP:
					{
					setState(29);
					operator();
					}
					break;
				case KEYWORD:
					{
					setState(30);
					keyword();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(33);
				args();
				setState(34);
				match(T__2);
				}
				break;
			case T__3:
			case INT:
			case ID:
			case FLOAT:
			case STRING:
			case IDENTIFIER:
			case TYPE:
				_localctx = new SexprAtomContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(36);
				atom();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OperatorContext extends ParserRuleContext {
		public TerminalNode OP() { return getToken(lispieParser.OP, 0); }
		public OperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator; }
	}

	public final OperatorContext operator() throws RecognitionException {
		OperatorContext _localctx = new OperatorContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			match(OP);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeywordContext extends ParserRuleContext {
		public TerminalNode KEYWORD() { return getToken(lispieParser.KEYWORD, 0); }
		public KeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword; }
	}

	public final KeywordContext keyword() throws RecognitionException {
		KeywordContext _localctx = new KeywordContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_keyword);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			match(KEYWORD);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgsContext extends ParserRuleContext {
		public List<AtomContext> atom() {
			return getRuleContexts(AtomContext.class);
		}
		public AtomContext atom(int i) {
			return getRuleContext(AtomContext.class,i);
		}
		public List<SexprContext> sexpr() {
			return getRuleContexts(SexprContext.class);
		}
		public SexprContext sexpr(int i) {
			return getRuleContext(SexprContext.class,i);
		}
		public ArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_args; }
	}

	public final ArgsContext args() throws RecognitionException {
		ArgsContext _localctx = new ArgsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_args);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__3) | (1L << INT) | (1L << ID) | (1L << FLOAT) | (1L << STRING) | (1L << IDENTIFIER) | (1L << TYPE))) != 0)) {
				{
				setState(45);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(43);
					atom();
					}
					break;
				case 2:
					{
					setState(44);
					sexpr();
					}
					break;
				}
				}
				setState(49);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AtomContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(lispieParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(lispieParser.FLOAT, 0); }
		public TerminalNode STRING() { return getToken(lispieParser.STRING, 0); }
		public TerminalNode IDENTIFIER() { return getToken(lispieParser.IDENTIFIER, 0); }
		public TerminalNode TYPE() { return getToken(lispieParser.TYPE, 0); }
		public TerminalNode ID() { return getToken(lispieParser.ID, 0); }
		public SelfContext self() {
			return getRuleContext(SelfContext.class,0);
		}
		public AtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atom; }
	}

	public final AtomContext atom() throws RecognitionException {
		AtomContext _localctx = new AtomContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_atom);
		try {
			setState(57);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(50);
				match(INT);
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(51);
				match(FLOAT);
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(52);
				match(STRING);
				}
				break;
			case IDENTIFIER:
				enterOuterAlt(_localctx, 4);
				{
				setState(53);
				match(IDENTIFIER);
				}
				break;
			case TYPE:
				enterOuterAlt(_localctx, 5);
				{
				setState(54);
				match(TYPE);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 6);
				{
				setState(55);
				match(ID);
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 7);
				{
				setState(56);
				self();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SelfContext extends ParserRuleContext {
		public SelfContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_self; }
	}

	public final SelfContext self() throws RecognitionException {
		SelfContext _localctx = new SelfContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_self);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\20@\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\3\2\7\2\25\n\2"+
		"\f\2\16\2\30\13\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\5\4\"\n\4\3\4\3\4\3"+
		"\4\3\4\5\4(\n\4\3\5\3\5\3\6\3\6\3\7\3\7\7\7\60\n\7\f\7\16\7\63\13\7\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b<\n\b\3\t\3\t\3\t\2\2\n\2\4\6\b\n\f\16\20"+
		"\2\2\2B\2\22\3\2\2\2\4\33\3\2\2\2\6\'\3\2\2\2\b)\3\2\2\2\n+\3\2\2\2\f"+
		"\61\3\2\2\2\16;\3\2\2\2\20=\3\2\2\2\22\26\5\4\3\2\23\25\5\6\4\2\24\23"+
		"\3\2\2\2\25\30\3\2\2\2\26\24\3\2\2\2\26\27\3\2\2\2\27\31\3\2\2\2\30\26"+
		"\3\2\2\2\31\32\7\2\2\3\32\3\3\2\2\2\33\34\7\3\2\2\34\35\7\n\2\2\35\5\3"+
		"\2\2\2\36!\7\4\2\2\37\"\5\b\5\2 \"\5\n\6\2!\37\3\2\2\2! \3\2\2\2\"#\3"+
		"\2\2\2#$\5\f\7\2$%\7\5\2\2%(\3\2\2\2&(\5\16\b\2\'\36\3\2\2\2\'&\3\2\2"+
		"\2(\7\3\2\2\2)*\7\7\2\2*\t\3\2\2\2+,\7\b\2\2,\13\3\2\2\2-\60\5\16\b\2"+
		".\60\5\6\4\2/-\3\2\2\2/.\3\2\2\2\60\63\3\2\2\2\61/\3\2\2\2\61\62\3\2\2"+
		"\2\62\r\3\2\2\2\63\61\3\2\2\2\64<\7\t\2\2\65<\7\13\2\2\66<\7\f\2\2\67"+
		"<\7\r\2\28<\7\16\2\29<\7\n\2\2:<\5\20\t\2;\64\3\2\2\2;\65\3\2\2\2;\66"+
		"\3\2\2\2;\67\3\2\2\2;8\3\2\2\2;9\3\2\2\2;:\3\2\2\2<\17\3\2\2\2=>\7\6\2"+
		"\2>\21\3\2\2\2\b\26!\'/\61;";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}