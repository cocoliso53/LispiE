// Generated from /Users/cuau/Projects/LispiE/grammar/lispie.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class lispieLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, OP=4, KEYWORD=5, INT=6, FLOAT=7, STRING=8, IDENTIFIER=9, 
		TYPE=10, COMMENT=11, WS=12;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "OP", "KEYWORD", "INT", "FLOAT", "STRING", "IDENTIFIER", 
			"TYPE", "COMMENT", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'self'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, "OP", "KEYWORD", "INT", "FLOAT", "STRING", "IDENTIFIER", 
			"TYPE", "COMMENT", "WS"
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


	public lispieLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "lispie.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\16\u00eb\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5v\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6"+
		"\u008f\n\6\3\7\6\7\u0092\n\7\r\7\16\7\u0093\3\b\6\b\u0097\n\b\r\b\16\b"+
		"\u0098\3\b\3\b\6\b\u009d\n\b\r\b\16\b\u009e\3\t\3\t\7\t\u00a3\n\t\f\t"+
		"\16\t\u00a6\13\t\3\t\3\t\3\n\3\n\7\n\u00ac\n\n\f\n\16\n\u00af\13\n\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u00d6\n\13\3\f\3\f\3\f\3"+
		"\f\7\f\u00dc\n\f\f\f\16\f\u00df\13\f\3\f\3\f\3\f\3\f\3\r\6\r\u00e6\n\r"+
		"\r\r\16\r\u00e7\3\r\3\r\3\u00dd\2\16\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21"+
		"\n\23\13\25\f\27\r\31\16\3\2\b\6\2,-//\61\61>@\3\2\62;\5\2\13\f\17\17"+
		"$$\6\2&&C\\aac|\7\2&&\62;C\\aac|\5\2\13\f\17\17\"\"\2\u0109\2\3\3\2\2"+
		"\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3"+
		"\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2"+
		"\2\3\33\3\2\2\2\5\35\3\2\2\2\7\37\3\2\2\2\tu\3\2\2\2\13\u008e\3\2\2\2"+
		"\r\u0091\3\2\2\2\17\u0096\3\2\2\2\21\u00a0\3\2\2\2\23\u00a9\3\2\2\2\25"+
		"\u00d5\3\2\2\2\27\u00d7\3\2\2\2\31\u00e5\3\2\2\2\33\34\7*\2\2\34\4\3\2"+
		"\2\2\35\36\7+\2\2\36\6\3\2\2\2\37 \7u\2\2 !\7g\2\2!\"\7n\2\2\"#\7h\2\2"+
		"#\b\3\2\2\2$v\t\2\2\2%&\7c\2\2&\'\7p\2\2\'v\7f\2\2()\7q\2\2)v\7t\2\2*"+
		"+\7p\2\2+,\7q\2\2,v\7v\2\2-.\7k\2\2.v\7h\2\2/\60\7n\2\2\60\61\7g\2\2\61"+
		"v\7v\2\2\62\63\7f\2\2\63\64\7g\2\2\64\65\7h\2\2\65\66\7w\2\2\66v\7p\2"+
		"\2\678\7f\2\289\7g\2\29:\7h\2\2:;\7x\2\2;<\7c\2\2<v\7t\2\2=>\7f\2\2>?"+
		"\7g\2\2?@\7h\2\2@A\7e\2\2AB\7q\2\2BC\7p\2\2CD\7u\2\2Dv\7v\2\2EF\7f\2\2"+
		"FG\7g\2\2GH\7h\2\2HI\7k\2\2IJ\7p\2\2JK\7v\2\2KL\7g\2\2LM\7t\2\2MN\7h\2"+
		"\2NO\7c\2\2OP\7e\2\2Pv\7g\2\2QR\7h\2\2RS\7w\2\2ST\7p\2\2TU\7e\2\2UV\7"+
		"v\2\2VW\7k\2\2WX\7q\2\2Xv\7p\2\2YZ\7g\2\2Z[\7x\2\2[\\\7g\2\2\\]\7p\2\2"+
		"]v\7v\2\2^_\7g\2\2_`\7o\2\2`a\7k\2\2av\7v\2\2bc\7o\2\2cd\7c\2\2dv\7r\2"+
		"\2ef\7h\2\2fg\7k\2\2gh\7n\2\2hi\7v\2\2ij\7g\2\2jv\7t\2\2kl\7h\2\2lm\7"+
		"q\2\2mv\7t\2\2no\7y\2\2op\7j\2\2pq\7k\2\2qr\7n\2\2rv\7g\2\2st\7<\2\2t"+
		"v\5\13\6\2u$\3\2\2\2u%\3\2\2\2u(\3\2\2\2u*\3\2\2\2u-\3\2\2\2u/\3\2\2\2"+
		"u\62\3\2\2\2u\67\3\2\2\2u=\3\2\2\2uE\3\2\2\2uQ\3\2\2\2uY\3\2\2\2u^\3\2"+
		"\2\2ub\3\2\2\2ue\3\2\2\2uk\3\2\2\2un\3\2\2\2us\3\2\2\2v\n\3\2\2\2wx\7"+
		"g\2\2xy\7z\2\2yz\7v\2\2z{\7g\2\2{|\7t\2\2|}\7p\2\2}~\7c\2\2~\u008f\7n"+
		"\2\2\177\u0080\7k\2\2\u0080\u0081\7p\2\2\u0081\u0082\7v\2\2\u0082\u0083"+
		"\7g\2\2\u0083\u0084\7t\2\2\u0084\u0085\7p\2\2\u0085\u0086\7c\2\2\u0086"+
		"\u008f\7n\2\2\u0087\u0088\7r\2\2\u0088\u0089\7c\2\2\u0089\u008a\7{\2\2"+
		"\u008a\u008b\7c\2\2\u008b\u008c\7d\2\2\u008c\u008d\7n\2\2\u008d\u008f"+
		"\7g\2\2\u008ew\3\2\2\2\u008e\177\3\2\2\2\u008e\u0087\3\2\2\2\u008f\f\3"+
		"\2\2\2\u0090\u0092\t\3\2\2\u0091\u0090\3\2\2\2\u0092\u0093\3\2\2\2\u0093"+
		"\u0091\3\2\2\2\u0093\u0094\3\2\2\2\u0094\16\3\2\2\2\u0095\u0097\t\3\2"+
		"\2\u0096\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098\u0096\3\2\2\2\u0098\u0099"+
		"\3\2\2\2\u0099\u009a\3\2\2\2\u009a\u009c\7\60\2\2\u009b\u009d\t\3\2\2"+
		"\u009c\u009b\3\2\2\2\u009d\u009e\3\2\2\2\u009e\u009c\3\2\2\2\u009e\u009f"+
		"\3\2\2\2\u009f\20\3\2\2\2\u00a0\u00a4\7$\2\2\u00a1\u00a3\n\4\2\2\u00a2"+
		"\u00a1\3\2\2\2\u00a3\u00a6\3\2\2\2\u00a4\u00a2\3\2\2\2\u00a4\u00a5\3\2"+
		"\2\2\u00a5\u00a7\3\2\2\2\u00a6\u00a4\3\2\2\2\u00a7\u00a8\7$\2\2\u00a8"+
		"\22\3\2\2\2\u00a9\u00ad\t\5\2\2\u00aa\u00ac\t\6\2\2\u00ab\u00aa\3\2\2"+
		"\2\u00ac\u00af\3\2\2\2\u00ad\u00ab\3\2\2\2\u00ad\u00ae\3\2\2\2\u00ae\24"+
		"\3\2\2\2\u00af\u00ad\3\2\2\2\u00b0\u00b1\7w\2\2\u00b1\u00b2\7k\2\2\u00b2"+
		"\u00b3\7p\2\2\u00b3\u00b4\7v\2\2\u00b4\u00b5\7\64\2\2\u00b5\u00b6\7\67"+
		"\2\2\u00b6\u00d6\78\2\2\u00b7\u00b8\7k\2\2\u00b8\u00b9\7p\2\2\u00b9\u00ba"+
		"\7v\2\2\u00ba\u00bb\7\64\2\2\u00bb\u00bc\7\67\2\2\u00bc\u00d6\78\2\2\u00bd"+
		"\u00be\7c\2\2\u00be\u00bf\7f\2\2\u00bf\u00c0\7f\2\2\u00c0\u00c1\7t\2\2"+
		"\u00c1\u00c2\7g\2\2\u00c2\u00c3\7u\2\2\u00c3\u00d6\7u\2\2\u00c4\u00c5"+
		"\7d\2\2\u00c5\u00c6\7q\2\2\u00c6\u00c7\7q\2\2\u00c7\u00d6\7n\2\2\u00c8"+
		"\u00c9\7d\2\2\u00c9\u00ca\7{\2\2\u00ca\u00cb\7v\2\2\u00cb\u00cc\7g\2\2"+
		"\u00cc\u00cd\7u\2\2\u00cd\u00ce\7\65\2\2\u00ce\u00d6\7\64\2\2\u00cf\u00d0"+
		"\7u\2\2\u00d0\u00d1\7v\2\2\u00d1\u00d2\7t\2\2\u00d2\u00d3\7k\2\2\u00d3"+
		"\u00d4\7p\2\2\u00d4\u00d6\7i\2\2\u00d5\u00b0\3\2\2\2\u00d5\u00b7\3\2\2"+
		"\2\u00d5\u00bd\3\2\2\2\u00d5\u00c4\3\2\2\2\u00d5\u00c8\3\2\2\2\u00d5\u00cf"+
		"\3\2\2\2\u00d6\26\3\2\2\2\u00d7\u00d8\7=\2\2\u00d8\u00d9\7=\2\2\u00d9"+
		"\u00dd\3\2\2\2\u00da\u00dc\13\2\2\2\u00db\u00da\3\2\2\2\u00dc\u00df\3"+
		"\2\2\2\u00dd\u00de\3\2\2\2\u00dd\u00db\3\2\2\2\u00de\u00e0\3\2\2\2\u00df"+
		"\u00dd\3\2\2\2\u00e0\u00e1\7\f\2\2\u00e1\u00e2\3\2\2\2\u00e2\u00e3\b\f"+
		"\2\2\u00e3\30\3\2\2\2\u00e4\u00e6\t\7\2\2\u00e5\u00e4\3\2\2\2\u00e6\u00e7"+
		"\3\2\2\2\u00e7\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9"+
		"\u00ea\b\r\2\2\u00ea\32\3\2\2\2\r\2u\u008e\u0093\u0098\u009e\u00a4\u00ad"+
		"\u00d5\u00dd\u00e7\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}