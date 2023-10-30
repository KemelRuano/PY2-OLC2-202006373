// Generated from c:/Users/LENOVO/Desktop/PY2-OLC2-202006373/OLC2_PY2/backend-app/gramatica/Sintax.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class SintaxParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

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
	public static final int
		RULE_inicio = 0, RULE_lista = 1, RULE_lista_proceso = 2, RULE_comentarios = 3, 
		RULE_retornar = 4, RULE_declaracion = 5, RULE_dec_tipo_valor = 6, RULE_dec_valor = 7, 
		RULE_asignacion = 8, RULE_asignacionVariable = 9, RULE_asignacionVector = 10, 
		RULE_subasig = 11, RULE_print = 12, RULE_if = 13, RULE_elseif_ = 14, RULE_else = 15, 
		RULE_switch = 16, RULE_caso = 17, RULE_default = 18, RULE_while = 19, 
		RULE_for = 20, RULE_guard = 21, RULE_dec_vector = 22, RULE_dec_vector_V_C = 23, 
		RULE_subVC = 24, RULE_funcvectorList = 25, RULE_tipevector = 26, RULE_bloque = 27, 
		RULE_control = 28, RULE_tipo = 29, RULE_expresion = 30, RULE_funciones = 31, 
		RULE_parametros = 32, RULE_existeExInt = 33, RULE_tipofuncion = 34, RULE_tipoinout = 35, 
		RULE_funcLLamada = 36, RULE_parametrosLLamada = 37, RULE_estructs = 38, 
		RULE_lista_struct = 39, RULE_funcionestruct = 40;
	private static String[] makeRuleNames() {
		return new String[] {
			"inicio", "lista", "lista_proceso", "comentarios", "retornar", "declaracion", 
			"dec_tipo_valor", "dec_valor", "asignacion", "asignacionVariable", "asignacionVector", 
			"subasig", "print", "if", "elseif_", "else", "switch", "caso", "default", 
			"while", "for", "guard", "dec_vector", "dec_vector_V_C", "subVC", "funcvectorList", 
			"tipevector", "bloque", "control", "tipo", "expresion", "funciones", 
			"parametros", "existeExInt", "tipofuncion", "tipoinout", "funcLLamada", 
			"parametrosLLamada", "estructs", "lista_struct", "funcionestruct"
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

	@Override
	public String getGrammarFileName() { return "Sintax.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SintaxParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InicioContext extends ParserRuleContext {
		public ListaContext lista() {
			return getRuleContext(ListaContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SintaxParser.EOF, 0); }
		public InicioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inicio; }
	}

	public final InicioContext inicio() throws RecognitionException {
		InicioContext _localctx = new InicioContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_inicio);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			lista();
			setState(83);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ListaContext extends ParserRuleContext {
		public List<Lista_procesoContext> lista_proceso() {
			return getRuleContexts(Lista_procesoContext.class);
		}
		public Lista_procesoContext lista_proceso(int i) {
			return getRuleContext(Lista_procesoContext.class,i);
		}
		public List<FuncionesContext> funciones() {
			return getRuleContexts(FuncionesContext.class);
		}
		public FuncionesContext funciones(int i) {
			return getRuleContext(FuncionesContext.class,i);
		}
		public ListaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista; }
	}

	public final ListaContext lista() throws RecognitionException {
		ListaContext _localctx = new ListaContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_lista);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & 481107688239L) != 0)) {
				{
				setState(87);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case VAR:
				case LET:
				case PRINT:
				case IF:
				case SWITCH:
				case WHILE:
				case FOR:
				case BREAK:
				case CONTINUE:
				case RETURN:
				case GUARD:
				case STRUCT:
				case ID:
				case COMMENT:
				case COMMENT_MULT:
					{
					setState(85);
					lista_proceso();
					}
					break;
				case FUNC:
					{
					setState(86);
					funciones();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(91);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Lista_procesoContext extends ParserRuleContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public IfContext if_() {
			return getRuleContext(IfContext.class,0);
		}
		public SwitchContext switch_() {
			return getRuleContext(SwitchContext.class,0);
		}
		public WhileContext while_() {
			return getRuleContext(WhileContext.class,0);
		}
		public ForContext for_() {
			return getRuleContext(ForContext.class,0);
		}
		public GuardContext guard() {
			return getRuleContext(GuardContext.class,0);
		}
		public FuncvectorListContext funcvectorList() {
			return getRuleContext(FuncvectorListContext.class,0);
		}
		public FuncLLamadaContext funcLLamada() {
			return getRuleContext(FuncLLamadaContext.class,0);
		}
		public RetornarContext retornar() {
			return getRuleContext(RetornarContext.class,0);
		}
		public EstructsContext estructs() {
			return getRuleContext(EstructsContext.class,0);
		}
		public ComentariosContext comentarios() {
			return getRuleContext(ComentariosContext.class,0);
		}
		public Lista_procesoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_proceso; }
	}

	public final Lista_procesoContext lista_proceso() throws RecognitionException {
		Lista_procesoContext _localctx = new Lista_procesoContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_lista_proceso);
		try {
			setState(105);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(92);
				declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(93);
				asignacion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(94);
				print();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(95);
				if_();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(96);
				switch_();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(97);
				while_();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(98);
				for_();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(99);
				guard();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(100);
				funcvectorList();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(101);
				funcLLamada();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(102);
				retornar();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(103);
				estructs();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(104);
				comentarios();
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class ComentariosContext extends ParserRuleContext {
		public TerminalNode COMMENT() { return getToken(SintaxParser.COMMENT, 0); }
		public TerminalNode COMMENT_MULT() { return getToken(SintaxParser.COMMENT_MULT, 0); }
		public ComentariosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comentarios; }
	}

	public final ComentariosContext comentarios() throws RecognitionException {
		ComentariosContext _localctx = new ComentariosContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_comentarios);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			_la = _input.LA(1);
			if ( !(_la==COMMENT || _la==COMMENT_MULT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	@SuppressWarnings("CheckReturnValue")
	public static class RetornarContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(SintaxParser.RETURN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(SintaxParser.BREAK, 0); }
		public TerminalNode CONTINUE() { return getToken(SintaxParser.CONTINUE, 0); }
		public RetornarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_retornar; }
	}

	public final RetornarContext retornar() throws RecognitionException {
		RetornarContext _localctx = new RetornarContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_retornar);
		try {
			setState(115);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN:
				enterOuterAlt(_localctx, 1);
				{
				setState(109);
				match(RETURN);
				setState(111);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(110);
					expresion(0);
					}
					break;
				}
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 2);
				{
				setState(113);
				match(BREAK);
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(114);
				match(CONTINUE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeclaracionContext extends ParserRuleContext {
		public Dec_tipo_valorContext dec_tipo_valor() {
			return getRuleContext(Dec_tipo_valorContext.class,0);
		}
		public Dec_valorContext dec_valor() {
			return getRuleContext(Dec_valorContext.class,0);
		}
		public Dec_vectorContext dec_vector() {
			return getRuleContext(Dec_vectorContext.class,0);
		}
		public Dec_vector_V_CContext dec_vector_V_C() {
			return getRuleContext(Dec_vector_V_CContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_declaracion);
		try {
			setState(121);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(117);
				dec_tipo_valor();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(118);
				dec_valor();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(119);
				dec_vector();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(120);
				dec_vector_V_C();
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class Dec_tipo_valorContext extends ParserRuleContext {
		public Dec_tipo_valorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_tipo_valor; }
	 
		public Dec_tipo_valorContext() { }
		public void copyFrom(Dec_tipo_valorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclaracionTipoContext extends Dec_tipo_valorContext {
		public Token op;
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SintaxParser.DOSPUNTOS, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode VAR() { return getToken(SintaxParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SintaxParser.LET, 0); }
		public TerminalNode IGUAL() { return getToken(SintaxParser.IGUAL, 0); }
		public DeclaracionTipoContext(Dec_tipo_valorContext ctx) { copyFrom(ctx); }
	}

	public final Dec_tipo_valorContext dec_tipo_valor() throws RecognitionException {
		Dec_tipo_valorContext _localctx = new Dec_tipo_valorContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_dec_tipo_valor);
		int _la;
		try {
			_localctx = new DeclaracionTipoContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			((DeclaracionTipoContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
				((DeclaracionTipoContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(124);
			match(ID);
			setState(125);
			match(DOSPUNTOS);
			setState(126);
			tipo();
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IGUAL) {
				{
				setState(127);
				match(IGUAL);
				}
			}

			setState(130);
			expresion(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Dec_valorContext extends ParserRuleContext {
		public Dec_valorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_valor; }
	 
		public Dec_valorContext() { }
		public void copyFrom(Dec_valorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclaracionTipoImplicitoContext extends Dec_valorContext {
		public Token op;
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(SintaxParser.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode VAR() { return getToken(SintaxParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SintaxParser.LET, 0); }
		public DeclaracionTipoImplicitoContext(Dec_valorContext ctx) { copyFrom(ctx); }
	}

	public final Dec_valorContext dec_valor() throws RecognitionException {
		Dec_valorContext _localctx = new Dec_valorContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_dec_valor);
		int _la;
		try {
			_localctx = new DeclaracionTipoImplicitoContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			((DeclaracionTipoImplicitoContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
				((DeclaracionTipoImplicitoContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(133);
			match(ID);
			setState(134);
			match(IGUAL);
			setState(135);
			expresion(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionContext extends ParserRuleContext {
		public AsignacionVariableContext asignacionVariable() {
			return getRuleContext(AsignacionVariableContext.class,0);
		}
		public AsignacionVectorContext asignacionVector() {
			return getRuleContext(AsignacionVectorContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_asignacion);
		try {
			setState(139);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(137);
				asignacionVariable();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(138);
				asignacionVector();
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionVariableContext extends ParserRuleContext {
		public Token op;
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode INCREMENTO() { return getToken(SintaxParser.INCREMENTO, 0); }
		public TerminalNode DECREMENTO() { return getToken(SintaxParser.DECREMENTO, 0); }
		public TerminalNode IGUAL() { return getToken(SintaxParser.IGUAL, 0); }
		public AsignacionVariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacionVariable; }
	}

	public final AsignacionVariableContext asignacionVariable() throws RecognitionException {
		AsignacionVariableContext _localctx = new AsignacionVariableContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_asignacionVariable);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			match(ID);
			setState(142);
			((AsignacionVariableContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 3136L) != 0)) ) {
				((AsignacionVariableContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(143);
			expresion(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionVectorContext extends ParserRuleContext {
		public Token op;
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode OPEN_SQUARE_BRACKET() { return getToken(SintaxParser.OPEN_SQUARE_BRACKET, 0); }
		public SubasigContext subasig() {
			return getRuleContext(SubasigContext.class,0);
		}
		public TerminalNode CLOSE_SQUARE_BRACKET() { return getToken(SintaxParser.CLOSE_SQUARE_BRACKET, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(SintaxParser.IGUAL, 0); }
		public TerminalNode INCREMENTO() { return getToken(SintaxParser.INCREMENTO, 0); }
		public AsignacionVectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacionVector; }
	}

	public final AsignacionVectorContext asignacionVector() throws RecognitionException {
		AsignacionVectorContext _localctx = new AsignacionVectorContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_asignacionVector);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(145);
			match(ID);
			setState(146);
			match(OPEN_SQUARE_BRACKET);
			setState(147);
			subasig();
			setState(148);
			match(CLOSE_SQUARE_BRACKET);
			setState(149);
			((AsignacionVectorContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==IGUAL || _la==INCREMENTO) ) {
				((AsignacionVectorContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(150);
			expresion(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SubasigContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(SintaxParser.NUMBER, 0); }
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public SubasigContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subasig; }
	}

	public final SubasigContext subasig() throws RecognitionException {
		SubasigContext _localctx = new SubasigContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_subasig);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			_la = _input.LA(1);
			if ( !(_la==NUMBER || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrintContext extends ParserRuleContext {
		public TerminalNode PRINT() { return getToken(SintaxParser.PRINT, 0); }
		public TerminalNode OPEN_PAREN() { return getToken(SintaxParser.OPEN_PAREN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode CLOSE_PAREN() { return getToken(SintaxParser.CLOSE_PAREN, 0); }
		public List<TerminalNode> COMA() { return getTokens(SintaxParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SintaxParser.COMA, i);
		}
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			match(PRINT);
			setState(155);
			match(OPEN_PAREN);
			setState(156);
			expresion(0);
			setState(161);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(157);
				match(COMA);
				setState(158);
				expresion(0);
				}
				}
				setState(163);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(164);
			match(CLOSE_PAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IfContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(SintaxParser.IF, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public List<Elseif_Context> elseif_() {
			return getRuleContexts(Elseif_Context.class);
		}
		public Elseif_Context elseif_(int i) {
			return getRuleContext(Elseif_Context.class,i);
		}
		public ElseContext else_() {
			return getRuleContext(ElseContext.class,0);
		}
		public IfContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if; }
	}

	public final IfContext if_() throws RecognitionException {
		IfContext _localctx = new IfContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_if);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(166);
			match(IF);
			setState(167);
			expresion(0);
			setState(168);
			bloque();
			setState(172);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(169);
					elseif_();
					}
					} 
				}
				setState(174);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			setState(176);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(175);
				else_();
				}
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

	@SuppressWarnings("CheckReturnValue")
	public static class Elseif_Context extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(SintaxParser.ELSE, 0); }
		public TerminalNode IF() { return getToken(SintaxParser.IF, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public Elseif_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseif_; }
	}

	public final Elseif_Context elseif_() throws RecognitionException {
		Elseif_Context _localctx = new Elseif_Context(_ctx, getState());
		enterRule(_localctx, 28, RULE_elseif_);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			match(ELSE);
			setState(179);
			match(IF);
			setState(180);
			expresion(0);
			setState(181);
			bloque();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ElseContext extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(SintaxParser.ELSE, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public ElseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else; }
	}

	public final ElseContext else_() throws RecognitionException {
		ElseContext _localctx = new ElseContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_else);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(ELSE);
			setState(184);
			bloque();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(SintaxParser.SWITCH, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode OPEN_BRACKET() { return getToken(SintaxParser.OPEN_BRACKET, 0); }
		public TerminalNode CLOSE_BRACKET() { return getToken(SintaxParser.CLOSE_BRACKET, 0); }
		public List<CasoContext> caso() {
			return getRuleContexts(CasoContext.class);
		}
		public CasoContext caso(int i) {
			return getRuleContext(CasoContext.class,i);
		}
		public DefaultContext default_() {
			return getRuleContext(DefaultContext.class,0);
		}
		public SwitchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch; }
	}

	public final SwitchContext switch_() throws RecognitionException {
		SwitchContext _localctx = new SwitchContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_switch);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(186);
			match(SWITCH);
			setState(187);
			expresion(0);
			setState(188);
			match(OPEN_BRACKET);
			setState(190); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(189);
				caso();
				}
				}
				setState(192); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(195);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(194);
				default_();
				}
			}

			setState(197);
			match(CLOSE_BRACKET);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CasoContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(SintaxParser.CASE, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode DOSPUNTOS() { return getToken(SintaxParser.DOSPUNTOS, 0); }
		public List<Lista_procesoContext> lista_proceso() {
			return getRuleContexts(Lista_procesoContext.class);
		}
		public Lista_procesoContext lista_proceso(int i) {
			return getRuleContext(Lista_procesoContext.class,i);
		}
		public CasoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caso; }
	}

	public final CasoContext caso() throws RecognitionException {
		CasoContext _localctx = new CasoContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_caso);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(199);
			match(CASE);
			setState(200);
			expresion(0);
			setState(201);
			match(DOSPUNTOS);
			setState(205);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & 481103493935L) != 0)) {
				{
				{
				setState(202);
				lista_proceso();
				}
				}
				setState(207);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DefaultContext extends ParserRuleContext {
		public TerminalNode DEFAULT() { return getToken(SintaxParser.DEFAULT, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SintaxParser.DOSPUNTOS, 0); }
		public List<Lista_procesoContext> lista_proceso() {
			return getRuleContexts(Lista_procesoContext.class);
		}
		public Lista_procesoContext lista_proceso(int i) {
			return getRuleContext(Lista_procesoContext.class,i);
		}
		public DefaultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default; }
	}

	public final DefaultContext default_() throws RecognitionException {
		DefaultContext _localctx = new DefaultContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_default);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(DEFAULT);
			setState(209);
			match(DOSPUNTOS);
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & 481103493935L) != 0)) {
				{
				{
				setState(210);
				lista_proceso();
				}
				}
				setState(215);
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

	@SuppressWarnings("CheckReturnValue")
	public static class WhileContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(SintaxParser.WHILE, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public WhileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while; }
	}

	public final WhileContext while_() throws RecognitionException {
		WhileContext _localctx = new WhileContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_while);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(216);
			match(WHILE);
			setState(217);
			expresion(0);
			setState(218);
			bloque();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ForContext extends ParserRuleContext {
		public ExpresionContext left;
		public ExpresionContext right;
		public TerminalNode FOR() { return getToken(SintaxParser.FOR, 0); }
		public TerminalNode IN() { return getToken(SintaxParser.IN, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode GUION_BAJO() { return getToken(SintaxParser.GUION_BAJO, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode TRESPUNTOS() { return getToken(SintaxParser.TRESPUNTOS, 0); }
		public ForContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for; }
	}

	public final ForContext for_() throws RecognitionException {
		ForContext _localctx = new ForContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_for);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(220);
			match(FOR);
			setState(221);
			_la = _input.LA(1);
			if ( !(_la==GUION_BAJO || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(222);
			match(IN);
			setState(223);
			((ForContext)_localctx).left = expresion(0);
			setState(226);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TRESPUNTOS) {
				{
				setState(224);
				match(TRESPUNTOS);
				setState(225);
				((ForContext)_localctx).right = expresion(0);
				}
			}

			setState(228);
			bloque();
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

	@SuppressWarnings("CheckReturnValue")
	public static class GuardContext extends ParserRuleContext {
		public TerminalNode GUARD() { return getToken(SintaxParser.GUARD, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SintaxParser.ELSE, 0); }
		public TerminalNode OPEN_BRACKET() { return getToken(SintaxParser.OPEN_BRACKET, 0); }
		public TerminalNode CLOSE_BRACKET() { return getToken(SintaxParser.CLOSE_BRACKET, 0); }
		public List<Lista_procesoContext> lista_proceso() {
			return getRuleContexts(Lista_procesoContext.class);
		}
		public Lista_procesoContext lista_proceso(int i) {
			return getRuleContext(Lista_procesoContext.class,i);
		}
		public GuardContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guard; }
	}

	public final GuardContext guard() throws RecognitionException {
		GuardContext _localctx = new GuardContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_guard);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(230);
			match(GUARD);
			setState(231);
			expresion(0);
			setState(232);
			match(ELSE);
			setState(233);
			match(OPEN_BRACKET);
			setState(237);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & 481103493935L) != 0)) {
				{
				{
				setState(234);
				lista_proceso();
				}
				}
				setState(239);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(240);
			match(CLOSE_BRACKET);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Dec_vectorContext extends ParserRuleContext {
		public Token op;
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SintaxParser.DOSPUNTOS, 0); }
		public List<TerminalNode> OPEN_SQUARE_BRACKET() { return getTokens(SintaxParser.OPEN_SQUARE_BRACKET); }
		public TerminalNode OPEN_SQUARE_BRACKET(int i) {
			return getToken(SintaxParser.OPEN_SQUARE_BRACKET, i);
		}
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public List<TerminalNode> CLOSE_SQUARE_BRACKET() { return getTokens(SintaxParser.CLOSE_SQUARE_BRACKET); }
		public TerminalNode CLOSE_SQUARE_BRACKET(int i) {
			return getToken(SintaxParser.CLOSE_SQUARE_BRACKET, i);
		}
		public TerminalNode IGUAL() { return getToken(SintaxParser.IGUAL, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode VAR() { return getToken(SintaxParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SintaxParser.LET, 0); }
		public List<TerminalNode> COMA() { return getTokens(SintaxParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SintaxParser.COMA, i);
		}
		public Dec_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_vector; }
	}

	public final Dec_vectorContext dec_vector() throws RecognitionException {
		Dec_vectorContext _localctx = new Dec_vectorContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_dec_vector);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
			((Dec_vectorContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
				((Dec_vectorContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(243);
			match(ID);
			setState(244);
			match(DOSPUNTOS);
			setState(245);
			match(OPEN_SQUARE_BRACKET);
			setState(246);
			tipo();
			setState(247);
			match(CLOSE_SQUARE_BRACKET);
			setState(248);
			match(IGUAL);
			setState(249);
			match(OPEN_SQUARE_BRACKET);
			setState(250);
			expresion(0);
			setState(255);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(251);
				match(COMA);
				setState(252);
				expresion(0);
				}
				}
				setState(257);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(258);
			match(CLOSE_SQUARE_BRACKET);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Dec_vector_V_CContext extends ParserRuleContext {
		public Token op;
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SintaxParser.DOSPUNTOS, 0); }
		public TerminalNode OPEN_SQUARE_BRACKET() { return getToken(SintaxParser.OPEN_SQUARE_BRACKET, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CLOSE_SQUARE_BRACKET() { return getToken(SintaxParser.CLOSE_SQUARE_BRACKET, 0); }
		public TerminalNode IGUAL() { return getToken(SintaxParser.IGUAL, 0); }
		public SubVCContext subVC() {
			return getRuleContext(SubVCContext.class,0);
		}
		public TerminalNode VAR() { return getToken(SintaxParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SintaxParser.LET, 0); }
		public Dec_vector_V_CContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_vector_V_C; }
	}

	public final Dec_vector_V_CContext dec_vector_V_C() throws RecognitionException {
		Dec_vector_V_CContext _localctx = new Dec_vector_V_CContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_dec_vector_V_C);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			((Dec_vector_V_CContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
				((Dec_vector_V_CContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(261);
			match(ID);
			setState(262);
			match(DOSPUNTOS);
			setState(263);
			match(OPEN_SQUARE_BRACKET);
			setState(264);
			tipo();
			setState(265);
			match(CLOSE_SQUARE_BRACKET);
			setState(266);
			match(IGUAL);
			setState(267);
			subVC();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SubVCContext extends ParserRuleContext {
		public TerminalNode OPEN_SQUARE_BRACKET() { return getToken(SintaxParser.OPEN_SQUARE_BRACKET, 0); }
		public TerminalNode CLOSE_SQUARE_BRACKET() { return getToken(SintaxParser.CLOSE_SQUARE_BRACKET, 0); }
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public SubVCContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subVC; }
	}

	public final SubVCContext subVC() throws RecognitionException {
		SubVCContext _localctx = new SubVCContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_subVC);
		try {
			setState(272);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OPEN_SQUARE_BRACKET:
				enterOuterAlt(_localctx, 1);
				{
				setState(269);
				match(OPEN_SQUARE_BRACKET);
				setState(270);
				match(CLOSE_SQUARE_BRACKET);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(271);
				match(ID);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncvectorListContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SintaxParser.PUNTO, 0); }
		public TipevectorContext tipevector() {
			return getRuleContext(TipevectorContext.class,0);
		}
		public TerminalNode OPEN_PAREN() { return getToken(SintaxParser.OPEN_PAREN, 0); }
		public TerminalNode CLOSE_PAREN() { return getToken(SintaxParser.CLOSE_PAREN, 0); }
		public TerminalNode AT() { return getToken(SintaxParser.AT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public FuncvectorListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcvectorList; }
	}

	public final FuncvectorListContext funcvectorList() throws RecognitionException {
		FuncvectorListContext _localctx = new FuncvectorListContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_funcvectorList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(274);
			match(ID);
			setState(275);
			match(PUNTO);
			setState(276);
			tipevector();
			setState(277);
			match(OPEN_PAREN);
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AT) {
				{
				setState(278);
				match(AT);
				}
			}

			setState(282);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & -2305843004918726338L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 51L) != 0)) {
				{
				setState(281);
				expresion(0);
				}
			}

			setState(284);
			match(CLOSE_PAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TipevectorContext extends ParserRuleContext {
		public TerminalNode APPEND() { return getToken(SintaxParser.APPEND, 0); }
		public TerminalNode REMOVELAST() { return getToken(SintaxParser.REMOVELAST, 0); }
		public TerminalNode REMOVE() { return getToken(SintaxParser.REMOVE, 0); }
		public TipevectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipevector; }
	}

	public final TipevectorContext tipevector() throws RecognitionException {
		TipevectorContext _localctx = new TipevectorContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_tipevector);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 7318349394477056L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	@SuppressWarnings("CheckReturnValue")
	public static class BloqueContext extends ParserRuleContext {
		public TerminalNode OPEN_BRACKET() { return getToken(SintaxParser.OPEN_BRACKET, 0); }
		public TerminalNode CLOSE_BRACKET() { return getToken(SintaxParser.CLOSE_BRACKET, 0); }
		public List<Lista_procesoContext> lista_proceso() {
			return getRuleContexts(Lista_procesoContext.class);
		}
		public Lista_procesoContext lista_proceso(int i) {
			return getRuleContext(Lista_procesoContext.class,i);
		}
		public BloqueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloque; }
	}

	public final BloqueContext bloque() throws RecognitionException {
		BloqueContext _localctx = new BloqueContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_bloque);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(288);
			match(OPEN_BRACKET);
			setState(292);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & 481103493935L) != 0)) {
				{
				{
				setState(289);
				lista_proceso();
				}
				}
				setState(294);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(295);
			match(CLOSE_BRACKET);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ControlContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(SintaxParser.RETURN, 0); }
		public TerminalNode BREAK() { return getToken(SintaxParser.BREAK, 0); }
		public TerminalNode CONTINUE() { return getToken(SintaxParser.CONTINUE, 0); }
		public ControlContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control; }
	}

	public final ControlContext control() throws RecognitionException {
		ControlContext _localctx = new ControlContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_control);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(297);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 123145302310912L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TipoContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(SintaxParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SintaxParser.FLOAT, 0); }
		public TerminalNode STRING() { return getToken(SintaxParser.STRING, 0); }
		public TerminalNode BOOL() { return getToken(SintaxParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(SintaxParser.CHARACTER, 0); }
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_tipo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(299);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 62L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpresionContext extends ParserRuleContext {
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	 
		public ExpresionContext() { }
		public void copyFrom(ExpresionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorCountContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SintaxParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(SintaxParser.COUNT, 0); }
		public VectorCountContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OpLogicoContext extends ExpresionContext {
		public ExpresionContext left;
		public Token op;
		public ExpresionContext right;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode AND() { return getToken(SintaxParser.AND, 0); }
		public TerminalNode OR() { return getToken(SintaxParser.OR, 0); }
		public OpLogicoContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SinValorContext extends ExpresionContext {
		public TerminalNode SINVALOR() { return getToken(SintaxParser.SINVALOR, 0); }
		public SinValorContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorAsignacionContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode OPEN_SQUARE_BRACKET() { return getToken(SintaxParser.OPEN_SQUARE_BRACKET, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode CLOSE_SQUARE_BRACKET() { return getToken(SintaxParser.CLOSE_SQUARE_BRACKET, 0); }
		public VectorAsignacionContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OpAritmeticoContext extends ExpresionContext {
		public ExpresionContext left;
		public Token op;
		public ExpresionContext right;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode MUL() { return getToken(SintaxParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SintaxParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(SintaxParser.MOD, 0); }
		public TerminalNode MAS() { return getToken(SintaxParser.MAS, 0); }
		public TerminalNode MENOS() { return getToken(SintaxParser.MENOS, 0); }
		public OpAritmeticoContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NumberValorContext extends ExpresionContext {
		public TerminalNode NUMBER() { return getToken(SintaxParser.NUMBER, 0); }
		public NumberValorContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorsimpleContext extends ExpresionContext {
		public TerminalNode NULO() { return getToken(SintaxParser.NULO, 0); }
		public ValorsimpleContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CasteoContext extends ExpresionContext {
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode OPEN_PAREN() { return getToken(SintaxParser.OPEN_PAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode CLOSE_PAREN() { return getToken(SintaxParser.CLOSE_PAREN, 0); }
		public CasteoContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorVacioContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SintaxParser.PUNTO, 0); }
		public TerminalNode ISEMPTY() { return getToken(SintaxParser.ISEMPTY, 0); }
		public VectorVacioContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParentexprContext extends ExpresionContext {
		public TerminalNode OPEN_PAREN() { return getToken(SintaxParser.OPEN_PAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode CLOSE_PAREN() { return getToken(SintaxParser.CLOSE_PAREN, 0); }
		public ParentexprContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SelfCALLContext extends ExpresionContext {
		public TerminalNode SELF() { return getToken(SintaxParser.SELF, 0); }
		public TerminalNode PUNTO() { return getToken(SintaxParser.PUNTO, 0); }
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(SintaxParser.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public SelfCALLContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringValorContext extends ExpresionContext {
		public TerminalNode STRING_SINTAX() { return getToken(SintaxParser.STRING_SINTAX, 0); }
		public StringValorContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OpRelacionalContext extends ExpresionContext {
		public ExpresionContext left;
		public Token op;
		public ExpresionContext right;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode MAYOR() { return getToken(SintaxParser.MAYOR, 0); }
		public TerminalNode MENOR() { return getToken(SintaxParser.MENOR, 0); }
		public TerminalNode MAYOR_IGUAL() { return getToken(SintaxParser.MAYOR_IGUAL, 0); }
		public TerminalNode MENOR_IGUAL() { return getToken(SintaxParser.MENOR_IGUAL, 0); }
		public OpRelacionalContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BoolValorContext extends ExpresionContext {
		public TerminalNode VALORBOOL() { return getToken(SintaxParser.VALORBOOL, 0); }
		public BoolValorContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintLLamadaContext extends ExpresionContext {
		public FuncLLamadaContext funcLLamada() {
			return getRuleContext(FuncLLamadaContext.class,0);
		}
		public PrintLLamadaContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OpLogicoNotContext extends ExpresionContext {
		public TerminalNode NOT() { return getToken(SintaxParser.NOT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public OpLogicoNotContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdValorContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public IdValorContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OpComparativoContext extends ExpresionContext {
		public ExpresionContext left;
		public Token op;
		public ExpresionContext right;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode IGUAL_IGUAL() { return getToken(SintaxParser.IGUAL_IGUAL, 0); }
		public TerminalNode DIFERENTE() { return getToken(SintaxParser.DIFERENTE, 0); }
		public OpComparativoContext(ExpresionContext ctx) { copyFrom(ctx); }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(336);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				{
				_localctx = new OpLogicoNotContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(302);
				match(NOT);
				setState(303);
				expresion(14);
				}
				break;
			case 2:
				{
				_localctx = new ParentexprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(304);
				match(OPEN_PAREN);
				setState(305);
				expresion(0);
				setState(306);
				match(CLOSE_PAREN);
				}
				break;
			case 3:
				{
				_localctx = new VectorVacioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(308);
				match(ID);
				setState(309);
				match(PUNTO);
				setState(310);
				match(ISEMPTY);
				}
				break;
			case 4:
				{
				_localctx = new VectorCountContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(311);
				match(ID);
				setState(312);
				match(PUNTO);
				setState(313);
				match(COUNT);
				}
				break;
			case 5:
				{
				_localctx = new VectorAsignacionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(314);
				match(ID);
				setState(315);
				match(OPEN_SQUARE_BRACKET);
				setState(316);
				expresion(0);
				setState(317);
				match(CLOSE_SQUARE_BRACKET);
				}
				break;
			case 6:
				{
				_localctx = new PrintLLamadaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(319);
				funcLLamada();
				}
				break;
			case 7:
				{
				_localctx = new ValorsimpleContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(320);
				match(NULO);
				}
				break;
			case 8:
				{
				_localctx = new NumberValorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(321);
				match(NUMBER);
				}
				break;
			case 9:
				{
				_localctx = new StringValorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(322);
				match(STRING_SINTAX);
				}
				break;
			case 10:
				{
				_localctx = new BoolValorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(323);
				match(VALORBOOL);
				}
				break;
			case 11:
				{
				_localctx = new SinValorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(324);
				match(SINVALOR);
				}
				break;
			case 12:
				{
				_localctx = new IdValorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(325);
				match(ID);
				}
				break;
			case 13:
				{
				_localctx = new CasteoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(326);
				tipo();
				setState(327);
				match(OPEN_PAREN);
				setState(328);
				expresion(0);
				setState(329);
				match(CLOSE_PAREN);
				}
				break;
			case 14:
				{
				_localctx = new SelfCALLContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(331);
				match(SELF);
				setState(332);
				match(PUNTO);
				setState(333);
				match(ID);
				setState(334);
				match(IGUAL);
				setState(335);
				expresion(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(355);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(353);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
					case 1:
						{
						_localctx = new OpAritmeticoContext(new ExpresionContext(_parentctx, _parentState));
						((OpAritmeticoContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(338);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(339);
						((OpAritmeticoContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 14680064L) != 0)) ) {
							((OpAritmeticoContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(340);
						((OpAritmeticoContext)_localctx).right = expresion(20);
						}
						break;
					case 2:
						{
						_localctx = new OpAritmeticoContext(new ExpresionContext(_parentctx, _parentState));
						((OpAritmeticoContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(341);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(342);
						((OpAritmeticoContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAS || _la==MENOS) ) {
							((OpAritmeticoContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(343);
						((OpAritmeticoContext)_localctx).right = expresion(19);
						}
						break;
					case 3:
						{
						_localctx = new OpComparativoContext(new ExpresionContext(_parentctx, _parentState));
						((OpComparativoContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(344);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(345);
						((OpComparativoContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==IGUAL_IGUAL || _la==DIFERENTE) ) {
							((OpComparativoContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(346);
						((OpComparativoContext)_localctx).right = expresion(18);
						}
						break;
					case 4:
						{
						_localctx = new OpRelacionalContext(new ExpresionContext(_parentctx, _parentState));
						((OpRelacionalContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(347);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(348);
						((OpRelacionalContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 251658240L) != 0)) ) {
							((OpRelacionalContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(349);
						((OpRelacionalContext)_localctx).right = expresion(17);
						}
						break;
					case 5:
						{
						_localctx = new OpLogicoContext(new ExpresionContext(_parentctx, _parentState));
						((OpLogicoContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(350);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(351);
						((OpLogicoContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==AND || _la==OR) ) {
							((OpLogicoContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(352);
						((OpLogicoContext)_localctx).right = expresion(16);
						}
						break;
					}
					} 
				}
				setState(357);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncionesContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(SintaxParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode OPEN_PAREN() { return getToken(SintaxParser.OPEN_PAREN, 0); }
		public TerminalNode CLOSE_PAREN() { return getToken(SintaxParser.CLOSE_PAREN, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public ParametrosContext parametros() {
			return getRuleContext(ParametrosContext.class,0);
		}
		public TerminalNode FLECHA() { return getToken(SintaxParser.FLECHA, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public FuncionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funciones; }
	}

	public final FuncionesContext funciones() throws RecognitionException {
		FuncionesContext _localctx = new FuncionesContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_funciones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(358);
			match(FUNC);
			setState(359);
			match(ID);
			setState(360);
			match(OPEN_PAREN);
			setState(362);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==GUION_BAJO || _la==ID) {
				{
				setState(361);
				parametros();
				}
			}

			setState(364);
			match(CLOSE_PAREN);
			setState(367);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FLECHA) {
				{
				setState(365);
				match(FLECHA);
				setState(366);
				tipo();
				}
			}

			setState(369);
			bloque();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(SintaxParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SintaxParser.ID, i);
		}
		public List<TerminalNode> DOSPUNTOS() { return getTokens(SintaxParser.DOSPUNTOS); }
		public TerminalNode DOSPUNTOS(int i) {
			return getToken(SintaxParser.DOSPUNTOS, i);
		}
		public List<ExisteExIntContext> existeExInt() {
			return getRuleContexts(ExisteExIntContext.class);
		}
		public ExisteExIntContext existeExInt(int i) {
			return getRuleContext(ExisteExIntContext.class,i);
		}
		public List<TipoinoutContext> tipoinout() {
			return getRuleContexts(TipoinoutContext.class);
		}
		public TipoinoutContext tipoinout(int i) {
			return getRuleContext(TipoinoutContext.class,i);
		}
		public List<TipofuncionContext> tipofuncion() {
			return getRuleContexts(TipofuncionContext.class);
		}
		public TipofuncionContext tipofuncion(int i) {
			return getRuleContext(TipofuncionContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(SintaxParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SintaxParser.COMA, i);
		}
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(372);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				{
				setState(371);
				existeExInt();
				}
				break;
			}
			setState(374);
			match(ID);
			setState(375);
			match(DOSPUNTOS);
			setState(377);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INOUT) {
				{
				setState(376);
				tipoinout();
				}
			}

			setState(380);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 16446L) != 0)) {
				{
				setState(379);
				tipofuncion();
				}
			}

			setState(396);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(382);
				match(COMA);
				setState(384);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
				case 1:
					{
					setState(383);
					existeExInt();
					}
					break;
				}
				setState(386);
				match(ID);
				setState(387);
				match(DOSPUNTOS);
				setState(389);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(388);
					tipoinout();
					}
				}

				setState(392);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 16446L) != 0)) {
					{
					setState(391);
					tipofuncion();
					}
				}

				}
				}
				setState(398);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExisteExIntContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode GUION_BAJO() { return getToken(SintaxParser.GUION_BAJO, 0); }
		public ExisteExIntContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_existeExInt; }
	}

	public final ExisteExIntContext existeExInt() throws RecognitionException {
		ExisteExIntContext _localctx = new ExisteExIntContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_existeExInt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(399);
			_la = _input.LA(1);
			if ( !(_la==GUION_BAJO || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TipofuncionContext extends ParserRuleContext {
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode OPEN_SQUARE_BRACKET() { return getToken(SintaxParser.OPEN_SQUARE_BRACKET, 0); }
		public TerminalNode CLOSE_SQUARE_BRACKET() { return getToken(SintaxParser.CLOSE_SQUARE_BRACKET, 0); }
		public TipofuncionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipofuncion; }
	}

	public final TipofuncionContext tipofuncion() throws RecognitionException {
		TipofuncionContext _localctx = new TipofuncionContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_tipofuncion);
		try {
			setState(406);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
			case FLOAT:
			case STRING:
			case BOOL:
			case CHARACTER:
				enterOuterAlt(_localctx, 1);
				{
				setState(401);
				tipo();
				}
				break;
			case OPEN_SQUARE_BRACKET:
				enterOuterAlt(_localctx, 2);
				{
				setState(402);
				match(OPEN_SQUARE_BRACKET);
				setState(403);
				tipo();
				setState(404);
				match(CLOSE_SQUARE_BRACKET);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TipoinoutContext extends ParserRuleContext {
		public TerminalNode INOUT() { return getToken(SintaxParser.INOUT, 0); }
		public TipoinoutContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipoinout; }
	}

	public final TipoinoutContext tipoinout() throws RecognitionException {
		TipoinoutContext _localctx = new TipoinoutContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_tipoinout);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(408);
			match(INOUT);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncLLamadaContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode OPEN_PAREN() { return getToken(SintaxParser.OPEN_PAREN, 0); }
		public TerminalNode CLOSE_PAREN() { return getToken(SintaxParser.CLOSE_PAREN, 0); }
		public ParametrosLLamadaContext parametrosLLamada() {
			return getRuleContext(ParametrosLLamadaContext.class,0);
		}
		public FuncLLamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcLLamada; }
	}

	public final FuncLLamadaContext funcLLamada() throws RecognitionException {
		FuncLLamadaContext _localctx = new FuncLLamadaContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_funcLLamada);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(410);
			match(ID);
			setState(411);
			match(OPEN_PAREN);
			setState(413);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & -2017612628767014594L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 51L) != 0)) {
				{
				setState(412);
				parametrosLLamada();
				}
			}

			setState(415);
			match(CLOSE_PAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosLLamadaContext extends ParserRuleContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> ID() { return getTokens(SintaxParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SintaxParser.ID, i);
		}
		public List<TerminalNode> DOSPUNTOS() { return getTokens(SintaxParser.DOSPUNTOS); }
		public TerminalNode DOSPUNTOS(int i) {
			return getToken(SintaxParser.DOSPUNTOS, i);
		}
		public List<TerminalNode> REFERENCIA() { return getTokens(SintaxParser.REFERENCIA); }
		public TerminalNode REFERENCIA(int i) {
			return getToken(SintaxParser.REFERENCIA, i);
		}
		public List<TerminalNode> COMA() { return getTokens(SintaxParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SintaxParser.COMA, i);
		}
		public ParametrosLLamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosLLamada; }
	}

	public final ParametrosLLamadaContext parametrosLLamada() throws RecognitionException {
		ParametrosLLamadaContext _localctx = new ParametrosLLamadaContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_parametrosLLamada);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(419);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				{
				setState(417);
				match(ID);
				setState(418);
				match(DOSPUNTOS);
				}
				break;
			}
			setState(422);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==REFERENCIA) {
				{
				setState(421);
				match(REFERENCIA);
				}
			}

			setState(424);
			expresion(0);
			setState(436);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(425);
				match(COMA);
				setState(428);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
				case 1:
					{
					setState(426);
					match(ID);
					setState(427);
					match(DOSPUNTOS);
					}
					break;
				}
				setState(431);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==REFERENCIA) {
					{
					setState(430);
					match(REFERENCIA);
					}
				}

				setState(433);
				expresion(0);
				}
				}
				setState(438);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EstructsContext extends ParserRuleContext {
		public TerminalNode STRUCT() { return getToken(SintaxParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode OPEN_BRACKET() { return getToken(SintaxParser.OPEN_BRACKET, 0); }
		public TerminalNode CLOSE_BRACKET() { return getToken(SintaxParser.CLOSE_BRACKET, 0); }
		public List<Lista_structContext> lista_struct() {
			return getRuleContexts(Lista_structContext.class);
		}
		public Lista_structContext lista_struct(int i) {
			return getRuleContext(Lista_structContext.class,i);
		}
		public EstructsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_estructs; }
	}

	public final EstructsContext estructs() throws RecognitionException {
		EstructsContext _localctx = new EstructsContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_estructs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(439);
			match(STRUCT);
			setState(440);
			match(ID);
			setState(441);
			match(OPEN_BRACKET);
			setState(445);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1188950327395614720L) != 0)) {
				{
				{
				setState(442);
				lista_struct();
				}
				}
				setState(447);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(448);
			match(CLOSE_BRACKET);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Lista_structContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode LET() { return getToken(SintaxParser.LET, 0); }
		public TerminalNode VAR() { return getToken(SintaxParser.VAR, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SintaxParser.DOSPUNTOS, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(SintaxParser.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public FuncionestructContext funcionestruct() {
			return getRuleContext(FuncionestructContext.class,0);
		}
		public TerminalNode MUTATING() { return getToken(SintaxParser.MUTATING, 0); }
		public Lista_structContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_struct; }
	}

	public final Lista_structContext lista_struct() throws RecognitionException {
		Lista_structContext _localctx = new Lista_structContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_lista_struct);
		int _la;
		try {
			setState(464);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
			case LET:
				enterOuterAlt(_localctx, 1);
				{
				setState(450);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(451);
				match(ID);
				setState(454);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPUNTOS) {
					{
					setState(452);
					match(DOSPUNTOS);
					setState(453);
					tipo();
					}
				}

				setState(458);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IGUAL) {
					{
					setState(456);
					match(IGUAL);
					setState(457);
					expresion(0);
					}
				}

				}
				break;
			case FUNC:
			case MUTATING:
				enterOuterAlt(_localctx, 2);
				{
				setState(461);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MUTATING) {
					{
					setState(460);
					match(MUTATING);
					}
				}

				setState(463);
				funcionestruct();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncionestructContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(SintaxParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SintaxParser.ID, 0); }
		public TerminalNode OPEN_PAREN() { return getToken(SintaxParser.OPEN_PAREN, 0); }
		public TerminalNode CLOSE_PAREN() { return getToken(SintaxParser.CLOSE_PAREN, 0); }
		public TerminalNode OPEN_BRACKET() { return getToken(SintaxParser.OPEN_BRACKET, 0); }
		public TerminalNode CLOSE_BRACKET() { return getToken(SintaxParser.CLOSE_BRACKET, 0); }
		public List<Lista_procesoContext> lista_proceso() {
			return getRuleContexts(Lista_procesoContext.class);
		}
		public Lista_procesoContext lista_proceso(int i) {
			return getRuleContext(Lista_procesoContext.class,i);
		}
		public FuncionestructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionestruct; }
	}

	public final FuncionestructContext funcionestruct() throws RecognitionException {
		FuncionestructContext _localctx = new FuncionestructContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_funcionestruct);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(466);
			match(FUNC);
			setState(467);
			match(ID);
			setState(468);
			match(OPEN_PAREN);
			setState(469);
			match(CLOSE_PAREN);
			setState(470);
			match(OPEN_BRACKET);
			setState(474);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & 481103493935L) != 0)) {
				{
				{
				setState(471);
				lista_proceso();
				}
				}
				setState(476);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(477);
			match(CLOSE_BRACKET);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 30:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 19);
		case 1:
			return precpred(_ctx, 18);
		case 2:
			return precpred(_ctx, 17);
		case 3:
			return precpred(_ctx, 16);
		case 4:
			return precpred(_ctx, 15);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001H\u01e0\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0005"+
		"\u0001X\b\u0001\n\u0001\f\u0001[\t\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002j\b"+
		"\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0003\u0004p\b"+
		"\u0004\u0001\u0004\u0001\u0004\u0003\u0004t\b\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0003\u0005z\b\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u0081\b\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\b\u0001\b\u0003\b\u008c\b\b\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0005\f\u00a0\b\f"+
		"\n\f\f\f\u00a3\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0005"+
		"\r\u00ab\b\r\n\r\f\r\u00ae\t\r\u0001\r\u0003\r\u00b1\b\r\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0004\u0010\u00bf"+
		"\b\u0010\u000b\u0010\f\u0010\u00c0\u0001\u0010\u0003\u0010\u00c4\b\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0005\u0011\u00cc\b\u0011\n\u0011\f\u0011\u00cf\t\u0011\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0005\u0012\u00d4\b\u0012\n\u0012\f\u0012\u00d7\t\u0012"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u00e3\b\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0005\u0015\u00ec\b\u0015\n\u0015\f\u0015\u00ef\t\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0005\u0016\u00fe\b\u0016\n\u0016\f\u0016\u0101\t\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0003\u0018\u0111\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0003\u0019\u0118\b\u0019\u0001\u0019\u0003\u0019"+
		"\u011b\b\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001b"+
		"\u0001\u001b\u0005\u001b\u0123\b\u001b\n\u001b\f\u001b\u0126\t\u001b\u0001"+
		"\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001d\u0001\u001d\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0003\u001e\u0151"+
		"\b\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0005\u001e\u0162\b\u001e\n"+
		"\u001e\f\u001e\u0165\t\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0003\u001f\u016b\b\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003"+
		"\u001f\u0170\b\u001f\u0001\u001f\u0001\u001f\u0001 \u0003 \u0175\b \u0001"+
		" \u0001 \u0001 \u0003 \u017a\b \u0001 \u0003 \u017d\b \u0001 \u0001 \u0003"+
		" \u0181\b \u0001 \u0001 \u0001 \u0003 \u0186\b \u0001 \u0003 \u0189\b"+
		" \u0005 \u018b\b \n \f \u018e\t \u0001!\u0001!\u0001\"\u0001\"\u0001\""+
		"\u0001\"\u0001\"\u0003\"\u0197\b\"\u0001#\u0001#\u0001$\u0001$\u0001$"+
		"\u0003$\u019e\b$\u0001$\u0001$\u0001%\u0001%\u0003%\u01a4\b%\u0001%\u0003"+
		"%\u01a7\b%\u0001%\u0001%\u0001%\u0001%\u0003%\u01ad\b%\u0001%\u0003%\u01b0"+
		"\b%\u0001%\u0005%\u01b3\b%\n%\f%\u01b6\t%\u0001&\u0001&\u0001&\u0001&"+
		"\u0005&\u01bc\b&\n&\f&\u01bf\t&\u0001&\u0001&\u0001\'\u0001\'\u0001\'"+
		"\u0001\'\u0003\'\u01c7\b\'\u0001\'\u0001\'\u0003\'\u01cb\b\'\u0001\'\u0003"+
		"\'\u01ce\b\'\u0001\'\u0003\'\u01d1\b\'\u0001(\u0001(\u0001(\u0001(\u0001"+
		"(\u0001(\u0005(\u01d9\b(\n(\f(\u01dc\t(\u0001(\u0001(\u0001(\u0000\u0001"+
		"<)\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a"+
		"\u001c\u001e \"$&(*,.02468:<>@BDFHJLNP\u0000\u000e\u0001\u0000FG\u0001"+
		"\u0000!\"\u0002\u0000\u0006\u0006\n\u000b\u0002\u0000\u0006\u0006\n\n"+
		"\u0002\u0000>>EE\u0002\u0000\u0012\u0012EE\u0002\u00001134\u0001\u0000"+
		",.\u0001\u0000\u0001\u0005\u0001\u0000\u0015\u0017\u0001\u0000\u0013\u0014"+
		"\u0001\u0000\u001c\u001d\u0001\u0000\u0018\u001b\u0001\u0000\u001e\u001f"+
		"\u0202\u0000R\u0001\u0000\u0000\u0000\u0002Y\u0001\u0000\u0000\u0000\u0004"+
		"i\u0001\u0000\u0000\u0000\u0006k\u0001\u0000\u0000\u0000\bs\u0001\u0000"+
		"\u0000\u0000\ny\u0001\u0000\u0000\u0000\f{\u0001\u0000\u0000\u0000\u000e"+
		"\u0084\u0001\u0000\u0000\u0000\u0010\u008b\u0001\u0000\u0000\u0000\u0012"+
		"\u008d\u0001\u0000\u0000\u0000\u0014\u0091\u0001\u0000\u0000\u0000\u0016"+
		"\u0098\u0001\u0000\u0000\u0000\u0018\u009a\u0001\u0000\u0000\u0000\u001a"+
		"\u00a6\u0001\u0000\u0000\u0000\u001c\u00b2\u0001\u0000\u0000\u0000\u001e"+
		"\u00b7\u0001\u0000\u0000\u0000 \u00ba\u0001\u0000\u0000\u0000\"\u00c7"+
		"\u0001\u0000\u0000\u0000$\u00d0\u0001\u0000\u0000\u0000&\u00d8\u0001\u0000"+
		"\u0000\u0000(\u00dc\u0001\u0000\u0000\u0000*\u00e6\u0001\u0000\u0000\u0000"+
		",\u00f2\u0001\u0000\u0000\u0000.\u0104\u0001\u0000\u0000\u00000\u0110"+
		"\u0001\u0000\u0000\u00002\u0112\u0001\u0000\u0000\u00004\u011e\u0001\u0000"+
		"\u0000\u00006\u0120\u0001\u0000\u0000\u00008\u0129\u0001\u0000\u0000\u0000"+
		":\u012b\u0001\u0000\u0000\u0000<\u0150\u0001\u0000\u0000\u0000>\u0166"+
		"\u0001\u0000\u0000\u0000@\u0174\u0001\u0000\u0000\u0000B\u018f\u0001\u0000"+
		"\u0000\u0000D\u0196\u0001\u0000\u0000\u0000F\u0198\u0001\u0000\u0000\u0000"+
		"H\u019a\u0001\u0000\u0000\u0000J\u01a3\u0001\u0000\u0000\u0000L\u01b7"+
		"\u0001\u0000\u0000\u0000N\u01d0\u0001\u0000\u0000\u0000P\u01d2\u0001\u0000"+
		"\u0000\u0000RS\u0003\u0002\u0001\u0000ST\u0005\u0000\u0000\u0001T\u0001"+
		"\u0001\u0000\u0000\u0000UX\u0003\u0004\u0002\u0000VX\u0003>\u001f\u0000"+
		"WU\u0001\u0000\u0000\u0000WV\u0001\u0000\u0000\u0000X[\u0001\u0000\u0000"+
		"\u0000YW\u0001\u0000\u0000\u0000YZ\u0001\u0000\u0000\u0000Z\u0003\u0001"+
		"\u0000\u0000\u0000[Y\u0001\u0000\u0000\u0000\\j\u0003\n\u0005\u0000]j"+
		"\u0003\u0010\b\u0000^j\u0003\u0018\f\u0000_j\u0003\u001a\r\u0000`j\u0003"+
		" \u0010\u0000aj\u0003&\u0013\u0000bj\u0003(\u0014\u0000cj\u0003*\u0015"+
		"\u0000dj\u00032\u0019\u0000ej\u0003H$\u0000fj\u0003\b\u0004\u0000gj\u0003"+
		"L&\u0000hj\u0003\u0006\u0003\u0000i\\\u0001\u0000\u0000\u0000i]\u0001"+
		"\u0000\u0000\u0000i^\u0001\u0000\u0000\u0000i_\u0001\u0000\u0000\u0000"+
		"i`\u0001\u0000\u0000\u0000ia\u0001\u0000\u0000\u0000ib\u0001\u0000\u0000"+
		"\u0000ic\u0001\u0000\u0000\u0000id\u0001\u0000\u0000\u0000ie\u0001\u0000"+
		"\u0000\u0000if\u0001\u0000\u0000\u0000ig\u0001\u0000\u0000\u0000ih\u0001"+
		"\u0000\u0000\u0000j\u0005\u0001\u0000\u0000\u0000kl\u0007\u0000\u0000"+
		"\u0000l\u0007\u0001\u0000\u0000\u0000mo\u0005.\u0000\u0000np\u0003<\u001e"+
		"\u0000on\u0001\u0000\u0000\u0000op\u0001\u0000\u0000\u0000pt\u0001\u0000"+
		"\u0000\u0000qt\u0005,\u0000\u0000rt\u0005-\u0000\u0000sm\u0001\u0000\u0000"+
		"\u0000sq\u0001\u0000\u0000\u0000sr\u0001\u0000\u0000\u0000t\t\u0001\u0000"+
		"\u0000\u0000uz\u0003\f\u0006\u0000vz\u0003\u000e\u0007\u0000wz\u0003,"+
		"\u0016\u0000xz\u0003.\u0017\u0000yu\u0001\u0000\u0000\u0000yv\u0001\u0000"+
		"\u0000\u0000yw\u0001\u0000\u0000\u0000yx\u0001\u0000\u0000\u0000z\u000b"+
		"\u0001\u0000\u0000\u0000{|\u0007\u0001\u0000\u0000|}\u0005E\u0000\u0000"+
		"}~\u0005\u0007\u0000\u0000~\u0080\u0003:\u001d\u0000\u007f\u0081\u0005"+
		"\u0006\u0000\u0000\u0080\u007f\u0001\u0000\u0000\u0000\u0080\u0081\u0001"+
		"\u0000\u0000\u0000\u0081\u0082\u0001\u0000\u0000\u0000\u0082\u0083\u0003"+
		"<\u001e\u0000\u0083\r\u0001\u0000\u0000\u0000\u0084\u0085\u0007\u0001"+
		"\u0000\u0000\u0085\u0086\u0005E\u0000\u0000\u0086\u0087\u0005\u0006\u0000"+
		"\u0000\u0087\u0088\u0003<\u001e\u0000\u0088\u000f\u0001\u0000\u0000\u0000"+
		"\u0089\u008c\u0003\u0012\t\u0000\u008a\u008c\u0003\u0014\n\u0000\u008b"+
		"\u0089\u0001\u0000\u0000\u0000\u008b\u008a\u0001\u0000\u0000\u0000\u008c"+
		"\u0011\u0001\u0000\u0000\u0000\u008d\u008e\u0005E\u0000\u0000\u008e\u008f"+
		"\u0007\u0002\u0000\u0000\u008f\u0090\u0003<\u001e\u0000\u0090\u0013\u0001"+
		"\u0000\u0000\u0000\u0091\u0092\u0005E\u0000\u0000\u0092\u0093\u0005\u000e"+
		"\u0000\u0000\u0093\u0094\u0003\u0016\u000b\u0000\u0094\u0095\u0005\u000f"+
		"\u0000\u0000\u0095\u0096\u0007\u0003\u0000\u0000\u0096\u0097\u0003<\u001e"+
		"\u0000\u0097\u0015\u0001\u0000\u0000\u0000\u0098\u0099\u0007\u0004\u0000"+
		"\u0000\u0099\u0017\u0001\u0000\u0000\u0000\u009a\u009b\u0005#\u0000\u0000"+
		"\u009b\u009c\u0005\b\u0000\u0000\u009c\u00a1\u0003<\u001e\u0000\u009d"+
		"\u009e\u0005\u0010\u0000\u0000\u009e\u00a0\u0003<\u001e\u0000\u009f\u009d"+
		"\u0001\u0000\u0000\u0000\u00a0\u00a3\u0001\u0000\u0000\u0000\u00a1\u009f"+
		"\u0001\u0000\u0000\u0000\u00a1\u00a2\u0001\u0000\u0000\u0000\u00a2\u00a4"+
		"\u0001\u0000\u0000\u0000\u00a3\u00a1\u0001\u0000\u0000\u0000\u00a4\u00a5"+
		"\u0005\t\u0000\u0000\u00a5\u0019\u0001\u0000\u0000\u0000\u00a6\u00a7\u0005"+
		"$\u0000\u0000\u00a7\u00a8\u0003<\u001e\u0000\u00a8\u00ac\u00036\u001b"+
		"\u0000\u00a9\u00ab\u0003\u001c\u000e\u0000\u00aa\u00a9\u0001\u0000\u0000"+
		"\u0000\u00ab\u00ae\u0001\u0000\u0000\u0000\u00ac\u00aa\u0001\u0000\u0000"+
		"\u0000\u00ac\u00ad\u0001\u0000\u0000\u0000\u00ad\u00b0\u0001\u0000\u0000"+
		"\u0000\u00ae\u00ac\u0001\u0000\u0000\u0000\u00af\u00b1\u0003\u001e\u000f"+
		"\u0000\u00b0\u00af\u0001\u0000\u0000\u0000\u00b0\u00b1\u0001\u0000\u0000"+
		"\u0000\u00b1\u001b\u0001\u0000\u0000\u0000\u00b2\u00b3\u0005%\u0000\u0000"+
		"\u00b3\u00b4\u0005$\u0000\u0000\u00b4\u00b5\u0003<\u001e\u0000\u00b5\u00b6"+
		"\u00036\u001b\u0000\u00b6\u001d\u0001\u0000\u0000\u0000\u00b7\u00b8\u0005"+
		"%\u0000\u0000\u00b8\u00b9\u00036\u001b\u0000\u00b9\u001f\u0001\u0000\u0000"+
		"\u0000\u00ba\u00bb\u0005&\u0000\u0000\u00bb\u00bc\u0003<\u001e\u0000\u00bc"+
		"\u00be\u0005\f\u0000\u0000\u00bd\u00bf\u0003\"\u0011\u0000\u00be\u00bd"+
		"\u0001\u0000\u0000\u0000\u00bf\u00c0\u0001\u0000\u0000\u0000\u00c0\u00be"+
		"\u0001\u0000\u0000\u0000\u00c0\u00c1\u0001\u0000\u0000\u0000\u00c1\u00c3"+
		"\u0001\u0000\u0000\u0000\u00c2\u00c4\u0003$\u0012\u0000\u00c3\u00c2\u0001"+
		"\u0000\u0000\u0000\u00c3\u00c4\u0001\u0000\u0000\u0000\u00c4\u00c5\u0001"+
		"\u0000\u0000\u0000\u00c5\u00c6\u0005\r\u0000\u0000\u00c6!\u0001\u0000"+
		"\u0000\u0000\u00c7\u00c8\u0005\'\u0000\u0000\u00c8\u00c9\u0003<\u001e"+
		"\u0000\u00c9\u00cd\u0005\u0007\u0000\u0000\u00ca\u00cc\u0003\u0004\u0002"+
		"\u0000\u00cb\u00ca\u0001\u0000\u0000\u0000\u00cc\u00cf\u0001\u0000\u0000"+
		"\u0000\u00cd\u00cb\u0001\u0000\u0000\u0000\u00cd\u00ce\u0001\u0000\u0000"+
		"\u0000\u00ce#\u0001\u0000\u0000\u0000\u00cf\u00cd\u0001\u0000\u0000\u0000"+
		"\u00d0\u00d1\u0005(\u0000\u0000\u00d1\u00d5\u0005\u0007\u0000\u0000\u00d2"+
		"\u00d4\u0003\u0004\u0002\u0000\u00d3\u00d2\u0001\u0000\u0000\u0000\u00d4"+
		"\u00d7\u0001\u0000\u0000\u0000\u00d5\u00d3\u0001\u0000\u0000\u0000\u00d5"+
		"\u00d6\u0001\u0000\u0000\u0000\u00d6%\u0001\u0000\u0000\u0000\u00d7\u00d5"+
		"\u0001\u0000\u0000\u0000\u00d8\u00d9\u0005)\u0000\u0000\u00d9\u00da\u0003"+
		"<\u001e\u0000\u00da\u00db\u00036\u001b\u0000\u00db\'\u0001\u0000\u0000"+
		"\u0000\u00dc\u00dd\u0005*\u0000\u0000\u00dd\u00de\u0007\u0005\u0000\u0000"+
		"\u00de\u00df\u0005+\u0000\u0000\u00df\u00e2\u0003<\u001e\u0000\u00e0\u00e1"+
		"\u0005/\u0000\u0000\u00e1\u00e3\u0003<\u001e\u0000\u00e2\u00e0\u0001\u0000"+
		"\u0000\u0000\u00e2\u00e3\u0001\u0000\u0000\u0000\u00e3\u00e4\u0001\u0000"+
		"\u0000\u0000\u00e4\u00e5\u00036\u001b\u0000\u00e5)\u0001\u0000\u0000\u0000"+
		"\u00e6\u00e7\u00050\u0000\u0000\u00e7\u00e8\u0003<\u001e\u0000\u00e8\u00e9"+
		"\u0005%\u0000\u0000\u00e9\u00ed\u0005\f\u0000\u0000\u00ea\u00ec\u0003"+
		"\u0004\u0002\u0000\u00eb\u00ea\u0001\u0000\u0000\u0000\u00ec\u00ef\u0001"+
		"\u0000\u0000\u0000\u00ed\u00eb\u0001\u0000\u0000\u0000\u00ed\u00ee\u0001"+
		"\u0000\u0000\u0000\u00ee\u00f0\u0001\u0000\u0000\u0000\u00ef\u00ed\u0001"+
		"\u0000\u0000\u0000\u00f0\u00f1\u0005\r\u0000\u0000\u00f1+\u0001\u0000"+
		"\u0000\u0000\u00f2\u00f3\u0007\u0001\u0000\u0000\u00f3\u00f4\u0005E\u0000"+
		"\u0000\u00f4\u00f5\u0005\u0007\u0000\u0000\u00f5\u00f6\u0005\u000e\u0000"+
		"\u0000\u00f6\u00f7\u0003:\u001d\u0000\u00f7\u00f8\u0005\u000f\u0000\u0000"+
		"\u00f8\u00f9\u0005\u0006\u0000\u0000\u00f9\u00fa\u0005\u000e\u0000\u0000"+
		"\u00fa\u00ff\u0003<\u001e\u0000\u00fb\u00fc\u0005\u0010\u0000\u0000\u00fc"+
		"\u00fe\u0003<\u001e\u0000\u00fd\u00fb\u0001\u0000\u0000\u0000\u00fe\u0101"+
		"\u0001\u0000\u0000\u0000\u00ff\u00fd\u0001\u0000\u0000\u0000\u00ff\u0100"+
		"\u0001\u0000\u0000\u0000\u0100\u0102\u0001\u0000\u0000\u0000\u0101\u00ff"+
		"\u0001\u0000\u0000\u0000\u0102\u0103\u0005\u000f\u0000\u0000\u0103-\u0001"+
		"\u0000\u0000\u0000\u0104\u0105\u0007\u0001\u0000\u0000\u0105\u0106\u0005"+
		"E\u0000\u0000\u0106\u0107\u0005\u0007\u0000\u0000\u0107\u0108\u0005\u000e"+
		"\u0000\u0000\u0108\u0109\u0003:\u001d\u0000\u0109\u010a\u0005\u000f\u0000"+
		"\u0000\u010a\u010b\u0005\u0006\u0000\u0000\u010b\u010c\u00030\u0018\u0000"+
		"\u010c/\u0001\u0000\u0000\u0000\u010d\u010e\u0005\u000e\u0000\u0000\u010e"+
		"\u0111\u0005\u000f\u0000\u0000\u010f\u0111\u0005E\u0000\u0000\u0110\u010d"+
		"\u0001\u0000\u0000\u0000\u0110\u010f\u0001\u0000\u0000\u0000\u01111\u0001"+
		"\u0000\u0000\u0000\u0112\u0113\u0005E\u0000\u0000\u0113\u0114\u0005\u0011"+
		"\u0000\u0000\u0114\u0115\u00034\u001a\u0000\u0115\u0117\u0005\b\u0000"+
		"\u0000\u0116\u0118\u00055\u0000\u0000\u0117\u0116\u0001\u0000\u0000\u0000"+
		"\u0117\u0118\u0001\u0000\u0000\u0000\u0118\u011a\u0001\u0000\u0000\u0000"+
		"\u0119\u011b\u0003<\u001e\u0000\u011a\u0119\u0001\u0000\u0000\u0000\u011a"+
		"\u011b\u0001\u0000\u0000\u0000\u011b\u011c\u0001\u0000\u0000\u0000\u011c"+
		"\u011d\u0005\t\u0000\u0000\u011d3\u0001\u0000\u0000\u0000\u011e\u011f"+
		"\u0007\u0006\u0000\u0000\u011f5\u0001\u0000\u0000\u0000\u0120\u0124\u0005"+
		"\f\u0000\u0000\u0121\u0123\u0003\u0004\u0002\u0000\u0122\u0121\u0001\u0000"+
		"\u0000\u0000\u0123\u0126\u0001\u0000\u0000\u0000\u0124\u0122\u0001\u0000"+
		"\u0000\u0000\u0124\u0125\u0001\u0000\u0000\u0000\u0125\u0127\u0001\u0000"+
		"\u0000\u0000\u0126\u0124\u0001\u0000\u0000\u0000\u0127\u0128\u0005\r\u0000"+
		"\u0000\u01287\u0001\u0000\u0000\u0000\u0129\u012a\u0007\u0007\u0000\u0000"+
		"\u012a9\u0001\u0000\u0000\u0000\u012b\u012c\u0007\b\u0000\u0000\u012c"+
		";\u0001\u0000\u0000\u0000\u012d\u012e\u0006\u001e\uffff\uffff\u0000\u012e"+
		"\u012f\u0005 \u0000\u0000\u012f\u0151\u0003<\u001e\u000e\u0130\u0131\u0005"+
		"\b\u0000\u0000\u0131\u0132\u0003<\u001e\u0000\u0132\u0133\u0005\t\u0000"+
		"\u0000\u0133\u0151\u0001\u0000\u0000\u0000\u0134\u0135\u0005E\u0000\u0000"+
		"\u0135\u0136\u0005\u0011\u0000\u0000\u0136\u0151\u00056\u0000\u0000\u0137"+
		"\u0138\u0005E\u0000\u0000\u0138\u0139\u0005\u0011\u0000\u0000\u0139\u0151"+
		"\u00052\u0000\u0000\u013a\u013b\u0005E\u0000\u0000\u013b\u013c\u0005\u000e"+
		"\u0000\u0000\u013c\u013d\u0003<\u001e\u0000\u013d\u013e\u0005\u000f\u0000"+
		"\u0000\u013e\u0151\u0001\u0000\u0000\u0000\u013f\u0151\u0003H$\u0000\u0140"+
		"\u0151\u0005@\u0000\u0000\u0141\u0151\u0005>\u0000\u0000\u0142\u0151\u0005"+
		"A\u0000\u0000\u0143\u0151\u0005?\u0000\u0000\u0144\u0151\u0005D\u0000"+
		"\u0000\u0145\u0151\u0005E\u0000\u0000\u0146\u0147\u0003:\u001d\u0000\u0147"+
		"\u0148\u0005\b\u0000\u0000\u0148\u0149\u0003<\u001e\u0000\u0149\u014a"+
		"\u0005\t\u0000\u0000\u014a\u0151\u0001\u0000\u0000\u0000\u014b\u014c\u0005"+
		"=\u0000\u0000\u014c\u014d\u0005\u0011\u0000\u0000\u014d\u014e\u0005E\u0000"+
		"\u0000\u014e\u014f\u0005\u0006\u0000\u0000\u014f\u0151\u0003<\u001e\u0001"+
		"\u0150\u012d\u0001\u0000\u0000\u0000\u0150\u0130\u0001\u0000\u0000\u0000"+
		"\u0150\u0134\u0001\u0000\u0000\u0000\u0150\u0137\u0001\u0000\u0000\u0000"+
		"\u0150\u013a\u0001\u0000\u0000\u0000\u0150\u013f\u0001\u0000\u0000\u0000"+
		"\u0150\u0140\u0001\u0000\u0000\u0000\u0150\u0141\u0001\u0000\u0000\u0000"+
		"\u0150\u0142\u0001\u0000\u0000\u0000\u0150\u0143\u0001\u0000\u0000\u0000"+
		"\u0150\u0144\u0001\u0000\u0000\u0000\u0150\u0145\u0001\u0000\u0000\u0000"+
		"\u0150\u0146\u0001\u0000\u0000\u0000\u0150\u014b\u0001\u0000\u0000\u0000"+
		"\u0151\u0163\u0001\u0000\u0000\u0000\u0152\u0153\n\u0013\u0000\u0000\u0153"+
		"\u0154\u0007\t\u0000\u0000\u0154\u0162\u0003<\u001e\u0014\u0155\u0156"+
		"\n\u0012\u0000\u0000\u0156\u0157\u0007\n\u0000\u0000\u0157\u0162\u0003"+
		"<\u001e\u0013\u0158\u0159\n\u0011\u0000\u0000\u0159\u015a\u0007\u000b"+
		"\u0000\u0000\u015a\u0162\u0003<\u001e\u0012\u015b\u015c\n\u0010\u0000"+
		"\u0000\u015c\u015d\u0007\f\u0000\u0000\u015d\u0162\u0003<\u001e\u0011"+
		"\u015e\u015f\n\u000f\u0000\u0000\u015f\u0160\u0007\r\u0000\u0000\u0160"+
		"\u0162\u0003<\u001e\u0010\u0161\u0152\u0001\u0000\u0000\u0000\u0161\u0155"+
		"\u0001\u0000\u0000\u0000\u0161\u0158\u0001\u0000\u0000\u0000\u0161\u015b"+
		"\u0001\u0000\u0000\u0000\u0161\u015e\u0001\u0000\u0000\u0000\u0162\u0165"+
		"\u0001\u0000\u0000\u0000\u0163\u0161\u0001\u0000\u0000\u0000\u0163\u0164"+
		"\u0001\u0000\u0000\u0000\u0164=\u0001\u0000\u0000\u0000\u0165\u0163\u0001"+
		"\u0000\u0000\u0000\u0166\u0167\u00057\u0000\u0000\u0167\u0168\u0005E\u0000"+
		"\u0000\u0168\u016a\u0005\b\u0000\u0000\u0169\u016b\u0003@ \u0000\u016a"+
		"\u0169\u0001\u0000\u0000\u0000\u016a\u016b\u0001\u0000\u0000\u0000\u016b"+
		"\u016c\u0001\u0000\u0000\u0000\u016c\u016f\u0005\t\u0000\u0000\u016d\u016e"+
		"\u00059\u0000\u0000\u016e\u0170\u0003:\u001d\u0000\u016f\u016d\u0001\u0000"+
		"\u0000\u0000\u016f\u0170\u0001\u0000\u0000\u0000\u0170\u0171\u0001\u0000"+
		"\u0000\u0000\u0171\u0172\u00036\u001b\u0000\u0172?\u0001\u0000\u0000\u0000"+
		"\u0173\u0175\u0003B!\u0000\u0174\u0173\u0001\u0000\u0000\u0000\u0174\u0175"+
		"\u0001\u0000\u0000\u0000\u0175\u0176\u0001\u0000\u0000\u0000\u0176\u0177"+
		"\u0005E\u0000\u0000\u0177\u0179\u0005\u0007\u0000\u0000\u0178\u017a\u0003"+
		"F#\u0000\u0179\u0178\u0001\u0000\u0000\u0000\u0179\u017a\u0001\u0000\u0000"+
		"\u0000\u017a\u017c\u0001\u0000\u0000\u0000\u017b\u017d\u0003D\"\u0000"+
		"\u017c\u017b\u0001\u0000\u0000\u0000\u017c\u017d\u0001\u0000\u0000\u0000"+
		"\u017d\u018c\u0001\u0000\u0000\u0000\u017e\u0180\u0005\u0010\u0000\u0000"+
		"\u017f\u0181\u0003B!\u0000\u0180\u017f\u0001\u0000\u0000\u0000\u0180\u0181"+
		"\u0001\u0000\u0000\u0000\u0181\u0182\u0001\u0000\u0000\u0000\u0182\u0183"+
		"\u0005E\u0000\u0000\u0183\u0185\u0005\u0007\u0000\u0000\u0184\u0186\u0003"+
		"F#\u0000\u0185\u0184\u0001\u0000\u0000\u0000\u0185\u0186\u0001\u0000\u0000"+
		"\u0000\u0186\u0188\u0001\u0000\u0000\u0000\u0187\u0189\u0003D\"\u0000"+
		"\u0188\u0187\u0001\u0000\u0000\u0000\u0188\u0189\u0001\u0000\u0000\u0000"+
		"\u0189\u018b\u0001\u0000\u0000\u0000\u018a\u017e\u0001\u0000\u0000\u0000"+
		"\u018b\u018e\u0001\u0000\u0000\u0000\u018c\u018a\u0001\u0000\u0000\u0000"+
		"\u018c\u018d\u0001\u0000\u0000\u0000\u018dA\u0001\u0000\u0000\u0000\u018e"+
		"\u018c\u0001\u0000\u0000\u0000\u018f\u0190\u0007\u0005\u0000\u0000\u0190"+
		"C\u0001\u0000\u0000\u0000\u0191\u0197\u0003:\u001d\u0000\u0192\u0193\u0005"+
		"\u000e\u0000\u0000\u0193\u0194\u0003:\u001d\u0000\u0194\u0195\u0005\u000f"+
		"\u0000\u0000\u0195\u0197\u0001\u0000\u0000\u0000\u0196\u0191\u0001\u0000"+
		"\u0000\u0000\u0196\u0192\u0001\u0000\u0000\u0000\u0197E\u0001\u0000\u0000"+
		"\u0000\u0198\u0199\u00058\u0000\u0000\u0199G\u0001\u0000\u0000\u0000\u019a"+
		"\u019b\u0005E\u0000\u0000\u019b\u019d\u0005\b\u0000\u0000\u019c\u019e"+
		"\u0003J%\u0000\u019d\u019c\u0001\u0000\u0000\u0000\u019d\u019e\u0001\u0000"+
		"\u0000\u0000\u019e\u019f\u0001\u0000\u0000\u0000\u019f\u01a0\u0005\t\u0000"+
		"\u0000\u01a0I\u0001\u0000\u0000\u0000\u01a1\u01a2\u0005E\u0000\u0000\u01a2"+
		"\u01a4\u0005\u0007\u0000\u0000\u01a3\u01a1\u0001\u0000\u0000\u0000\u01a3"+
		"\u01a4\u0001\u0000\u0000\u0000\u01a4\u01a6\u0001\u0000\u0000\u0000\u01a5"+
		"\u01a7\u0005:\u0000\u0000\u01a6\u01a5\u0001\u0000\u0000\u0000\u01a6\u01a7"+
		"\u0001\u0000\u0000\u0000\u01a7\u01a8\u0001\u0000\u0000\u0000\u01a8\u01b4"+
		"\u0003<\u001e\u0000\u01a9\u01ac\u0005\u0010\u0000\u0000\u01aa\u01ab\u0005"+
		"E\u0000\u0000\u01ab\u01ad\u0005\u0007\u0000\u0000\u01ac\u01aa\u0001\u0000"+
		"\u0000\u0000\u01ac\u01ad\u0001\u0000\u0000\u0000\u01ad\u01af\u0001\u0000"+
		"\u0000\u0000\u01ae\u01b0\u0005:\u0000\u0000\u01af\u01ae\u0001\u0000\u0000"+
		"\u0000\u01af\u01b0\u0001\u0000\u0000\u0000\u01b0\u01b1\u0001\u0000\u0000"+
		"\u0000\u01b1\u01b3\u0003<\u001e\u0000\u01b2\u01a9\u0001\u0000\u0000\u0000"+
		"\u01b3\u01b6\u0001\u0000\u0000\u0000\u01b4\u01b2\u0001\u0000\u0000\u0000"+
		"\u01b4\u01b5\u0001\u0000\u0000\u0000\u01b5K\u0001\u0000\u0000\u0000\u01b6"+
		"\u01b4\u0001\u0000\u0000\u0000\u01b7\u01b8\u0005;\u0000\u0000\u01b8\u01b9"+
		"\u0005E\u0000\u0000\u01b9\u01bd\u0005\f\u0000\u0000\u01ba\u01bc\u0003"+
		"N\'\u0000\u01bb\u01ba\u0001\u0000\u0000\u0000\u01bc\u01bf\u0001\u0000"+
		"\u0000\u0000\u01bd\u01bb\u0001\u0000\u0000\u0000\u01bd\u01be\u0001\u0000"+
		"\u0000\u0000\u01be\u01c0\u0001\u0000\u0000\u0000\u01bf\u01bd\u0001\u0000"+
		"\u0000\u0000\u01c0\u01c1\u0005\r\u0000\u0000\u01c1M\u0001\u0000\u0000"+
		"\u0000\u01c2\u01c3\u0007\u0001\u0000\u0000\u01c3\u01c6\u0005E\u0000\u0000"+
		"\u01c4\u01c5\u0005\u0007\u0000\u0000\u01c5\u01c7\u0003:\u001d\u0000\u01c6"+
		"\u01c4\u0001\u0000\u0000\u0000\u01c6\u01c7\u0001\u0000\u0000\u0000\u01c7"+
		"\u01ca\u0001\u0000\u0000\u0000\u01c8\u01c9\u0005\u0006\u0000\u0000\u01c9"+
		"\u01cb\u0003<\u001e\u0000\u01ca\u01c8\u0001\u0000\u0000\u0000\u01ca\u01cb"+
		"\u0001\u0000\u0000\u0000\u01cb\u01d1\u0001\u0000\u0000\u0000\u01cc\u01ce"+
		"\u0005<\u0000\u0000\u01cd\u01cc\u0001\u0000\u0000\u0000\u01cd\u01ce\u0001"+
		"\u0000\u0000\u0000\u01ce\u01cf\u0001\u0000\u0000\u0000\u01cf\u01d1\u0003"+
		"P(\u0000\u01d0\u01c2\u0001\u0000\u0000\u0000\u01d0\u01cd\u0001\u0000\u0000"+
		"\u0000\u01d1O\u0001\u0000\u0000\u0000\u01d2\u01d3\u00057\u0000\u0000\u01d3"+
		"\u01d4\u0005E\u0000\u0000\u01d4\u01d5\u0005\b\u0000\u0000\u01d5\u01d6"+
		"\u0005\t\u0000\u0000\u01d6\u01da\u0005\f\u0000\u0000\u01d7\u01d9\u0003"+
		"\u0004\u0002\u0000\u01d8\u01d7\u0001\u0000\u0000\u0000\u01d9\u01dc\u0001"+
		"\u0000\u0000\u0000\u01da\u01d8\u0001\u0000\u0000\u0000\u01da\u01db\u0001"+
		"\u0000\u0000\u0000\u01db\u01dd\u0001\u0000\u0000\u0000\u01dc\u01da\u0001"+
		"\u0000\u0000\u0000\u01dd\u01de\u0005\r\u0000\u0000\u01deQ\u0001\u0000"+
		"\u0000\u0000/WYiosy\u0080\u008b\u00a1\u00ac\u00b0\u00c0\u00c3\u00cd\u00d5"+
		"\u00e2\u00ed\u00ff\u0110\u0117\u011a\u0124\u0150\u0161\u0163\u016a\u016f"+
		"\u0174\u0179\u017c\u0180\u0185\u0188\u018c\u0196\u019d\u01a3\u01a6\u01ac"+
		"\u01af\u01b4\u01bd\u01c6\u01ca\u01cd\u01d0\u01da";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}