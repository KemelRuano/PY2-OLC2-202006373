// Generated from c:\Users\LENOVO\Desktop\OLC2_PY1\backend-app\grammar\Lexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class LexerLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, STRING=3, BOOL=4, CHARACTER=5, IGUAL=6, DOSPUNTOS=7, OPEN_PAREN=8, 
		CLOSE_PAREN=9, INCREMENTO=10, DECREMENTO=11, OPEN_BRACKET=12, CLOSE_BRACKET=13, 
		OPEN_SQUARE_BRACKET=14, CLOSE_SQUARE_BRACKET=15, COMA=16, PUNTO=17, GUION_BAJO=18, 
		MAS=19, MENOS=20, MUL=21, DIV=22, MOD=23, MAYOR=24, MENOR=25, MAYOR_IGUAL=26, 
		MENOR_IGUAL=27, IGUAL_IGUAL=28, DIFERENTE=29, AND=30, OR=31, NOT=32, VAR=33, 
		LET=34, PRINT=35, IF=36, ELSE=37, SWITCH=38, CASE=39, DEFAULT=40, WHILE=41, 
		FOR=42, IN=43, BREAK=44, CONTINUE=45, RETURN=46, TRESPUNTOS=47, GUARD=48, 
		APPEND=49, COUNT=50, REMOVE=51, REMOVELAST=52, AT=53, ISEMPTY=54, FUNC=55, 
		INOUT=56, FLECHA=57, REFERENCIA=58, STRUCT=59, MUTATING=60, SELF=61, NUMBER=62, 
		VALORBOOL=63, NULO=64, STRING_SINTAX=65, ESCAPE_SEQUENCE=66, CARACTER=67, 
		SINVALOR=68, ID=69, COMMENT=70, COMMENT_MULT=71, WS=72;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "IGUAL", "DOSPUNTOS", 
			"OPEN_PAREN", "CLOSE_PAREN", "INCREMENTO", "DECREMENTO", "OPEN_BRACKET", 
			"CLOSE_BRACKET", "OPEN_SQUARE_BRACKET", "CLOSE_SQUARE_BRACKET", "COMA", 
			"PUNTO", "GUION_BAJO", "MAS", "MENOS", "MUL", "DIV", "MOD", "MAYOR", 
			"MENOR", "MAYOR_IGUAL", "MENOR_IGUAL", "IGUAL_IGUAL", "DIFERENTE", "AND", 
			"OR", "NOT", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", 
			"WHILE", "FOR", "IN", "BREAK", "CONTINUE", "RETURN", "TRESPUNTOS", "GUARD", 
			"APPEND", "COUNT", "REMOVE", "REMOVELAST", "AT", "ISEMPTY", "FUNC", "INOUT", 
			"FLECHA", "REFERENCIA", "STRUCT", "MUTATING", "SELF", "NUMBER", "VALORBOOL", 
			"NULO", "STRING_SINTAX", "ESCAPE_SEQUENCE", "CARACTER", "SINVALOR", "ID", 
			"COMMENT", "COMMENT_MULT", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'='", 
			"':'", "'('", "')'", "'+='", "'-='", "'{'", "'}'", "'['", "']'", "','", 
			"'.'", "'_'", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'>='", 
			"'<='", "'=='", "'!='", "'&&'", "'||'", "'!'", "'var'", "'let'", "'print'", 
			"'if'", "'else'", "'switch'", "'case'", "'default'", "'while'", "'for'", 
			"'in'", "'break'", "'continue'", "'return'", "'...'", "'guard'", "'append'", 
			"'count'", "'remove'", "'removeLast'", "'at:'", "'IsEmpty'", "'func'", 
			"'inout'", "'->'", "'&'", "'struct'", "'mutating'", "'self'", null, null, 
			"'nil'", null, null, null, "'?'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "IGUAL", "DOSPUNTOS", 
			"OPEN_PAREN", "CLOSE_PAREN", "INCREMENTO", "DECREMENTO", "OPEN_BRACKET", 
			"CLOSE_BRACKET", "OPEN_SQUARE_BRACKET", "CLOSE_SQUARE_BRACKET", "COMA", 
			"PUNTO", "GUION_BAJO", "MAS", "MENOS", "MUL", "DIV", "MOD", "MAYOR", 
			"MENOR", "MAYOR_IGUAL", "MENOR_IGUAL", "IGUAL_IGUAL", "DIFERENTE", "AND", 
			"OR", "NOT", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", 
			"WHILE", "FOR", "IN", "BREAK", "CONTINUE", "RETURN", "TRESPUNTOS", "GUARD", 
			"APPEND", "COUNT", "REMOVE", "REMOVELAST", "AT", "ISEMPTY", "FUNC", "INOUT", 
			"FLECHA", "REFERENCIA", "STRUCT", "MUTATING", "SELF", "NUMBER", "VALORBOOL", 
			"NULO", "STRING_SINTAX", "ESCAPE_SEQUENCE", "CARACTER", "SINVALOR", "ID", 
			"COMMENT", "COMMENT_MULT", "WS"
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


	public LexerLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Lexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2J\u01f1\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7"+
		"\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\16\3\16"+
		"\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25"+
		"\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\33\3\34"+
		"\3\34\3\34\3\35\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3!\3"+
		"!\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3%\3%\3%\3&\3&\3&\3&\3"+
		"&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)\3)\3"+
		"*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3,\3,\3,\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3"+
		".\3.\3.\3.\3.\3/\3/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\61\3\61\3\61"+
		"\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\63"+
		"\3\63\3\63\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\65"+
		"\3\65\3\65\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\67\3\67\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\38\38\38\38\38\39\39\39\39\39\39\3:\3:\3:\3;\3;\3"+
		"<\3<\3<\3<\3<\3<\3<\3=\3=\3=\3=\3=\3=\3=\3=\3=\3>\3>\3>\3>\3>\3?\5?\u019a"+
		"\n?\3?\6?\u019d\n?\r?\16?\u019e\3?\3?\6?\u01a3\n?\r?\16?\u01a4\5?\u01a7"+
		"\n?\3@\3@\3@\3@\3@\3@\3@\3@\3@\5@\u01b2\n@\3A\3A\3A\3A\3B\3B\3B\7B\u01bb"+
		"\nB\fB\16B\u01be\13B\3B\3B\3C\3C\3C\3D\3D\3D\3D\3E\3E\3F\3F\7F\u01cd\n"+
		"F\fF\16F\u01d0\13F\3G\3G\3G\3G\7G\u01d6\nG\fG\16G\u01d9\13G\3G\3G\3H\3"+
		"H\3H\3H\7H\u01e1\nH\fH\16H\u01e4\13H\3H\3H\3H\3H\3H\3I\6I\u01ec\nI\rI"+
		"\16I\u01ed\3I\3I\3\u01e2\2J\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25"+
		"\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32"+
		"\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a"+
		"\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087"+
		"E\u0089F\u008bG\u008dH\u008fI\u0091J\3\2\n\3\2\62;\6\2\f\f\17\17$$))\5"+
		"\2ppttvv\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\4\2\f\f\17\17\5\2\13\f\16\17"+
		"\"\"\2\u01fb\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2"+
		"\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2"+
		"\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2"+
		"\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2"+
		"\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3"+
		"\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2"+
		"\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2"+
		"S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3"+
		"\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2"+
		"\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2"+
		"y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083"+
		"\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2"+
		"\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\3\u0093\3\2\2\2\5\u0097"+
		"\3\2\2\2\7\u009d\3\2\2\2\t\u00a4\3\2\2\2\13\u00a9\3\2\2\2\r\u00b3\3\2"+
		"\2\2\17\u00b5\3\2\2\2\21\u00b7\3\2\2\2\23\u00b9\3\2\2\2\25\u00bb\3\2\2"+
		"\2\27\u00be\3\2\2\2\31\u00c1\3\2\2\2\33\u00c3\3\2\2\2\35\u00c5\3\2\2\2"+
		"\37\u00c7\3\2\2\2!\u00c9\3\2\2\2#\u00cb\3\2\2\2%\u00cd\3\2\2\2\'\u00cf"+
		"\3\2\2\2)\u00d1\3\2\2\2+\u00d3\3\2\2\2-\u00d5\3\2\2\2/\u00d7\3\2\2\2\61"+
		"\u00d9\3\2\2\2\63\u00db\3\2\2\2\65\u00dd\3\2\2\2\67\u00e0\3\2\2\29\u00e3"+
		"\3\2\2\2;\u00e6\3\2\2\2=\u00e9\3\2\2\2?\u00ec\3\2\2\2A\u00ef\3\2\2\2C"+
		"\u00f1\3\2\2\2E\u00f5\3\2\2\2G\u00f9\3\2\2\2I\u00ff\3\2\2\2K\u0102\3\2"+
		"\2\2M\u0107\3\2\2\2O\u010e\3\2\2\2Q\u0113\3\2\2\2S\u011b\3\2\2\2U\u0121"+
		"\3\2\2\2W\u0125\3\2\2\2Y\u0128\3\2\2\2[\u012e\3\2\2\2]\u0137\3\2\2\2_"+
		"\u013e\3\2\2\2a\u0142\3\2\2\2c\u0148\3\2\2\2e\u014f\3\2\2\2g\u0155\3\2"+
		"\2\2i\u015c\3\2\2\2k\u0167\3\2\2\2m\u016b\3\2\2\2o\u0173\3\2\2\2q\u0178"+
		"\3\2\2\2s\u017e\3\2\2\2u\u0181\3\2\2\2w\u0183\3\2\2\2y\u018a\3\2\2\2{"+
		"\u0193\3\2\2\2}\u0199\3\2\2\2\177\u01b1\3\2\2\2\u0081\u01b3\3\2\2\2\u0083"+
		"\u01b7\3\2\2\2\u0085\u01c1\3\2\2\2\u0087\u01c4\3\2\2\2\u0089\u01c8\3\2"+
		"\2\2\u008b\u01ca\3\2\2\2\u008d\u01d1\3\2\2\2\u008f\u01dc\3\2\2\2\u0091"+
		"\u01eb\3\2\2\2\u0093\u0094\7K\2\2\u0094\u0095\7p\2\2\u0095\u0096\7v\2"+
		"\2\u0096\4\3\2\2\2\u0097\u0098\7H\2\2\u0098\u0099\7n\2\2\u0099\u009a\7"+
		"q\2\2\u009a\u009b\7c\2\2\u009b\u009c\7v\2\2\u009c\6\3\2\2\2\u009d\u009e"+
		"\7U\2\2\u009e\u009f\7v\2\2\u009f\u00a0\7t\2\2\u00a0\u00a1\7k\2\2\u00a1"+
		"\u00a2\7p\2\2\u00a2\u00a3\7i\2\2\u00a3\b\3\2\2\2\u00a4\u00a5\7D\2\2\u00a5"+
		"\u00a6\7q\2\2\u00a6\u00a7\7q\2\2\u00a7\u00a8\7n\2\2\u00a8\n\3\2\2\2\u00a9"+
		"\u00aa\7E\2\2\u00aa\u00ab\7j\2\2\u00ab\u00ac\7c\2\2\u00ac\u00ad\7t\2\2"+
		"\u00ad\u00ae\7c\2\2\u00ae\u00af\7e\2\2\u00af\u00b0\7v\2\2\u00b0\u00b1"+
		"\7g\2\2\u00b1\u00b2\7t\2\2\u00b2\f\3\2\2\2\u00b3\u00b4\7?\2\2\u00b4\16"+
		"\3\2\2\2\u00b5\u00b6\7<\2\2\u00b6\20\3\2\2\2\u00b7\u00b8\7*\2\2\u00b8"+
		"\22\3\2\2\2\u00b9\u00ba\7+\2\2\u00ba\24\3\2\2\2\u00bb\u00bc\7-\2\2\u00bc"+
		"\u00bd\7?\2\2\u00bd\26\3\2\2\2\u00be\u00bf\7/\2\2\u00bf\u00c0\7?\2\2\u00c0"+
		"\30\3\2\2\2\u00c1\u00c2\7}\2\2\u00c2\32\3\2\2\2\u00c3\u00c4\7\177\2\2"+
		"\u00c4\34\3\2\2\2\u00c5\u00c6\7]\2\2\u00c6\36\3\2\2\2\u00c7\u00c8\7_\2"+
		"\2\u00c8 \3\2\2\2\u00c9\u00ca\7.\2\2\u00ca\"\3\2\2\2\u00cb\u00cc\7\60"+
		"\2\2\u00cc$\3\2\2\2\u00cd\u00ce\7a\2\2\u00ce&\3\2\2\2\u00cf\u00d0\7-\2"+
		"\2\u00d0(\3\2\2\2\u00d1\u00d2\7/\2\2\u00d2*\3\2\2\2\u00d3\u00d4\7,\2\2"+
		"\u00d4,\3\2\2\2\u00d5\u00d6\7\61\2\2\u00d6.\3\2\2\2\u00d7\u00d8\7\'\2"+
		"\2\u00d8\60\3\2\2\2\u00d9\u00da\7@\2\2\u00da\62\3\2\2\2\u00db\u00dc\7"+
		">\2\2\u00dc\64\3\2\2\2\u00dd\u00de\7@\2\2\u00de\u00df\7?\2\2\u00df\66"+
		"\3\2\2\2\u00e0\u00e1\7>\2\2\u00e1\u00e2\7?\2\2\u00e28\3\2\2\2\u00e3\u00e4"+
		"\7?\2\2\u00e4\u00e5\7?\2\2\u00e5:\3\2\2\2\u00e6\u00e7\7#\2\2\u00e7\u00e8"+
		"\7?\2\2\u00e8<\3\2\2\2\u00e9\u00ea\7(\2\2\u00ea\u00eb\7(\2\2\u00eb>\3"+
		"\2\2\2\u00ec\u00ed\7~\2\2\u00ed\u00ee\7~\2\2\u00ee@\3\2\2\2\u00ef\u00f0"+
		"\7#\2\2\u00f0B\3\2\2\2\u00f1\u00f2\7x\2\2\u00f2\u00f3\7c\2\2\u00f3\u00f4"+
		"\7t\2\2\u00f4D\3\2\2\2\u00f5\u00f6\7n\2\2\u00f6\u00f7\7g\2\2\u00f7\u00f8"+
		"\7v\2\2\u00f8F\3\2\2\2\u00f9\u00fa\7r\2\2\u00fa\u00fb\7t\2\2\u00fb\u00fc"+
		"\7k\2\2\u00fc\u00fd\7p\2\2\u00fd\u00fe\7v\2\2\u00feH\3\2\2\2\u00ff\u0100"+
		"\7k\2\2\u0100\u0101\7h\2\2\u0101J\3\2\2\2\u0102\u0103\7g\2\2\u0103\u0104"+
		"\7n\2\2\u0104\u0105\7u\2\2\u0105\u0106\7g\2\2\u0106L\3\2\2\2\u0107\u0108"+
		"\7u\2\2\u0108\u0109\7y\2\2\u0109\u010a\7k\2\2\u010a\u010b\7v\2\2\u010b"+
		"\u010c\7e\2\2\u010c\u010d\7j\2\2\u010dN\3\2\2\2\u010e\u010f\7e\2\2\u010f"+
		"\u0110\7c\2\2\u0110\u0111\7u\2\2\u0111\u0112\7g\2\2\u0112P\3\2\2\2\u0113"+
		"\u0114\7f\2\2\u0114\u0115\7g\2\2\u0115\u0116\7h\2\2\u0116\u0117\7c\2\2"+
		"\u0117\u0118\7w\2\2\u0118\u0119\7n\2\2\u0119\u011a\7v\2\2\u011aR\3\2\2"+
		"\2\u011b\u011c\7y\2\2\u011c\u011d\7j\2\2\u011d\u011e\7k\2\2\u011e\u011f"+
		"\7n\2\2\u011f\u0120\7g\2\2\u0120T\3\2\2\2\u0121\u0122\7h\2\2\u0122\u0123"+
		"\7q\2\2\u0123\u0124\7t\2\2\u0124V\3\2\2\2\u0125\u0126\7k\2\2\u0126\u0127"+
		"\7p\2\2\u0127X\3\2\2\2\u0128\u0129\7d\2\2\u0129\u012a\7t\2\2\u012a\u012b"+
		"\7g\2\2\u012b\u012c\7c\2\2\u012c\u012d\7m\2\2\u012dZ\3\2\2\2\u012e\u012f"+
		"\7e\2\2\u012f\u0130\7q\2\2\u0130\u0131\7p\2\2\u0131\u0132\7v\2\2\u0132"+
		"\u0133\7k\2\2\u0133\u0134\7p\2\2\u0134\u0135\7w\2\2\u0135\u0136\7g\2\2"+
		"\u0136\\\3\2\2\2\u0137\u0138\7t\2\2\u0138\u0139\7g\2\2\u0139\u013a\7v"+
		"\2\2\u013a\u013b\7w\2\2\u013b\u013c\7t\2\2\u013c\u013d\7p\2\2\u013d^\3"+
		"\2\2\2\u013e\u013f\7\60\2\2\u013f\u0140\7\60\2\2\u0140\u0141\7\60\2\2"+
		"\u0141`\3\2\2\2\u0142\u0143\7i\2\2\u0143\u0144\7w\2\2\u0144\u0145\7c\2"+
		"\2\u0145\u0146\7t\2\2\u0146\u0147\7f\2\2\u0147b\3\2\2\2\u0148\u0149\7"+
		"c\2\2\u0149\u014a\7r\2\2\u014a\u014b\7r\2\2\u014b\u014c\7g\2\2\u014c\u014d"+
		"\7p\2\2\u014d\u014e\7f\2\2\u014ed\3\2\2\2\u014f\u0150\7e\2\2\u0150\u0151"+
		"\7q\2\2\u0151\u0152\7w\2\2\u0152\u0153\7p\2\2\u0153\u0154\7v\2\2\u0154"+
		"f\3\2\2\2\u0155\u0156\7t\2\2\u0156\u0157\7g\2\2\u0157\u0158\7o\2\2\u0158"+
		"\u0159\7q\2\2\u0159\u015a\7x\2\2\u015a\u015b\7g\2\2\u015bh\3\2\2\2\u015c"+
		"\u015d\7t\2\2\u015d\u015e\7g\2\2\u015e\u015f\7o\2\2\u015f\u0160\7q\2\2"+
		"\u0160\u0161\7x\2\2\u0161\u0162\7g\2\2\u0162\u0163\7N\2\2\u0163\u0164"+
		"\7c\2\2\u0164\u0165\7u\2\2\u0165\u0166\7v\2\2\u0166j\3\2\2\2\u0167\u0168"+
		"\7c\2\2\u0168\u0169\7v\2\2\u0169\u016a\7<\2\2\u016al\3\2\2\2\u016b\u016c"+
		"\7K\2\2\u016c\u016d\7u\2\2\u016d\u016e\7G\2\2\u016e\u016f\7o\2\2\u016f"+
		"\u0170\7r\2\2\u0170\u0171\7v\2\2\u0171\u0172\7{\2\2\u0172n\3\2\2\2\u0173"+
		"\u0174\7h\2\2\u0174\u0175\7w\2\2\u0175\u0176\7p\2\2\u0176\u0177\7e\2\2"+
		"\u0177p\3\2\2\2\u0178\u0179\7k\2\2\u0179\u017a\7p\2\2\u017a\u017b\7q\2"+
		"\2\u017b\u017c\7w\2\2\u017c\u017d\7v\2\2\u017dr\3\2\2\2\u017e\u017f\7"+
		"/\2\2\u017f\u0180\7@\2\2\u0180t\3\2\2\2\u0181\u0182\7(\2\2\u0182v\3\2"+
		"\2\2\u0183\u0184\7u\2\2\u0184\u0185\7v\2\2\u0185\u0186\7t\2\2\u0186\u0187"+
		"\7w\2\2\u0187\u0188\7e\2\2\u0188\u0189\7v\2\2\u0189x\3\2\2\2\u018a\u018b"+
		"\7o\2\2\u018b\u018c\7w\2\2\u018c\u018d\7v\2\2\u018d\u018e\7c\2\2\u018e"+
		"\u018f\7v\2\2\u018f\u0190\7k\2\2\u0190\u0191\7p\2\2\u0191\u0192\7i\2\2"+
		"\u0192z\3\2\2\2\u0193\u0194\7u\2\2\u0194\u0195\7g\2\2\u0195\u0196\7n\2"+
		"\2\u0196\u0197\7h\2\2\u0197|\3\2\2\2\u0198\u019a\7/\2\2\u0199\u0198\3"+
		"\2\2\2\u0199\u019a\3\2\2\2\u019a\u019c\3\2\2\2\u019b\u019d\t\2\2\2\u019c"+
		"\u019b\3\2\2\2\u019d\u019e\3\2\2\2\u019e\u019c\3\2\2\2\u019e\u019f\3\2"+
		"\2\2\u019f\u01a6\3\2\2\2\u01a0\u01a2\7\60\2\2\u01a1\u01a3\t\2\2\2\u01a2"+
		"\u01a1\3\2\2\2\u01a3\u01a4\3\2\2\2\u01a4\u01a2\3\2\2\2\u01a4\u01a5\3\2"+
		"\2\2\u01a5\u01a7\3\2\2\2\u01a6\u01a0\3\2\2\2\u01a6\u01a7\3\2\2\2\u01a7"+
		"~\3\2\2\2\u01a8\u01a9\7v\2\2\u01a9\u01aa\7t\2\2\u01aa\u01ab\7w\2\2\u01ab"+
		"\u01b2\7g\2\2\u01ac\u01ad\7h\2\2\u01ad\u01ae\7c\2\2\u01ae\u01af\7n\2\2"+
		"\u01af\u01b0\7u\2\2\u01b0\u01b2\7g\2\2\u01b1\u01a8\3\2\2\2\u01b1\u01ac"+
		"\3\2\2\2\u01b2\u0080\3\2\2\2\u01b3\u01b4\7p\2\2\u01b4\u01b5\7k\2\2\u01b5"+
		"\u01b6\7n\2\2\u01b6\u0082\3\2\2\2\u01b7\u01bc\7$\2\2\u01b8\u01bb\5\u0085"+
		"C\2\u01b9\u01bb\n\3\2\2\u01ba\u01b8\3\2\2\2\u01ba\u01b9\3\2\2\2\u01bb"+
		"\u01be\3\2\2\2\u01bc\u01ba\3\2\2\2\u01bc\u01bd\3\2\2\2\u01bd\u01bf\3\2"+
		"\2\2\u01be\u01bc\3\2\2\2\u01bf\u01c0\7$\2\2\u01c0\u0084\3\2\2\2\u01c1"+
		"\u01c2\7^\2\2\u01c2\u01c3\t\4\2\2\u01c3\u0086\3\2\2\2\u01c4\u01c5\7$\2"+
		"\2\u01c5\u01c6\n\5\2\2\u01c6\u01c7\7$\2\2\u01c7\u0088\3\2\2\2\u01c8\u01c9"+
		"\7A\2\2\u01c9\u008a\3\2\2\2\u01ca\u01ce\t\6\2\2\u01cb\u01cd\t\7\2\2\u01cc"+
		"\u01cb\3\2\2\2\u01cd\u01d0\3\2\2\2\u01ce\u01cc\3\2\2\2\u01ce\u01cf\3\2"+
		"\2\2\u01cf\u008c\3\2\2\2\u01d0\u01ce\3\2\2\2\u01d1\u01d2\7\61\2\2\u01d2"+
		"\u01d3\7\61\2\2\u01d3\u01d7\3\2\2\2\u01d4\u01d6\n\b\2\2\u01d5\u01d4\3"+
		"\2\2\2\u01d6\u01d9\3\2\2\2\u01d7\u01d5\3\2\2\2\u01d7\u01d8\3\2\2\2\u01d8"+
		"\u01da\3\2\2\2\u01d9\u01d7\3\2\2\2\u01da\u01db\bG\2\2\u01db\u008e\3\2"+
		"\2\2\u01dc\u01dd\7\61\2\2\u01dd\u01de\7,\2\2\u01de\u01e2\3\2\2\2\u01df"+
		"\u01e1\13\2\2\2\u01e0\u01df\3\2\2\2\u01e1\u01e4\3\2\2\2\u01e2\u01e3\3"+
		"\2\2\2\u01e2\u01e0\3\2\2\2\u01e3\u01e5\3\2\2\2\u01e4\u01e2\3\2\2\2\u01e5"+
		"\u01e6\7,\2\2\u01e6\u01e7\7\61\2\2\u01e7\u01e8\3\2\2\2\u01e8\u01e9\bH"+
		"\2\2\u01e9\u0090\3\2\2\2\u01ea\u01ec\t\t\2\2\u01eb\u01ea\3\2\2\2\u01ec"+
		"\u01ed\3\2\2\2\u01ed\u01eb\3\2\2\2\u01ed\u01ee\3\2\2\2\u01ee\u01ef\3\2"+
		"\2\2\u01ef\u01f0\bI\2\2\u01f0\u0092\3\2\2\2\16\2\u0199\u019e\u01a4\u01a6"+
		"\u01b1\u01ba\u01bc\u01ce\u01d7\u01e2\u01ed\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}