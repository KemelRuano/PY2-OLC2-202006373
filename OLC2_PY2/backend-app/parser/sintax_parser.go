// Code generated from C:\Users\LENOVO\Desktop\PY2-OLC2-202006373\OLC2_PY2\backend-app\gramatica\Sintax.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Sintax
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SintaxParser struct {
	*antlr.BaseParser
}

var SintaxParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sintaxParserInit() {
	staticData := &SintaxParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'='",
		"':'", "'('", "')'", "'+='", "'-='", "'{'", "'}'", "'['", "']'", "','",
		"'.'", "'_'", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'>='",
		"'<='", "'=='", "'!='", "'&&'", "'||'", "'!'", "'var'", "'let'", "'print'",
		"'if'", "'else'", "'switch'", "'case'", "'default'", "'while'", "'for'",
		"'in'", "'break'", "'continue'", "'return'", "'...'", "'guard'", "'append'",
		"'count'", "'remove'", "'removeLast'", "'at:'", "'IsEmpty'", "'func'",
		"'inout'", "'->'", "'&'", "'struct'", "'mutating'", "'self'", "", "",
		"'nil'", "", "", "", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "IGUAL", "DOSPUNTOS",
		"OPEN_PAREN", "CLOSE_PAREN", "INCREMENTO", "DECREMENTO", "OPEN_BRACKET",
		"CLOSE_BRACKET", "OPEN_SQUARE_BRACKET", "CLOSE_SQUARE_BRACKET", "COMA",
		"PUNTO", "GUION_BAJO", "MAS", "MENOS", "MUL", "DIV", "MOD", "MAYOR",
		"MENOR", "MAYOR_IGUAL", "MENOR_IGUAL", "IGUAL_IGUAL", "DIFERENTE", "AND",
		"OR", "NOT", "VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE",
		"DEFAULT", "WHILE", "FOR", "IN", "BREAK", "CONTINUE", "RETURN", "TRESPUNTOS",
		"GUARD", "APPEND", "COUNT", "REMOVE", "REMOVELAST", "AT", "ISEMPTY",
		"FUNC", "INOUT", "FLECHA", "REFERENCIA", "STRUCT", "MUTATING", "SELF",
		"NUMBER", "VALORBOOL", "NULO", "STRING_SINTAX", "ESCAPE_SEQUENCE", "CARACTER",
		"SINVALOR", "ID", "COMMENT", "COMMENT_MULT", "WS",
	}
	staticData.RuleNames = []string{
		"inicio", "lista", "lista_proceso", "comentarios", "retornar", "declaracion",
		"dec_tipo_valor", "dec_valor", "asignacion", "asignacionVariable", "asignacionVector",
		"subasig", "print", "if", "elseif_", "else", "switch", "caso", "default",
		"while", "for", "guard", "dec_vector", "dec_vector_V_C", "subVC", "funcvectorList",
		"tipevector", "bloque", "control", "tipo", "expresion", "funciones",
		"parametros", "existeExInt", "tipofuncion", "tipoinout", "funcLLamada",
		"parametrosLLamada", "estructs", "lista_struct", "funcionestruct",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 480, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 5, 1, 88, 8, 1, 10, 1, 12, 1, 91, 9, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 106,
		8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 112, 8, 4, 1, 4, 1, 4, 3, 4, 116, 8,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 122, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 3, 6, 129, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 3, 8, 140, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5,
		12, 160, 8, 12, 10, 12, 12, 12, 163, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 5, 13, 171, 8, 13, 10, 13, 12, 13, 174, 9, 13, 1, 13, 3,
		13, 177, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 4, 16, 191, 8, 16, 11, 16, 12, 16, 192, 1,
		16, 3, 16, 196, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17,
		204, 8, 17, 10, 17, 12, 17, 207, 9, 17, 1, 18, 1, 18, 1, 18, 5, 18, 212,
		8, 18, 10, 18, 12, 18, 215, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 227, 8, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 236, 8, 21, 10, 21, 12, 21, 239, 9,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 5, 22, 254, 8, 22, 10, 22, 12, 22, 257, 9, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 3, 24, 273, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 3, 25, 280, 8, 25, 1, 25, 3, 25, 283, 8, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 5, 27, 291, 8, 27, 10, 27, 12, 27, 294, 9, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 337,
		8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 354, 8, 30, 10, 30, 12, 30,
		357, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 363, 8, 31, 1, 31, 1, 31,
		1, 31, 3, 31, 368, 8, 31, 1, 31, 1, 31, 1, 32, 3, 32, 373, 8, 32, 1, 32,
		1, 32, 1, 32, 3, 32, 378, 8, 32, 1, 32, 3, 32, 381, 8, 32, 1, 32, 1, 32,
		3, 32, 385, 8, 32, 1, 32, 1, 32, 1, 32, 3, 32, 390, 8, 32, 1, 32, 3, 32,
		393, 8, 32, 5, 32, 395, 8, 32, 10, 32, 12, 32, 398, 9, 32, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 407, 8, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 3, 36, 414, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37, 3, 37,
		420, 8, 37, 1, 37, 3, 37, 423, 8, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37,
		429, 8, 37, 1, 37, 3, 37, 432, 8, 37, 1, 37, 5, 37, 435, 8, 37, 10, 37,
		12, 37, 438, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 444, 8, 38, 10,
		38, 12, 38, 447, 9, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39,
		455, 8, 39, 1, 39, 1, 39, 3, 39, 459, 8, 39, 1, 39, 3, 39, 462, 8, 39,
		1, 39, 3, 39, 465, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 5,
		40, 473, 8, 40, 10, 40, 12, 40, 476, 9, 40, 1, 40, 1, 40, 1, 40, 0, 1,
		60, 41, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 0, 14, 1, 0, 70, 71, 1, 0, 33, 34, 2, 0, 6, 6,
		10, 11, 2, 0, 6, 6, 10, 10, 2, 0, 62, 62, 69, 69, 2, 0, 18, 18, 69, 69,
		2, 0, 49, 49, 51, 52, 1, 0, 44, 46, 1, 0, 1, 5, 1, 0, 21, 23, 1, 0, 19,
		20, 1, 0, 28, 29, 1, 0, 24, 27, 1, 0, 30, 31, 514, 0, 82, 1, 0, 0, 0, 2,
		89, 1, 0, 0, 0, 4, 105, 1, 0, 0, 0, 6, 107, 1, 0, 0, 0, 8, 115, 1, 0, 0,
		0, 10, 121, 1, 0, 0, 0, 12, 123, 1, 0, 0, 0, 14, 132, 1, 0, 0, 0, 16, 139,
		1, 0, 0, 0, 18, 141, 1, 0, 0, 0, 20, 145, 1, 0, 0, 0, 22, 152, 1, 0, 0,
		0, 24, 154, 1, 0, 0, 0, 26, 166, 1, 0, 0, 0, 28, 178, 1, 0, 0, 0, 30, 183,
		1, 0, 0, 0, 32, 186, 1, 0, 0, 0, 34, 199, 1, 0, 0, 0, 36, 208, 1, 0, 0,
		0, 38, 216, 1, 0, 0, 0, 40, 220, 1, 0, 0, 0, 42, 230, 1, 0, 0, 0, 44, 242,
		1, 0, 0, 0, 46, 260, 1, 0, 0, 0, 48, 272, 1, 0, 0, 0, 50, 274, 1, 0, 0,
		0, 52, 286, 1, 0, 0, 0, 54, 288, 1, 0, 0, 0, 56, 297, 1, 0, 0, 0, 58, 299,
		1, 0, 0, 0, 60, 336, 1, 0, 0, 0, 62, 358, 1, 0, 0, 0, 64, 372, 1, 0, 0,
		0, 66, 399, 1, 0, 0, 0, 68, 406, 1, 0, 0, 0, 70, 408, 1, 0, 0, 0, 72, 410,
		1, 0, 0, 0, 74, 419, 1, 0, 0, 0, 76, 439, 1, 0, 0, 0, 78, 464, 1, 0, 0,
		0, 80, 466, 1, 0, 0, 0, 82, 83, 3, 2, 1, 0, 83, 84, 5, 0, 0, 1, 84, 1,
		1, 0, 0, 0, 85, 88, 3, 4, 2, 0, 86, 88, 3, 62, 31, 0, 87, 85, 1, 0, 0,
		0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90,
		1, 0, 0, 0, 90, 3, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 106, 3, 10, 5, 0,
		93, 106, 3, 16, 8, 0, 94, 106, 3, 24, 12, 0, 95, 106, 3, 26, 13, 0, 96,
		106, 3, 32, 16, 0, 97, 106, 3, 38, 19, 0, 98, 106, 3, 40, 20, 0, 99, 106,
		3, 42, 21, 0, 100, 106, 3, 50, 25, 0, 101, 106, 3, 72, 36, 0, 102, 106,
		3, 8, 4, 0, 103, 106, 3, 76, 38, 0, 104, 106, 3, 6, 3, 0, 105, 92, 1, 0,
		0, 0, 105, 93, 1, 0, 0, 0, 105, 94, 1, 0, 0, 0, 105, 95, 1, 0, 0, 0, 105,
		96, 1, 0, 0, 0, 105, 97, 1, 0, 0, 0, 105, 98, 1, 0, 0, 0, 105, 99, 1, 0,
		0, 0, 105, 100, 1, 0, 0, 0, 105, 101, 1, 0, 0, 0, 105, 102, 1, 0, 0, 0,
		105, 103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 5, 1, 0, 0, 0, 107, 108,
		7, 0, 0, 0, 108, 7, 1, 0, 0, 0, 109, 111, 5, 46, 0, 0, 110, 112, 3, 60,
		30, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 116, 1, 0, 0, 0,
		113, 116, 5, 44, 0, 0, 114, 116, 5, 45, 0, 0, 115, 109, 1, 0, 0, 0, 115,
		113, 1, 0, 0, 0, 115, 114, 1, 0, 0, 0, 116, 9, 1, 0, 0, 0, 117, 122, 3,
		12, 6, 0, 118, 122, 3, 14, 7, 0, 119, 122, 3, 44, 22, 0, 120, 122, 3, 46,
		23, 0, 121, 117, 1, 0, 0, 0, 121, 118, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		121, 120, 1, 0, 0, 0, 122, 11, 1, 0, 0, 0, 123, 124, 7, 1, 0, 0, 124, 125,
		5, 69, 0, 0, 125, 126, 5, 7, 0, 0, 126, 128, 3, 58, 29, 0, 127, 129, 5,
		6, 0, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 1, 0, 0,
		0, 130, 131, 3, 60, 30, 0, 131, 13, 1, 0, 0, 0, 132, 133, 7, 1, 0, 0, 133,
		134, 5, 69, 0, 0, 134, 135, 5, 6, 0, 0, 135, 136, 3, 60, 30, 0, 136, 15,
		1, 0, 0, 0, 137, 140, 3, 18, 9, 0, 138, 140, 3, 20, 10, 0, 139, 137, 1,
		0, 0, 0, 139, 138, 1, 0, 0, 0, 140, 17, 1, 0, 0, 0, 141, 142, 5, 69, 0,
		0, 142, 143, 7, 2, 0, 0, 143, 144, 3, 60, 30, 0, 144, 19, 1, 0, 0, 0, 145,
		146, 5, 69, 0, 0, 146, 147, 5, 14, 0, 0, 147, 148, 3, 22, 11, 0, 148, 149,
		5, 15, 0, 0, 149, 150, 7, 3, 0, 0, 150, 151, 3, 60, 30, 0, 151, 21, 1,
		0, 0, 0, 152, 153, 7, 4, 0, 0, 153, 23, 1, 0, 0, 0, 154, 155, 5, 35, 0,
		0, 155, 156, 5, 8, 0, 0, 156, 161, 3, 60, 30, 0, 157, 158, 5, 16, 0, 0,
		158, 160, 3, 60, 30, 0, 159, 157, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161,
		159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 161,
		1, 0, 0, 0, 164, 165, 5, 9, 0, 0, 165, 25, 1, 0, 0, 0, 166, 167, 5, 36,
		0, 0, 167, 168, 3, 60, 30, 0, 168, 172, 3, 54, 27, 0, 169, 171, 3, 28,
		14, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175,
		177, 3, 30, 15, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 27,
		1, 0, 0, 0, 178, 179, 5, 37, 0, 0, 179, 180, 5, 36, 0, 0, 180, 181, 3,
		60, 30, 0, 181, 182, 3, 54, 27, 0, 182, 29, 1, 0, 0, 0, 183, 184, 5, 37,
		0, 0, 184, 185, 3, 54, 27, 0, 185, 31, 1, 0, 0, 0, 186, 187, 5, 38, 0,
		0, 187, 188, 3, 60, 30, 0, 188, 190, 5, 12, 0, 0, 189, 191, 3, 34, 17,
		0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192,
		193, 1, 0, 0, 0, 193, 195, 1, 0, 0, 0, 194, 196, 3, 36, 18, 0, 195, 194,
		1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 198, 5, 13,
		0, 0, 198, 33, 1, 0, 0, 0, 199, 200, 5, 39, 0, 0, 200, 201, 3, 60, 30,
		0, 201, 205, 5, 7, 0, 0, 202, 204, 3, 4, 2, 0, 203, 202, 1, 0, 0, 0, 204,
		207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 35, 1,
		0, 0, 0, 207, 205, 1, 0, 0, 0, 208, 209, 5, 40, 0, 0, 209, 213, 5, 7, 0,
		0, 210, 212, 3, 4, 2, 0, 211, 210, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213,
		211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 37, 1, 0, 0, 0, 215, 213, 1,
		0, 0, 0, 216, 217, 5, 41, 0, 0, 217, 218, 3, 60, 30, 0, 218, 219, 3, 54,
		27, 0, 219, 39, 1, 0, 0, 0, 220, 221, 5, 42, 0, 0, 221, 222, 7, 5, 0, 0,
		222, 223, 5, 43, 0, 0, 223, 226, 3, 60, 30, 0, 224, 225, 5, 47, 0, 0, 225,
		227, 3, 60, 30, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228,
		1, 0, 0, 0, 228, 229, 3, 54, 27, 0, 229, 41, 1, 0, 0, 0, 230, 231, 5, 48,
		0, 0, 231, 232, 3, 60, 30, 0, 232, 233, 5, 37, 0, 0, 233, 237, 5, 12, 0,
		0, 234, 236, 3, 4, 2, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 240, 241, 5, 13, 0, 0, 241, 43, 1, 0, 0, 0, 242, 243, 7, 1,
		0, 0, 243, 244, 5, 69, 0, 0, 244, 245, 5, 7, 0, 0, 245, 246, 5, 14, 0,
		0, 246, 247, 3, 58, 29, 0, 247, 248, 5, 15, 0, 0, 248, 249, 5, 6, 0, 0,
		249, 250, 5, 14, 0, 0, 250, 255, 3, 60, 30, 0, 251, 252, 5, 16, 0, 0, 252,
		254, 3, 60, 30, 0, 253, 251, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253,
		1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 258, 1, 0, 0, 0, 257, 255, 1, 0,
		0, 0, 258, 259, 5, 15, 0, 0, 259, 45, 1, 0, 0, 0, 260, 261, 7, 1, 0, 0,
		261, 262, 5, 69, 0, 0, 262, 263, 5, 7, 0, 0, 263, 264, 5, 14, 0, 0, 264,
		265, 3, 58, 29, 0, 265, 266, 5, 15, 0, 0, 266, 267, 5, 6, 0, 0, 267, 268,
		3, 48, 24, 0, 268, 47, 1, 0, 0, 0, 269, 270, 5, 14, 0, 0, 270, 273, 5,
		15, 0, 0, 271, 273, 5, 69, 0, 0, 272, 269, 1, 0, 0, 0, 272, 271, 1, 0,
		0, 0, 273, 49, 1, 0, 0, 0, 274, 275, 5, 69, 0, 0, 275, 276, 5, 17, 0, 0,
		276, 277, 3, 52, 26, 0, 277, 279, 5, 8, 0, 0, 278, 280, 5, 53, 0, 0, 279,
		278, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281, 283,
		3, 60, 30, 0, 282, 281, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 1,
		0, 0, 0, 284, 285, 5, 9, 0, 0, 285, 51, 1, 0, 0, 0, 286, 287, 7, 6, 0,
		0, 287, 53, 1, 0, 0, 0, 288, 292, 5, 12, 0, 0, 289, 291, 3, 4, 2, 0, 290,
		289, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293,
		1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296, 5, 13,
		0, 0, 296, 55, 1, 0, 0, 0, 297, 298, 7, 7, 0, 0, 298, 57, 1, 0, 0, 0, 299,
		300, 7, 8, 0, 0, 300, 59, 1, 0, 0, 0, 301, 302, 6, 30, -1, 0, 302, 303,
		5, 32, 0, 0, 303, 337, 3, 60, 30, 14, 304, 305, 5, 8, 0, 0, 305, 306, 3,
		60, 30, 0, 306, 307, 5, 9, 0, 0, 307, 337, 1, 0, 0, 0, 308, 309, 5, 69,
		0, 0, 309, 310, 5, 17, 0, 0, 310, 337, 5, 54, 0, 0, 311, 312, 5, 69, 0,
		0, 312, 313, 5, 17, 0, 0, 313, 337, 5, 50, 0, 0, 314, 315, 5, 69, 0, 0,
		315, 316, 5, 14, 0, 0, 316, 317, 3, 60, 30, 0, 317, 318, 5, 15, 0, 0, 318,
		337, 1, 0, 0, 0, 319, 337, 3, 72, 36, 0, 320, 337, 5, 64, 0, 0, 321, 337,
		5, 62, 0, 0, 322, 337, 5, 65, 0, 0, 323, 337, 5, 63, 0, 0, 324, 337, 5,
		68, 0, 0, 325, 337, 5, 69, 0, 0, 326, 327, 3, 58, 29, 0, 327, 328, 5, 8,
		0, 0, 328, 329, 3, 60, 30, 0, 329, 330, 5, 9, 0, 0, 330, 337, 1, 0, 0,
		0, 331, 332, 5, 61, 0, 0, 332, 333, 5, 17, 0, 0, 333, 334, 5, 69, 0, 0,
		334, 335, 5, 6, 0, 0, 335, 337, 3, 60, 30, 1, 336, 301, 1, 0, 0, 0, 336,
		304, 1, 0, 0, 0, 336, 308, 1, 0, 0, 0, 336, 311, 1, 0, 0, 0, 336, 314,
		1, 0, 0, 0, 336, 319, 1, 0, 0, 0, 336, 320, 1, 0, 0, 0, 336, 321, 1, 0,
		0, 0, 336, 322, 1, 0, 0, 0, 336, 323, 1, 0, 0, 0, 336, 324, 1, 0, 0, 0,
		336, 325, 1, 0, 0, 0, 336, 326, 1, 0, 0, 0, 336, 331, 1, 0, 0, 0, 337,
		355, 1, 0, 0, 0, 338, 339, 10, 19, 0, 0, 339, 340, 7, 9, 0, 0, 340, 354,
		3, 60, 30, 20, 341, 342, 10, 18, 0, 0, 342, 343, 7, 10, 0, 0, 343, 354,
		3, 60, 30, 19, 344, 345, 10, 17, 0, 0, 345, 346, 7, 11, 0, 0, 346, 354,
		3, 60, 30, 18, 347, 348, 10, 16, 0, 0, 348, 349, 7, 12, 0, 0, 349, 354,
		3, 60, 30, 17, 350, 351, 10, 15, 0, 0, 351, 352, 7, 13, 0, 0, 352, 354,
		3, 60, 30, 16, 353, 338, 1, 0, 0, 0, 353, 341, 1, 0, 0, 0, 353, 344, 1,
		0, 0, 0, 353, 347, 1, 0, 0, 0, 353, 350, 1, 0, 0, 0, 354, 357, 1, 0, 0,
		0, 355, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 61, 1, 0, 0, 0, 357,
		355, 1, 0, 0, 0, 358, 359, 5, 55, 0, 0, 359, 360, 5, 69, 0, 0, 360, 362,
		5, 8, 0, 0, 361, 363, 3, 64, 32, 0, 362, 361, 1, 0, 0, 0, 362, 363, 1,
		0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 367, 5, 9, 0, 0, 365, 366, 5, 57, 0,
		0, 366, 368, 3, 58, 29, 0, 367, 365, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0,
		368, 369, 1, 0, 0, 0, 369, 370, 3, 54, 27, 0, 370, 63, 1, 0, 0, 0, 371,
		373, 3, 66, 33, 0, 372, 371, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 374,
		1, 0, 0, 0, 374, 375, 5, 69, 0, 0, 375, 377, 5, 7, 0, 0, 376, 378, 3, 70,
		35, 0, 377, 376, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 380, 1, 0, 0, 0,
		379, 381, 3, 68, 34, 0, 380, 379, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381,
		396, 1, 0, 0, 0, 382, 384, 5, 16, 0, 0, 383, 385, 3, 66, 33, 0, 384, 383,
		1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 387, 5, 69,
		0, 0, 387, 389, 5, 7, 0, 0, 388, 390, 3, 70, 35, 0, 389, 388, 1, 0, 0,
		0, 389, 390, 1, 0, 0, 0, 390, 392, 1, 0, 0, 0, 391, 393, 3, 68, 34, 0,
		392, 391, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 395, 1, 0, 0, 0, 394,
		382, 1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397,
		1, 0, 0, 0, 397, 65, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399, 400, 7, 5,
		0, 0, 400, 67, 1, 0, 0, 0, 401, 407, 3, 58, 29, 0, 402, 403, 5, 14, 0,
		0, 403, 404, 3, 58, 29, 0, 404, 405, 5, 15, 0, 0, 405, 407, 1, 0, 0, 0,
		406, 401, 1, 0, 0, 0, 406, 402, 1, 0, 0, 0, 407, 69, 1, 0, 0, 0, 408, 409,
		5, 56, 0, 0, 409, 71, 1, 0, 0, 0, 410, 411, 5, 69, 0, 0, 411, 413, 5, 8,
		0, 0, 412, 414, 3, 74, 37, 0, 413, 412, 1, 0, 0, 0, 413, 414, 1, 0, 0,
		0, 414, 415, 1, 0, 0, 0, 415, 416, 5, 9, 0, 0, 416, 73, 1, 0, 0, 0, 417,
		418, 5, 69, 0, 0, 418, 420, 5, 7, 0, 0, 419, 417, 1, 0, 0, 0, 419, 420,
		1, 0, 0, 0, 420, 422, 1, 0, 0, 0, 421, 423, 5, 58, 0, 0, 422, 421, 1, 0,
		0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 436, 3, 60, 30,
		0, 425, 428, 5, 16, 0, 0, 426, 427, 5, 69, 0, 0, 427, 429, 5, 7, 0, 0,
		428, 426, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 431, 1, 0, 0, 0, 430,
		432, 5, 58, 0, 0, 431, 430, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432, 433,
		1, 0, 0, 0, 433, 435, 3, 60, 30, 0, 434, 425, 1, 0, 0, 0, 435, 438, 1,
		0, 0, 0, 436, 434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 75, 1, 0, 0,
		0, 438, 436, 1, 0, 0, 0, 439, 440, 5, 59, 0, 0, 440, 441, 5, 69, 0, 0,
		441, 445, 5, 12, 0, 0, 442, 444, 3, 78, 39, 0, 443, 442, 1, 0, 0, 0, 444,
		447, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 448,
		1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 448, 449, 5, 13, 0, 0, 449, 77, 1, 0,
		0, 0, 450, 451, 7, 1, 0, 0, 451, 454, 5, 69, 0, 0, 452, 453, 5, 7, 0, 0,
		453, 455, 3, 58, 29, 0, 454, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455,
		458, 1, 0, 0, 0, 456, 457, 5, 6, 0, 0, 457, 459, 3, 60, 30, 0, 458, 456,
		1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 465, 1, 0, 0, 0, 460, 462, 5, 60,
		0, 0, 461, 460, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0,
		463, 465, 3, 80, 40, 0, 464, 450, 1, 0, 0, 0, 464, 461, 1, 0, 0, 0, 465,
		79, 1, 0, 0, 0, 466, 467, 5, 55, 0, 0, 467, 468, 5, 69, 0, 0, 468, 469,
		5, 8, 0, 0, 469, 470, 5, 9, 0, 0, 470, 474, 5, 12, 0, 0, 471, 473, 3, 4,
		2, 0, 472, 471, 1, 0, 0, 0, 473, 476, 1, 0, 0, 0, 474, 472, 1, 0, 0, 0,
		474, 475, 1, 0, 0, 0, 475, 477, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 477,
		478, 5, 13, 0, 0, 478, 81, 1, 0, 0, 0, 47, 87, 89, 105, 111, 115, 121,
		128, 139, 161, 172, 176, 192, 195, 205, 213, 226, 237, 255, 272, 279, 282,
		292, 336, 353, 355, 362, 367, 372, 377, 380, 384, 389, 392, 396, 406, 413,
		419, 422, 428, 431, 436, 445, 454, 458, 461, 464, 474,
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

// SintaxParserInit initializes any static state used to implement SintaxParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSintaxParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SintaxParserInit() {
	staticData := &SintaxParserStaticData
	staticData.once.Do(sintaxParserInit)
}

// NewSintaxParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSintaxParser(input antlr.TokenStream) *SintaxParser {
	SintaxParserInit()
	this := new(SintaxParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SintaxParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Sintax.g4"

	return this
}

// SintaxParser tokens.
const (
	SintaxParserEOF                  = antlr.TokenEOF
	SintaxParserINT                  = 1
	SintaxParserFLOAT                = 2
	SintaxParserSTRING               = 3
	SintaxParserBOOL                 = 4
	SintaxParserCHARACTER            = 5
	SintaxParserIGUAL                = 6
	SintaxParserDOSPUNTOS            = 7
	SintaxParserOPEN_PAREN           = 8
	SintaxParserCLOSE_PAREN          = 9
	SintaxParserINCREMENTO           = 10
	SintaxParserDECREMENTO           = 11
	SintaxParserOPEN_BRACKET         = 12
	SintaxParserCLOSE_BRACKET        = 13
	SintaxParserOPEN_SQUARE_BRACKET  = 14
	SintaxParserCLOSE_SQUARE_BRACKET = 15
	SintaxParserCOMA                 = 16
	SintaxParserPUNTO                = 17
	SintaxParserGUION_BAJO           = 18
	SintaxParserMAS                  = 19
	SintaxParserMENOS                = 20
	SintaxParserMUL                  = 21
	SintaxParserDIV                  = 22
	SintaxParserMOD                  = 23
	SintaxParserMAYOR                = 24
	SintaxParserMENOR                = 25
	SintaxParserMAYOR_IGUAL          = 26
	SintaxParserMENOR_IGUAL          = 27
	SintaxParserIGUAL_IGUAL          = 28
	SintaxParserDIFERENTE            = 29
	SintaxParserAND                  = 30
	SintaxParserOR                   = 31
	SintaxParserNOT                  = 32
	SintaxParserVAR                  = 33
	SintaxParserLET                  = 34
	SintaxParserPRINT                = 35
	SintaxParserIF                   = 36
	SintaxParserELSE                 = 37
	SintaxParserSWITCH               = 38
	SintaxParserCASE                 = 39
	SintaxParserDEFAULT              = 40
	SintaxParserWHILE                = 41
	SintaxParserFOR                  = 42
	SintaxParserIN                   = 43
	SintaxParserBREAK                = 44
	SintaxParserCONTINUE             = 45
	SintaxParserRETURN               = 46
	SintaxParserTRESPUNTOS           = 47
	SintaxParserGUARD                = 48
	SintaxParserAPPEND               = 49
	SintaxParserCOUNT                = 50
	SintaxParserREMOVE               = 51
	SintaxParserREMOVELAST           = 52
	SintaxParserAT                   = 53
	SintaxParserISEMPTY              = 54
	SintaxParserFUNC                 = 55
	SintaxParserINOUT                = 56
	SintaxParserFLECHA               = 57
	SintaxParserREFERENCIA           = 58
	SintaxParserSTRUCT               = 59
	SintaxParserMUTATING             = 60
	SintaxParserSELF                 = 61
	SintaxParserNUMBER               = 62
	SintaxParserVALORBOOL            = 63
	SintaxParserNULO                 = 64
	SintaxParserSTRING_SINTAX        = 65
	SintaxParserESCAPE_SEQUENCE      = 66
	SintaxParserCARACTER             = 67
	SintaxParserSINVALOR             = 68
	SintaxParserID                   = 69
	SintaxParserCOMMENT              = 70
	SintaxParserCOMMENT_MULT         = 71
	SintaxParserWS                   = 72
)

// SintaxParser rules.
const (
	SintaxParserRULE_inicio             = 0
	SintaxParserRULE_lista              = 1
	SintaxParserRULE_lista_proceso      = 2
	SintaxParserRULE_comentarios        = 3
	SintaxParserRULE_retornar           = 4
	SintaxParserRULE_declaracion        = 5
	SintaxParserRULE_dec_tipo_valor     = 6
	SintaxParserRULE_dec_valor          = 7
	SintaxParserRULE_asignacion         = 8
	SintaxParserRULE_asignacionVariable = 9
	SintaxParserRULE_asignacionVector   = 10
	SintaxParserRULE_subasig            = 11
	SintaxParserRULE_print              = 12
	SintaxParserRULE_if                 = 13
	SintaxParserRULE_elseif_            = 14
	SintaxParserRULE_else               = 15
	SintaxParserRULE_switch             = 16
	SintaxParserRULE_caso               = 17
	SintaxParserRULE_default            = 18
	SintaxParserRULE_while              = 19
	SintaxParserRULE_for                = 20
	SintaxParserRULE_guard              = 21
	SintaxParserRULE_dec_vector         = 22
	SintaxParserRULE_dec_vector_V_C     = 23
	SintaxParserRULE_subVC              = 24
	SintaxParserRULE_funcvectorList     = 25
	SintaxParserRULE_tipevector         = 26
	SintaxParserRULE_bloque             = 27
	SintaxParserRULE_control            = 28
	SintaxParserRULE_tipo               = 29
	SintaxParserRULE_expresion          = 30
	SintaxParserRULE_funciones          = 31
	SintaxParserRULE_parametros         = 32
	SintaxParserRULE_existeExInt        = 33
	SintaxParserRULE_tipofuncion        = 34
	SintaxParserRULE_tipoinout          = 35
	SintaxParserRULE_funcLLamada        = 36
	SintaxParserRULE_parametrosLLamada  = 37
	SintaxParserRULE_estructs           = 38
	SintaxParserRULE_lista_struct       = 39
	SintaxParserRULE_funcionestruct     = 40
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lista() IListaContext
	EOF() antlr.TerminalNode

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_inicio
	return p
}

func InitEmptyInicioContext(p *InicioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_inicio
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) Lista() IListaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *InicioContext) EOF() antlr.TerminalNode {
	return s.GetToken(SintaxParserEOF, 0)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitInicio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Inicio() (localctx IInicioContext) {
	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SintaxParserRULE_inicio)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Lista()
	}
	{
		p.SetState(83)
		p.Match(SintaxParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaContext is an interface to support dynamic dispatch.
type IListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLista_proceso() []ILista_procesoContext
	Lista_proceso(i int) ILista_procesoContext
	AllFunciones() []IFuncionesContext
	Funciones(i int) IFuncionesContext

	// IsListaContext differentiates from other interfaces.
	IsListaContext()
}

type ListaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaContext() *ListaContext {
	var p = new(ListaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_lista
	return p
}

func InitEmptyListaContext(p *ListaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_lista
}

func (*ListaContext) IsListaContext() {}

func NewListaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaContext {
	var p = new(ListaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_lista

	return p
}

func (s *ListaContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaContext) AllLista_proceso() []ILista_procesoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_procesoContext); ok {
			len++
		}
	}

	tst := make([]ILista_procesoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_procesoContext); ok {
			tst[i] = t.(ILista_procesoContext)
			i++
		}
	}

	return tst
}

func (s *ListaContext) Lista_proceso(i int) ILista_procesoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_procesoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_procesoContext)
}

func (s *ListaContext) AllFunciones() []IFuncionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncionesContext); ok {
			len++
		}
	}

	tst := make([]IFuncionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncionesContext); ok {
			tst[i] = t.(IFuncionesContext)
			i++
		}
	}

	return tst
}

func (s *ListaContext) Funciones(i int) IFuncionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionesContext)
}

func (s *ListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitLista(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Lista() (localctx IListaContext) {
	localctx = NewListaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SintaxParserRULE_lista)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&481107688239) != 0 {
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SintaxParserVAR, SintaxParserLET, SintaxParserPRINT, SintaxParserIF, SintaxParserSWITCH, SintaxParserWHILE, SintaxParserFOR, SintaxParserBREAK, SintaxParserCONTINUE, SintaxParserRETURN, SintaxParserGUARD, SintaxParserSTRUCT, SintaxParserID, SintaxParserCOMMENT, SintaxParserCOMMENT_MULT:
			{
				p.SetState(85)
				p.Lista_proceso()
			}

		case SintaxParserFUNC:
			{
				p.SetState(86)
				p.Funciones()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_procesoContext is an interface to support dynamic dispatch.
type ILista_procesoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracion() IDeclaracionContext
	Asignacion() IAsignacionContext
	Print_() IPrintContext
	If_() IIfContext
	Switch_() ISwitchContext
	While() IWhileContext
	For_() IForContext
	Guard() IGuardContext
	FuncvectorList() IFuncvectorListContext
	FuncLLamada() IFuncLLamadaContext
	Retornar() IRetornarContext
	Estructs() IEstructsContext
	Comentarios() IComentariosContext

	// IsLista_procesoContext differentiates from other interfaces.
	IsLista_procesoContext()
}

type Lista_procesoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_procesoContext() *Lista_procesoContext {
	var p = new(Lista_procesoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_lista_proceso
	return p
}

func InitEmptyLista_procesoContext(p *Lista_procesoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_lista_proceso
}

func (*Lista_procesoContext) IsLista_procesoContext() {}

func NewLista_procesoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_procesoContext {
	var p = new(Lista_procesoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_lista_proceso

	return p
}

func (s *Lista_procesoContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_procesoContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *Lista_procesoContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *Lista_procesoContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *Lista_procesoContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *Lista_procesoContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
}

func (s *Lista_procesoContext) While() IWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileContext)
}

func (s *Lista_procesoContext) For_() IForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForContext)
}

func (s *Lista_procesoContext) Guard() IGuardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardContext)
}

func (s *Lista_procesoContext) FuncvectorList() IFuncvectorListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncvectorListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncvectorListContext)
}

func (s *Lista_procesoContext) FuncLLamada() IFuncLLamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncLLamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncLLamadaContext)
}

func (s *Lista_procesoContext) Retornar() IRetornarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetornarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetornarContext)
}

func (s *Lista_procesoContext) Estructs() IEstructsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEstructsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEstructsContext)
}

func (s *Lista_procesoContext) Comentarios() IComentariosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComentariosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComentariosContext)
}

func (s *Lista_procesoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_procesoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_procesoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitLista_proceso(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Lista_proceso() (localctx ILista_procesoContext) {
	localctx = NewLista_procesoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SintaxParserRULE_lista_proceso)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Declaracion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Asignacion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Print_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)
			p.If_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(96)
			p.Switch_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(97)
			p.While()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(98)
			p.For_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(99)
			p.Guard()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(100)
			p.FuncvectorList()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(101)
			p.FuncLLamada()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(102)
			p.Retornar()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(103)
			p.Estructs()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(104)
			p.Comentarios()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComentariosContext is an interface to support dynamic dispatch.
type IComentariosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode
	COMMENT_MULT() antlr.TerminalNode

	// IsComentariosContext differentiates from other interfaces.
	IsComentariosContext()
}

type ComentariosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComentariosContext() *ComentariosContext {
	var p = new(ComentariosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_comentarios
	return p
}

func InitEmptyComentariosContext(p *ComentariosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_comentarios
}

func (*ComentariosContext) IsComentariosContext() {}

func NewComentariosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComentariosContext {
	var p = new(ComentariosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_comentarios

	return p
}

func (s *ComentariosContext) GetParser() antlr.Parser { return s.parser }

func (s *ComentariosContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SintaxParserCOMMENT, 0)
}

func (s *ComentariosContext) COMMENT_MULT() antlr.TerminalNode {
	return s.GetToken(SintaxParserCOMMENT_MULT, 0)
}

func (s *ComentariosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComentariosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComentariosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitComentarios(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Comentarios() (localctx IComentariosContext) {
	localctx = NewComentariosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SintaxParserRULE_comentarios)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserCOMMENT || _la == SintaxParserCOMMENT_MULT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRetornarContext is an interface to support dynamic dispatch.
type IRetornarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expresion() IExpresionContext
	BREAK() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode

	// IsRetornarContext differentiates from other interfaces.
	IsRetornarContext()
}

type RetornarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetornarContext() *RetornarContext {
	var p = new(RetornarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_retornar
	return p
}

func InitEmptyRetornarContext(p *RetornarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_retornar
}

func (*RetornarContext) IsRetornarContext() {}

func NewRetornarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetornarContext {
	var p = new(RetornarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_retornar

	return p
}

func (s *RetornarContext) GetParser() antlr.Parser { return s.parser }

func (s *RetornarContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SintaxParserRETURN, 0)
}

func (s *RetornarContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *RetornarContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SintaxParserBREAK, 0)
}

func (s *RetornarContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SintaxParserCONTINUE, 0)
}

func (s *RetornarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetornarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetornarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitRetornar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Retornar() (localctx IRetornarContext) {
	localctx = NewRetornarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SintaxParserRULE_retornar)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SintaxParserRETURN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Match(SintaxParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(110)
				p.expresion(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case SintaxParserBREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.Match(SintaxParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SintaxParserCONTINUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(114)
			p.Match(SintaxParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dec_tipo_valor() IDec_tipo_valorContext
	Dec_valor() IDec_valorContext
	Dec_vector() IDec_vectorContext
	Dec_vector_V_C() IDec_vector_V_CContext

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_declaracion
	return p
}

func InitEmptyDeclaracionContext(p *DeclaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_declaracion
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Dec_tipo_valor() IDec_tipo_valorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDec_tipo_valorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDec_tipo_valorContext)
}

func (s *DeclaracionContext) Dec_valor() IDec_valorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDec_valorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDec_valorContext)
}

func (s *DeclaracionContext) Dec_vector() IDec_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDec_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDec_vectorContext)
}

func (s *DeclaracionContext) Dec_vector_V_C() IDec_vector_V_CContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDec_vector_V_CContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDec_vector_V_CContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitDeclaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SintaxParserRULE_declaracion)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Dec_tipo_valor()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Dec_valor()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.Dec_vector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.Dec_vector_V_C()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDec_tipo_valorContext is an interface to support dynamic dispatch.
type IDec_tipo_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDec_tipo_valorContext differentiates from other interfaces.
	IsDec_tipo_valorContext()
}

type Dec_tipo_valorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDec_tipo_valorContext() *Dec_tipo_valorContext {
	var p = new(Dec_tipo_valorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_dec_tipo_valor
	return p
}

func InitEmptyDec_tipo_valorContext(p *Dec_tipo_valorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_dec_tipo_valor
}

func (*Dec_tipo_valorContext) IsDec_tipo_valorContext() {}

func NewDec_tipo_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_tipo_valorContext {
	var p = new(Dec_tipo_valorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_dec_tipo_valor

	return p
}

func (s *Dec_tipo_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_tipo_valorContext) CopyAll(ctx *Dec_tipo_valorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Dec_tipo_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_tipo_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclaracionTipoContext struct {
	Dec_tipo_valorContext
	op antlr.Token
}

func NewDeclaracionTipoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclaracionTipoContext {
	var p = new(DeclaracionTipoContext)

	InitEmptyDec_tipo_valorContext(&p.Dec_tipo_valorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Dec_tipo_valorContext))

	return p
}

func (s *DeclaracionTipoContext) GetOp() antlr.Token { return s.op }

func (s *DeclaracionTipoContext) SetOp(v antlr.Token) { s.op = v }

func (s *DeclaracionTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionTipoContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *DeclaracionTipoContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SintaxParserDOSPUNTOS, 0)
}

func (s *DeclaracionTipoContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclaracionTipoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclaracionTipoContext) VAR() antlr.TerminalNode {
	return s.GetToken(SintaxParserVAR, 0)
}

func (s *DeclaracionTipoContext) LET() antlr.TerminalNode {
	return s.GetToken(SintaxParserLET, 0)
}

func (s *DeclaracionTipoContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL, 0)
}

func (s *DeclaracionTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitDeclaracionTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Dec_tipo_valor() (localctx IDec_tipo_valorContext) {
	localctx = NewDec_tipo_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SintaxParserRULE_dec_tipo_valor)
	var _la int

	localctx = NewDeclaracionTipoContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*DeclaracionTipoContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserVAR || _la == SintaxParserLET) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*DeclaracionTipoContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(124)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(SintaxParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Tipo()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserIGUAL {
		{
			p.SetState(127)
			p.Match(SintaxParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(130)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDec_valorContext is an interface to support dynamic dispatch.
type IDec_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDec_valorContext differentiates from other interfaces.
	IsDec_valorContext()
}

type Dec_valorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDec_valorContext() *Dec_valorContext {
	var p = new(Dec_valorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_dec_valor
	return p
}

func InitEmptyDec_valorContext(p *Dec_valorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_dec_valor
}

func (*Dec_valorContext) IsDec_valorContext() {}

func NewDec_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_valorContext {
	var p = new(Dec_valorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_dec_valor

	return p
}

func (s *Dec_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_valorContext) CopyAll(ctx *Dec_valorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Dec_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclaracionTipoImplicitoContext struct {
	Dec_valorContext
	op antlr.Token
}

func NewDeclaracionTipoImplicitoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclaracionTipoImplicitoContext {
	var p = new(DeclaracionTipoImplicitoContext)

	InitEmptyDec_valorContext(&p.Dec_valorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Dec_valorContext))

	return p
}

func (s *DeclaracionTipoImplicitoContext) GetOp() antlr.Token { return s.op }

func (s *DeclaracionTipoImplicitoContext) SetOp(v antlr.Token) { s.op = v }

func (s *DeclaracionTipoImplicitoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionTipoImplicitoContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *DeclaracionTipoImplicitoContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL, 0)
}

func (s *DeclaracionTipoImplicitoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclaracionTipoImplicitoContext) VAR() antlr.TerminalNode {
	return s.GetToken(SintaxParserVAR, 0)
}

func (s *DeclaracionTipoImplicitoContext) LET() antlr.TerminalNode {
	return s.GetToken(SintaxParserLET, 0)
}

func (s *DeclaracionTipoImplicitoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitDeclaracionTipoImplicito(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Dec_valor() (localctx IDec_valorContext) {
	localctx = NewDec_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SintaxParserRULE_dec_valor)
	var _la int

	localctx = NewDeclaracionTipoImplicitoContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*DeclaracionTipoImplicitoContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserVAR || _la == SintaxParserLET) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*DeclaracionTipoImplicitoContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(133)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(SintaxParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AsignacionVariable() IAsignacionVariableContext
	AsignacionVector() IAsignacionVectorContext

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) AsignacionVariable() IAsignacionVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionVariableContext)
}

func (s *AsignacionContext) AsignacionVector() IAsignacionVectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionVectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionVectorContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SintaxParserRULE_asignacion)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.AsignacionVariable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.AsignacionVector()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionVariableContext is an interface to support dynamic dispatch.
type IAsignacionVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	Expresion() IExpresionContext
	INCREMENTO() antlr.TerminalNode
	DECREMENTO() antlr.TerminalNode
	IGUAL() antlr.TerminalNode

	// IsAsignacionVariableContext differentiates from other interfaces.
	IsAsignacionVariableContext()
}

type AsignacionVariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAsignacionVariableContext() *AsignacionVariableContext {
	var p = new(AsignacionVariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_asignacionVariable
	return p
}

func InitEmptyAsignacionVariableContext(p *AsignacionVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_asignacionVariable
}

func (*AsignacionVariableContext) IsAsignacionVariableContext() {}

func NewAsignacionVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionVariableContext {
	var p = new(AsignacionVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_asignacionVariable

	return p
}

func (s *AsignacionVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionVariableContext) GetOp() antlr.Token { return s.op }

func (s *AsignacionVariableContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignacionVariableContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *AsignacionVariableContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionVariableContext) INCREMENTO() antlr.TerminalNode {
	return s.GetToken(SintaxParserINCREMENTO, 0)
}

func (s *AsignacionVariableContext) DECREMENTO() antlr.TerminalNode {
	return s.GetToken(SintaxParserDECREMENTO, 0)
}

func (s *AsignacionVariableContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL, 0)
}

func (s *AsignacionVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitAsignacionVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) AsignacionVariable() (localctx IAsignacionVariableContext) {
	localctx = NewAsignacionVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SintaxParserRULE_asignacionVariable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AsignacionVariableContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3136) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AsignacionVariableContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(143)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionVectorContext is an interface to support dynamic dispatch.
type IAsignacionVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	OPEN_SQUARE_BRACKET() antlr.TerminalNode
	Subasig() ISubasigContext
	CLOSE_SQUARE_BRACKET() antlr.TerminalNode
	Expresion() IExpresionContext
	IGUAL() antlr.TerminalNode
	INCREMENTO() antlr.TerminalNode

	// IsAsignacionVectorContext differentiates from other interfaces.
	IsAsignacionVectorContext()
}

type AsignacionVectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAsignacionVectorContext() *AsignacionVectorContext {
	var p = new(AsignacionVectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_asignacionVector
	return p
}

func InitEmptyAsignacionVectorContext(p *AsignacionVectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_asignacionVector
}

func (*AsignacionVectorContext) IsAsignacionVectorContext() {}

func NewAsignacionVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionVectorContext {
	var p = new(AsignacionVectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_asignacionVector

	return p
}

func (s *AsignacionVectorContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionVectorContext) GetOp() antlr.Token { return s.op }

func (s *AsignacionVectorContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignacionVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *AsignacionVectorContext) OPEN_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_SQUARE_BRACKET, 0)
}

func (s *AsignacionVectorContext) Subasig() ISubasigContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubasigContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubasigContext)
}

func (s *AsignacionVectorContext) CLOSE_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_SQUARE_BRACKET, 0)
}

func (s *AsignacionVectorContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionVectorContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL, 0)
}

func (s *AsignacionVectorContext) INCREMENTO() antlr.TerminalNode {
	return s.GetToken(SintaxParserINCREMENTO, 0)
}

func (s *AsignacionVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionVectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitAsignacionVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) AsignacionVector() (localctx IAsignacionVectorContext) {
	localctx = NewAsignacionVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SintaxParserRULE_asignacionVector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(SintaxParserOPEN_SQUARE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Subasig()
	}
	{
		p.SetState(148)
		p.Match(SintaxParserCLOSE_SQUARE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AsignacionVectorContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserIGUAL || _la == SintaxParserINCREMENTO) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AsignacionVectorContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(150)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubasigContext is an interface to support dynamic dispatch.
type ISubasigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsSubasigContext differentiates from other interfaces.
	IsSubasigContext()
}

type SubasigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubasigContext() *SubasigContext {
	var p = new(SubasigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_subasig
	return p
}

func InitEmptySubasigContext(p *SubasigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_subasig
}

func (*SubasigContext) IsSubasigContext() {}

func NewSubasigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubasigContext {
	var p = new(SubasigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_subasig

	return p
}

func (s *SubasigContext) GetParser() antlr.Parser { return s.parser }

func (s *SubasigContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SintaxParserNUMBER, 0)
}

func (s *SubasigContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *SubasigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubasigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubasigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitSubasig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Subasig() (localctx ISubasigContext) {
	localctx = NewSubasigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SintaxParserRULE_subasig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserNUMBER || _la == SintaxParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	CLOSE_PAREN() antlr.TerminalNode
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SintaxParserPRINT, 0)
}

func (s *PrintContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_PAREN, 0)
}

func (s *PrintContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *PrintContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_PAREN, 0)
}

func (s *PrintContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserCOMA)
}

func (s *PrintContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserCOMA, i)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SintaxParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(SintaxParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(SintaxParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.expresion(0)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SintaxParserCOMA {
		{
			p.SetState(157)
			p.Match(SintaxParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.expresion(0)
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
		p.Match(SintaxParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expresion() IExpresionContext
	Bloque() IBloqueContext
	AllElseif_() []IElseif_Context
	Elseif_(i int) IElseif_Context
	Else_() IElseContext

	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(SintaxParserIF, 0)
}

func (s *IfContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IfContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *IfContext) AllElseif_() []IElseif_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseif_Context); ok {
			len++
		}
	}

	tst := make([]IElseif_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseif_Context); ok {
			tst[i] = t.(IElseif_Context)
			i++
		}
	}

	return tst
}

func (s *IfContext) Elseif_(i int) IElseif_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseif_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseif_Context)
}

func (s *IfContext) Else_() IElseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseContext)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SintaxParserRULE_if)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(SintaxParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.expresion(0)
	}
	{
		p.SetState(168)
		p.Bloque()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(169)
				p.Elseif_()
			}

		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserELSE {
		{
			p.SetState(175)
			p.Else_()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseif_Context is an interface to support dynamic dispatch.
type IElseif_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expresion() IExpresionContext
	Bloque() IBloqueContext

	// IsElseif_Context differentiates from other interfaces.
	IsElseif_Context()
}

type Elseif_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseif_Context() *Elseif_Context {
	var p = new(Elseif_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_elseif_
	return p
}

func InitEmptyElseif_Context(p *Elseif_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_elseif_
}

func (*Elseif_Context) IsElseif_Context() {}

func NewElseif_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elseif_Context {
	var p = new(Elseif_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_elseif_

	return p
}

func (s *Elseif_Context) GetParser() antlr.Parser { return s.parser }

func (s *Elseif_Context) ELSE() antlr.TerminalNode {
	return s.GetToken(SintaxParserELSE, 0)
}

func (s *Elseif_Context) IF() antlr.TerminalNode {
	return s.GetToken(SintaxParserIF, 0)
}

func (s *Elseif_Context) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Elseif_Context) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Elseif_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elseif_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elseif_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitElseif_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Elseif_() (localctx IElseif_Context) {
	localctx = NewElseif_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SintaxParserRULE_elseif_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(SintaxParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(SintaxParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.expresion(0)
	}
	{
		p.SetState(181)
		p.Bloque()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseContext is an interface to support dynamic dispatch.
type IElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Bloque() IBloqueContext

	// IsElseContext differentiates from other interfaces.
	IsElseContext()
}

type ElseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseContext() *ElseContext {
	var p = new(ElseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_else
	return p
}

func InitEmptyElseContext(p *ElseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_else
}

func (*ElseContext) IsElseContext() {}

func NewElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseContext {
	var p = new(ElseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_else

	return p
}

func (s *ElseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SintaxParserELSE, 0)
}

func (s *ElseContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitElse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Else_() (localctx IElseContext) {
	localctx = NewElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SintaxParserRULE_else)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(SintaxParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Bloque()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expresion() IExpresionContext
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	AllCaso() []ICasoContext
	Caso(i int) ICasoContext
	Default_() IDefaultContext

	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SintaxParserSWITCH, 0)
}

func (s *SwitchContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SwitchContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_BRACKET, 0)
}

func (s *SwitchContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_BRACKET, 0)
}

func (s *SwitchContext) AllCaso() []ICasoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICasoContext); ok {
			len++
		}
	}

	tst := make([]ICasoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICasoContext); ok {
			tst[i] = t.(ICasoContext)
			i++
		}
	}

	return tst
}

func (s *SwitchContext) Caso(i int) ICasoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasoContext)
}

func (s *SwitchContext) Default_() IDefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SintaxParserRULE_switch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(SintaxParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.expresion(0)
	}
	{
		p.SetState(188)
		p.Match(SintaxParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SintaxParserCASE {
		{
			p.SetState(189)
			p.Caso()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserDEFAULT {
		{
			p.SetState(194)
			p.Default_()
		}

	}
	{
		p.SetState(197)
		p.Match(SintaxParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasoContext is an interface to support dynamic dispatch.
type ICasoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expresion() IExpresionContext
	DOSPUNTOS() antlr.TerminalNode
	AllLista_proceso() []ILista_procesoContext
	Lista_proceso(i int) ILista_procesoContext

	// IsCasoContext differentiates from other interfaces.
	IsCasoContext()
}

type CasoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasoContext() *CasoContext {
	var p = new(CasoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_caso
	return p
}

func InitEmptyCasoContext(p *CasoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_caso
}

func (*CasoContext) IsCasoContext() {}

func NewCasoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasoContext {
	var p = new(CasoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_caso

	return p
}

func (s *CasoContext) GetParser() antlr.Parser { return s.parser }

func (s *CasoContext) CASE() antlr.TerminalNode {
	return s.GetToken(SintaxParserCASE, 0)
}

func (s *CasoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CasoContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SintaxParserDOSPUNTOS, 0)
}

func (s *CasoContext) AllLista_proceso() []ILista_procesoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_procesoContext); ok {
			len++
		}
	}

	tst := make([]ILista_procesoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_procesoContext); ok {
			tst[i] = t.(ILista_procesoContext)
			i++
		}
	}

	return tst
}

func (s *CasoContext) Lista_proceso(i int) ILista_procesoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_procesoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_procesoContext)
}

func (s *CasoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitCaso(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Caso() (localctx ICasoContext) {
	localctx = NewCasoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SintaxParserRULE_caso)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(SintaxParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.expresion(0)
	}
	{
		p.SetState(201)
		p.Match(SintaxParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&481103493935) != 0 {
		{
			p.SetState(202)
			p.Lista_proceso()
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultContext is an interface to support dynamic dispatch.
type IDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	AllLista_proceso() []ILista_procesoContext
	Lista_proceso(i int) ILista_procesoContext

	// IsDefaultContext differentiates from other interfaces.
	IsDefaultContext()
}

type DefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultContext() *DefaultContext {
	var p = new(DefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_default
	return p
}

func InitEmptyDefaultContext(p *DefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_default
}

func (*DefaultContext) IsDefaultContext() {}

func NewDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultContext {
	var p = new(DefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_default

	return p
}

func (s *DefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SintaxParserDEFAULT, 0)
}

func (s *DefaultContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SintaxParserDOSPUNTOS, 0)
}

func (s *DefaultContext) AllLista_proceso() []ILista_procesoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_procesoContext); ok {
			len++
		}
	}

	tst := make([]ILista_procesoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_procesoContext); ok {
			tst[i] = t.(ILista_procesoContext)
			i++
		}
	}

	return tst
}

func (s *DefaultContext) Lista_proceso(i int) ILista_procesoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_procesoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_procesoContext)
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Default_() (localctx IDefaultContext) {
	localctx = NewDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SintaxParserRULE_default)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(SintaxParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(SintaxParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&481103493935) != 0 {
		{
			p.SetState(210)
			p.Lista_proceso()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileContext is an interface to support dynamic dispatch.
type IWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expresion() IExpresionContext
	Bloque() IBloqueContext

	// IsWhileContext differentiates from other interfaces.
	IsWhileContext()
}

type WhileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileContext() *WhileContext {
	var p = new(WhileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_while
	return p
}

func InitEmptyWhileContext(p *WhileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_while
}

func (*WhileContext) IsWhileContext() {}

func NewWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileContext {
	var p = new(WhileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_while

	return p
}

func (s *WhileContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SintaxParserWHILE, 0)
}

func (s *WhileContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *WhileContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) While() (localctx IWhileContext) {
	localctx = NewWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SintaxParserRULE_while)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(SintaxParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.expresion(0)
	}
	{
		p.SetState(218)
		p.Bloque()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForContext is an interface to support dynamic dispatch.
type IForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpresionContext

	// GetRight returns the right rule contexts.
	GetRight() IExpresionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpresionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpresionContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	IN() antlr.TerminalNode
	Bloque() IBloqueContext
	ID() antlr.TerminalNode
	GUION_BAJO() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	TRESPUNTOS() antlr.TerminalNode

	// IsForContext differentiates from other interfaces.
	IsForContext()
}

type ForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpresionContext
	right  IExpresionContext
}

func NewEmptyForContext() *ForContext {
	var p = new(ForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_for
	return p
}

func InitEmptyForContext(p *ForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_for
}

func (*ForContext) IsForContext() {}

func NewForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForContext {
	var p = new(ForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_for

	return p
}

func (s *ForContext) GetParser() antlr.Parser { return s.parser }

func (s *ForContext) GetLeft() IExpresionContext { return s.left }

func (s *ForContext) GetRight() IExpresionContext { return s.right }

func (s *ForContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *ForContext) SetRight(v IExpresionContext) { s.right = v }

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(SintaxParserFOR, 0)
}

func (s *ForContext) IN() antlr.TerminalNode {
	return s.GetToken(SintaxParserIN, 0)
}

func (s *ForContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ForContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *ForContext) GUION_BAJO() antlr.TerminalNode {
	return s.GetToken(SintaxParserGUION_BAJO, 0)
}

func (s *ForContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ForContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForContext) TRESPUNTOS() antlr.TerminalNode {
	return s.GetToken(SintaxParserTRESPUNTOS, 0)
}

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) For_() (localctx IForContext) {
	localctx = NewForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SintaxParserRULE_for)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(SintaxParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserGUION_BAJO || _la == SintaxParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(222)
		p.Match(SintaxParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)

		var _x = p.expresion(0)

		localctx.(*ForContext).left = _x
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserTRESPUNTOS {
		{
			p.SetState(224)
			p.Match(SintaxParserTRESPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)

			var _x = p.expresion(0)

			localctx.(*ForContext).right = _x
		}

	}
	{
		p.SetState(228)
		p.Bloque()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardContext is an interface to support dynamic dispatch.
type IGuardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expresion() IExpresionContext
	ELSE() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	AllLista_proceso() []ILista_procesoContext
	Lista_proceso(i int) ILista_procesoContext

	// IsGuardContext differentiates from other interfaces.
	IsGuardContext()
}

type GuardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardContext() *GuardContext {
	var p = new(GuardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_guard
	return p
}

func InitEmptyGuardContext(p *GuardContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_guard
}

func (*GuardContext) IsGuardContext() {}

func NewGuardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardContext {
	var p = new(GuardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_guard

	return p
}

func (s *GuardContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SintaxParserGUARD, 0)
}

func (s *GuardContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *GuardContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SintaxParserELSE, 0)
}

func (s *GuardContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_BRACKET, 0)
}

func (s *GuardContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_BRACKET, 0)
}

func (s *GuardContext) AllLista_proceso() []ILista_procesoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_procesoContext); ok {
			len++
		}
	}

	tst := make([]ILista_procesoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_procesoContext); ok {
			tst[i] = t.(ILista_procesoContext)
			i++
		}
	}

	return tst
}

func (s *GuardContext) Lista_proceso(i int) ILista_procesoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_procesoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_procesoContext)
}

func (s *GuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitGuard(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Guard() (localctx IGuardContext) {
	localctx = NewGuardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SintaxParserRULE_guard)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(SintaxParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.expresion(0)
	}
	{
		p.SetState(232)
		p.Match(SintaxParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Match(SintaxParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&481103493935) != 0 {
		{
			p.SetState(234)
			p.Lista_proceso()
		}

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(240)
		p.Match(SintaxParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDec_vectorContext is an interface to support dynamic dispatch.
type IDec_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	AllOPEN_SQUARE_BRACKET() []antlr.TerminalNode
	OPEN_SQUARE_BRACKET(i int) antlr.TerminalNode
	Tipo() ITipoContext
	AllCLOSE_SQUARE_BRACKET() []antlr.TerminalNode
	CLOSE_SQUARE_BRACKET(i int) antlr.TerminalNode
	IGUAL() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsDec_vectorContext differentiates from other interfaces.
	IsDec_vectorContext()
}

type Dec_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyDec_vectorContext() *Dec_vectorContext {
	var p = new(Dec_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_dec_vector
	return p
}

func InitEmptyDec_vectorContext(p *Dec_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_dec_vector
}

func (*Dec_vectorContext) IsDec_vectorContext() {}

func NewDec_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_vectorContext {
	var p = new(Dec_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_dec_vector

	return p
}

func (s *Dec_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_vectorContext) GetOp() antlr.Token { return s.op }

func (s *Dec_vectorContext) SetOp(v antlr.Token) { s.op = v }

func (s *Dec_vectorContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *Dec_vectorContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SintaxParserDOSPUNTOS, 0)
}

func (s *Dec_vectorContext) AllOPEN_SQUARE_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserOPEN_SQUARE_BRACKET)
}

func (s *Dec_vectorContext) OPEN_SQUARE_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_SQUARE_BRACKET, i)
}

func (s *Dec_vectorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Dec_vectorContext) AllCLOSE_SQUARE_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserCLOSE_SQUARE_BRACKET)
}

func (s *Dec_vectorContext) CLOSE_SQUARE_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_SQUARE_BRACKET, i)
}

func (s *Dec_vectorContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL, 0)
}

func (s *Dec_vectorContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Dec_vectorContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Dec_vectorContext) VAR() antlr.TerminalNode {
	return s.GetToken(SintaxParserVAR, 0)
}

func (s *Dec_vectorContext) LET() antlr.TerminalNode {
	return s.GetToken(SintaxParserLET, 0)
}

func (s *Dec_vectorContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserCOMA)
}

func (s *Dec_vectorContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserCOMA, i)
}

func (s *Dec_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dec_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitDec_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Dec_vector() (localctx IDec_vectorContext) {
	localctx = NewDec_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SintaxParserRULE_dec_vector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Dec_vectorContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserVAR || _la == SintaxParserLET) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Dec_vectorContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(243)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(SintaxParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(SintaxParserOPEN_SQUARE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Tipo()
	}
	{
		p.SetState(247)
		p.Match(SintaxParserCLOSE_SQUARE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(SintaxParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(SintaxParserOPEN_SQUARE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.expresion(0)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SintaxParserCOMA {
		{
			p.SetState(251)
			p.Match(SintaxParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.expresion(0)
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(258)
		p.Match(SintaxParserCLOSE_SQUARE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDec_vector_V_CContext is an interface to support dynamic dispatch.
type IDec_vector_V_CContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	OPEN_SQUARE_BRACKET() antlr.TerminalNode
	Tipo() ITipoContext
	CLOSE_SQUARE_BRACKET() antlr.TerminalNode
	IGUAL() antlr.TerminalNode
	SubVC() ISubVCContext
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode

	// IsDec_vector_V_CContext differentiates from other interfaces.
	IsDec_vector_V_CContext()
}

type Dec_vector_V_CContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyDec_vector_V_CContext() *Dec_vector_V_CContext {
	var p = new(Dec_vector_V_CContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_dec_vector_V_C
	return p
}

func InitEmptyDec_vector_V_CContext(p *Dec_vector_V_CContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_dec_vector_V_C
}

func (*Dec_vector_V_CContext) IsDec_vector_V_CContext() {}

func NewDec_vector_V_CContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_vector_V_CContext {
	var p = new(Dec_vector_V_CContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_dec_vector_V_C

	return p
}

func (s *Dec_vector_V_CContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_vector_V_CContext) GetOp() antlr.Token { return s.op }

func (s *Dec_vector_V_CContext) SetOp(v antlr.Token) { s.op = v }

func (s *Dec_vector_V_CContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *Dec_vector_V_CContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SintaxParserDOSPUNTOS, 0)
}

func (s *Dec_vector_V_CContext) OPEN_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_SQUARE_BRACKET, 0)
}

func (s *Dec_vector_V_CContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Dec_vector_V_CContext) CLOSE_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_SQUARE_BRACKET, 0)
}

func (s *Dec_vector_V_CContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL, 0)
}

func (s *Dec_vector_V_CContext) SubVC() ISubVCContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubVCContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubVCContext)
}

func (s *Dec_vector_V_CContext) VAR() antlr.TerminalNode {
	return s.GetToken(SintaxParserVAR, 0)
}

func (s *Dec_vector_V_CContext) LET() antlr.TerminalNode {
	return s.GetToken(SintaxParserLET, 0)
}

func (s *Dec_vector_V_CContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_vector_V_CContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dec_vector_V_CContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitDec_vector_V_C(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Dec_vector_V_C() (localctx IDec_vector_V_CContext) {
	localctx = NewDec_vector_V_CContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SintaxParserRULE_dec_vector_V_C)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Dec_vector_V_CContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserVAR || _la == SintaxParserLET) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Dec_vector_V_CContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(261)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Match(SintaxParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(SintaxParserOPEN_SQUARE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Tipo()
	}
	{
		p.SetState(265)
		p.Match(SintaxParserCLOSE_SQUARE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.Match(SintaxParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.SubVC()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubVCContext is an interface to support dynamic dispatch.
type ISubVCContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_SQUARE_BRACKET() antlr.TerminalNode
	CLOSE_SQUARE_BRACKET() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsSubVCContext differentiates from other interfaces.
	IsSubVCContext()
}

type SubVCContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubVCContext() *SubVCContext {
	var p = new(SubVCContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_subVC
	return p
}

func InitEmptySubVCContext(p *SubVCContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_subVC
}

func (*SubVCContext) IsSubVCContext() {}

func NewSubVCContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubVCContext {
	var p = new(SubVCContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_subVC

	return p
}

func (s *SubVCContext) GetParser() antlr.Parser { return s.parser }

func (s *SubVCContext) OPEN_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_SQUARE_BRACKET, 0)
}

func (s *SubVCContext) CLOSE_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_SQUARE_BRACKET, 0)
}

func (s *SubVCContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *SubVCContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubVCContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubVCContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitSubVC(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) SubVC() (localctx ISubVCContext) {
	localctx = NewSubVCContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SintaxParserRULE_subVC)
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SintaxParserOPEN_SQUARE_BRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(269)
			p.Match(SintaxParserOPEN_SQUARE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Match(SintaxParserCLOSE_SQUARE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SintaxParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncvectorListContext is an interface to support dynamic dispatch.
type IFuncvectorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	Tipevector() ITipevectorContext
	OPEN_PAREN() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	AT() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsFuncvectorListContext differentiates from other interfaces.
	IsFuncvectorListContext()
}

type FuncvectorListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncvectorListContext() *FuncvectorListContext {
	var p = new(FuncvectorListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_funcvectorList
	return p
}

func InitEmptyFuncvectorListContext(p *FuncvectorListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_funcvectorList
}

func (*FuncvectorListContext) IsFuncvectorListContext() {}

func NewFuncvectorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncvectorListContext {
	var p = new(FuncvectorListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_funcvectorList

	return p
}

func (s *FuncvectorListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncvectorListContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *FuncvectorListContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SintaxParserPUNTO, 0)
}

func (s *FuncvectorListContext) Tipevector() ITipevectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipevectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipevectorContext)
}

func (s *FuncvectorListContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_PAREN, 0)
}

func (s *FuncvectorListContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_PAREN, 0)
}

func (s *FuncvectorListContext) AT() antlr.TerminalNode {
	return s.GetToken(SintaxParserAT, 0)
}

func (s *FuncvectorListContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *FuncvectorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncvectorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncvectorListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitFuncvectorList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) FuncvectorList() (localctx IFuncvectorListContext) {
	localctx = NewFuncvectorListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SintaxParserRULE_funcvectorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.Match(SintaxParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.Tipevector()
	}
	{
		p.SetState(277)
		p.Match(SintaxParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserAT {
		{
			p.SetState(278)
			p.Match(SintaxParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2305843004918726338) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&51) != 0) {
		{
			p.SetState(281)
			p.expresion(0)
		}

	}
	{
		p.SetState(284)
		p.Match(SintaxParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipevectorContext is an interface to support dynamic dispatch.
type ITipevectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	APPEND() antlr.TerminalNode
	REMOVELAST() antlr.TerminalNode
	REMOVE() antlr.TerminalNode

	// IsTipevectorContext differentiates from other interfaces.
	IsTipevectorContext()
}

type TipevectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipevectorContext() *TipevectorContext {
	var p = new(TipevectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_tipevector
	return p
}

func InitEmptyTipevectorContext(p *TipevectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_tipevector
}

func (*TipevectorContext) IsTipevectorContext() {}

func NewTipevectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipevectorContext {
	var p = new(TipevectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_tipevector

	return p
}

func (s *TipevectorContext) GetParser() antlr.Parser { return s.parser }

func (s *TipevectorContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SintaxParserAPPEND, 0)
}

func (s *TipevectorContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SintaxParserREMOVELAST, 0)
}

func (s *TipevectorContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SintaxParserREMOVE, 0)
}

func (s *TipevectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipevectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipevectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitTipevector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Tipevector() (localctx ITipevectorContext) {
	localctx = NewTipevectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SintaxParserRULE_tipevector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7318349394477056) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	AllLista_proceso() []ILista_procesoContext
	Lista_proceso(i int) ILista_procesoContext

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_bloque
	return p
}

func InitEmptyBloqueContext(p *BloqueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_bloque
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_BRACKET, 0)
}

func (s *BloqueContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_BRACKET, 0)
}

func (s *BloqueContext) AllLista_proceso() []ILista_procesoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_procesoContext); ok {
			len++
		}
	}

	tst := make([]ILista_procesoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_procesoContext); ok {
			tst[i] = t.(ILista_procesoContext)
			i++
		}
	}

	return tst
}

func (s *BloqueContext) Lista_proceso(i int) ILista_procesoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_procesoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_procesoContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitBloque(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SintaxParserRULE_bloque)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(SintaxParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&481103493935) != 0 {
		{
			p.SetState(289)
			p.Lista_proceso()
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(295)
		p.Match(SintaxParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IControlContext is an interface to support dynamic dispatch.
type IControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	BREAK() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode

	// IsControlContext differentiates from other interfaces.
	IsControlContext()
}

type ControlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlContext() *ControlContext {
	var p = new(ControlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_control
	return p
}

func InitEmptyControlContext(p *ControlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_control
}

func (*ControlContext) IsControlContext() {}

func NewControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlContext {
	var p = new(ControlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_control

	return p
}

func (s *ControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SintaxParserRETURN, 0)
}

func (s *ControlContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SintaxParserBREAK, 0)
}

func (s *ControlContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SintaxParserCONTINUE, 0)
}

func (s *ControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Control() (localctx IControlContext) {
	localctx = NewControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SintaxParserRULE_control)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&123145302310912) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) INT() antlr.TerminalNode {
	return s.GetToken(SintaxParserINT, 0)
}

func (s *TipoContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SintaxParserFLOAT, 0)
}

func (s *TipoContext) STRING() antlr.TerminalNode {
	return s.GetToken(SintaxParserSTRING, 0)
}

func (s *TipoContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SintaxParserBOOL, 0)
}

func (s *TipoContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(SintaxParserCHARACTER, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SintaxParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorCountContext struct {
	ExpresionContext
}

func NewVectorCountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorCountContext {
	var p = new(VectorCountContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *VectorCountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorCountContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *VectorCountContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SintaxParserPUNTO, 0)
}

func (s *VectorCountContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SintaxParserCOUNT, 0)
}

func (s *VectorCountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitVectorCount(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpLogicoContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewOpLogicoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpLogicoContext {
	var p = new(OpLogicoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *OpLogicoContext) GetOp() antlr.Token { return s.op }

func (s *OpLogicoContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpLogicoContext) GetLeft() IExpresionContext { return s.left }

func (s *OpLogicoContext) GetRight() IExpresionContext { return s.right }

func (s *OpLogicoContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *OpLogicoContext) SetRight(v IExpresionContext) { s.right = v }

func (s *OpLogicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpLogicoContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *OpLogicoContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *OpLogicoContext) AND() antlr.TerminalNode {
	return s.GetToken(SintaxParserAND, 0)
}

func (s *OpLogicoContext) OR() antlr.TerminalNode {
	return s.GetToken(SintaxParserOR, 0)
}

func (s *OpLogicoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitOpLogico(s)

	default:
		return t.VisitChildren(s)
	}
}

type SinValorContext struct {
	ExpresionContext
}

func NewSinValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SinValorContext {
	var p = new(SinValorContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SinValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinValorContext) SINVALOR() antlr.TerminalNode {
	return s.GetToken(SintaxParserSINVALOR, 0)
}

func (s *SinValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitSinValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorAsignacionContext struct {
	ExpresionContext
}

func NewVectorAsignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorAsignacionContext {
	var p = new(VectorAsignacionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *VectorAsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *VectorAsignacionContext) OPEN_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_SQUARE_BRACKET, 0)
}

func (s *VectorAsignacionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *VectorAsignacionContext) CLOSE_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_SQUARE_BRACKET, 0)
}

func (s *VectorAsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitVectorAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpAritmeticoContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewOpAritmeticoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpAritmeticoContext {
	var p = new(OpAritmeticoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *OpAritmeticoContext) GetOp() antlr.Token { return s.op }

func (s *OpAritmeticoContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpAritmeticoContext) GetLeft() IExpresionContext { return s.left }

func (s *OpAritmeticoContext) GetRight() IExpresionContext { return s.right }

func (s *OpAritmeticoContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *OpAritmeticoContext) SetRight(v IExpresionContext) { s.right = v }

func (s *OpAritmeticoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpAritmeticoContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *OpAritmeticoContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *OpAritmeticoContext) MUL() antlr.TerminalNode {
	return s.GetToken(SintaxParserMUL, 0)
}

func (s *OpAritmeticoContext) DIV() antlr.TerminalNode {
	return s.GetToken(SintaxParserDIV, 0)
}

func (s *OpAritmeticoContext) MOD() antlr.TerminalNode {
	return s.GetToken(SintaxParserMOD, 0)
}

func (s *OpAritmeticoContext) MAS() antlr.TerminalNode {
	return s.GetToken(SintaxParserMAS, 0)
}

func (s *OpAritmeticoContext) MENOS() antlr.TerminalNode {
	return s.GetToken(SintaxParserMENOS, 0)
}

func (s *OpAritmeticoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitOpAritmetico(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberValorContext struct {
	ExpresionContext
}

func NewNumberValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberValorContext {
	var p = new(NumberValorContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *NumberValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValorContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SintaxParserNUMBER, 0)
}

func (s *NumberValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitNumberValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorsimpleContext struct {
	ExpresionContext
}

func NewValorsimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorsimpleContext {
	var p = new(ValorsimpleContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ValorsimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorsimpleContext) NULO() antlr.TerminalNode {
	return s.GetToken(SintaxParserNULO, 0)
}

func (s *ValorsimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitValorsimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type CasteoContext struct {
	ExpresionContext
}

func NewCasteoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CasteoContext {
	var p = new(CasteoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *CasteoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasteoContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *CasteoContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_PAREN, 0)
}

func (s *CasteoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CasteoContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_PAREN, 0)
}

func (s *CasteoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitCasteo(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorVacioContext struct {
	ExpresionContext
}

func NewVectorVacioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorVacioContext {
	var p = new(VectorVacioContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *VectorVacioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorVacioContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *VectorVacioContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SintaxParserPUNTO, 0)
}

func (s *VectorVacioContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SintaxParserISEMPTY, 0)
}

func (s *VectorVacioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitVectorVacio(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentexprContext struct {
	ExpresionContext
}

func NewParentexprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentexprContext {
	var p = new(ParentexprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ParentexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentexprContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_PAREN, 0)
}

func (s *ParentexprContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ParentexprContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_PAREN, 0)
}

func (s *ParentexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitParentexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfCALLContext struct {
	ExpresionContext
}

func NewSelfCALLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfCALLContext {
	var p = new(SelfCALLContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SelfCALLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfCALLContext) SELF() antlr.TerminalNode {
	return s.GetToken(SintaxParserSELF, 0)
}

func (s *SelfCALLContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SintaxParserPUNTO, 0)
}

func (s *SelfCALLContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *SelfCALLContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL, 0)
}

func (s *SelfCALLContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SelfCALLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitSelfCALL(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringValorContext struct {
	ExpresionContext
}

func NewStringValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValorContext {
	var p = new(StringValorContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *StringValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValorContext) STRING_SINTAX() antlr.TerminalNode {
	return s.GetToken(SintaxParserSTRING_SINTAX, 0)
}

func (s *StringValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitStringValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpRelacionalContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewOpRelacionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpRelacionalContext {
	var p = new(OpRelacionalContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *OpRelacionalContext) GetOp() antlr.Token { return s.op }

func (s *OpRelacionalContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpRelacionalContext) GetLeft() IExpresionContext { return s.left }

func (s *OpRelacionalContext) GetRight() IExpresionContext { return s.right }

func (s *OpRelacionalContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *OpRelacionalContext) SetRight(v IExpresionContext) { s.right = v }

func (s *OpRelacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpRelacionalContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *OpRelacionalContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *OpRelacionalContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SintaxParserMAYOR, 0)
}

func (s *OpRelacionalContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SintaxParserMENOR, 0)
}

func (s *OpRelacionalContext) MAYOR_IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserMAYOR_IGUAL, 0)
}

func (s *OpRelacionalContext) MENOR_IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserMENOR_IGUAL, 0)
}

func (s *OpRelacionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitOpRelacional(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolValorContext struct {
	ExpresionContext
}

func NewBoolValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolValorContext {
	var p = new(BoolValorContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *BoolValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValorContext) VALORBOOL() antlr.TerminalNode {
	return s.GetToken(SintaxParserVALORBOOL, 0)
}

func (s *BoolValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitBoolValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintLLamadaContext struct {
	ExpresionContext
}

func NewPrintLLamadaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintLLamadaContext {
	var p = new(PrintLLamadaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *PrintLLamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintLLamadaContext) FuncLLamada() IFuncLLamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncLLamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncLLamadaContext)
}

func (s *PrintLLamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitPrintLLamada(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpLogicoNotContext struct {
	ExpresionContext
}

func NewOpLogicoNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpLogicoNotContext {
	var p = new(OpLogicoNotContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *OpLogicoNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpLogicoNotContext) NOT() antlr.TerminalNode {
	return s.GetToken(SintaxParserNOT, 0)
}

func (s *OpLogicoNotContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *OpLogicoNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitOpLogicoNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdValorContext struct {
	ExpresionContext
}

func NewIdValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdValorContext {
	var p = new(IdValorContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IdValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdValorContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *IdValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitIdValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpComparativoContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewOpComparativoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpComparativoContext {
	var p = new(OpComparativoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *OpComparativoContext) GetOp() antlr.Token { return s.op }

func (s *OpComparativoContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpComparativoContext) GetLeft() IExpresionContext { return s.left }

func (s *OpComparativoContext) GetRight() IExpresionContext { return s.right }

func (s *OpComparativoContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *OpComparativoContext) SetRight(v IExpresionContext) { s.right = v }

func (s *OpComparativoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpComparativoContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *OpComparativoContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *OpComparativoContext) IGUAL_IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL_IGUAL, 0)
}

func (s *OpComparativoContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(SintaxParserDIFERENTE, 0)
}

func (s *OpComparativoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitOpComparativo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *SintaxParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, SintaxParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOpLogicoNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(302)
			p.Match(SintaxParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.expresion(14)
		}

	case 2:
		localctx = NewParentexprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(304)
			p.Match(SintaxParserOPEN_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.expresion(0)
		}
		{
			p.SetState(306)
			p.Match(SintaxParserCLOSE_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVectorVacioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(308)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Match(SintaxParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Match(SintaxParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewVectorCountContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(311)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(SintaxParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Match(SintaxParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewVectorAsignacionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(314)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(SintaxParserOPEN_SQUARE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.expresion(0)
		}
		{
			p.SetState(317)
			p.Match(SintaxParserCLOSE_SQUARE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewPrintLLamadaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(319)
			p.FuncLLamada()
		}

	case 7:
		localctx = NewValorsimpleContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(320)
			p.Match(SintaxParserNULO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewNumberValorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(321)
			p.Match(SintaxParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewStringValorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(322)
			p.Match(SintaxParserSTRING_SINTAX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewBoolValorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(323)
			p.Match(SintaxParserVALORBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewSinValorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(324)
			p.Match(SintaxParserSINVALOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewIdValorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(325)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewCasteoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(326)
			p.Tipo()
		}
		{
			p.SetState(327)
			p.Match(SintaxParserOPEN_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.expresion(0)
		}
		{
			p.SetState(329)
			p.Match(SintaxParserCLOSE_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewSelfCALLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(331)
			p.Match(SintaxParserSELF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(SintaxParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(SintaxParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.expresion(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(353)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpAritmeticoContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*OpAritmeticoContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SintaxParserRULE_expresion)
				p.SetState(338)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(339)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpAritmeticoContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14680064) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpAritmeticoContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(340)

					var _x = p.expresion(20)

					localctx.(*OpAritmeticoContext).right = _x
				}

			case 2:
				localctx = NewOpAritmeticoContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*OpAritmeticoContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SintaxParserRULE_expresion)
				p.SetState(341)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(342)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpAritmeticoContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SintaxParserMAS || _la == SintaxParserMENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpAritmeticoContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(343)

					var _x = p.expresion(19)

					localctx.(*OpAritmeticoContext).right = _x
				}

			case 3:
				localctx = NewOpComparativoContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*OpComparativoContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SintaxParserRULE_expresion)
				p.SetState(344)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(345)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpComparativoContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SintaxParserIGUAL_IGUAL || _la == SintaxParserDIFERENTE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpComparativoContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(346)

					var _x = p.expresion(18)

					localctx.(*OpComparativoContext).right = _x
				}

			case 4:
				localctx = NewOpRelacionalContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*OpRelacionalContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SintaxParserRULE_expresion)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(348)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpRelacionalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&251658240) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpRelacionalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(349)

					var _x = p.expresion(17)

					localctx.(*OpRelacionalContext).right = _x
				}

			case 5:
				localctx = NewOpLogicoContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*OpLogicoContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SintaxParserRULE_expresion)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(351)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpLogicoContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SintaxParserAND || _la == SintaxParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpLogicoContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(352)

					var _x = p.expresion(16)

					localctx.(*OpLogicoContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncionesContext is an interface to support dynamic dispatch.
type IFuncionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	Bloque() IBloqueContext
	Parametros() IParametrosContext
	FLECHA() antlr.TerminalNode
	Tipo() ITipoContext

	// IsFuncionesContext differentiates from other interfaces.
	IsFuncionesContext()
}

type FuncionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncionesContext() *FuncionesContext {
	var p = new(FuncionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_funciones
	return p
}

func InitEmptyFuncionesContext(p *FuncionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_funciones
}

func (*FuncionesContext) IsFuncionesContext() {}

func NewFuncionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionesContext {
	var p = new(FuncionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_funciones

	return p
}

func (s *FuncionesContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionesContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SintaxParserFUNC, 0)
}

func (s *FuncionesContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *FuncionesContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_PAREN, 0)
}

func (s *FuncionesContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_PAREN, 0)
}

func (s *FuncionesContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncionesContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *FuncionesContext) FLECHA() antlr.TerminalNode {
	return s.GetToken(SintaxParserFLECHA, 0)
}

func (s *FuncionesContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *FuncionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitFunciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Funciones() (localctx IFuncionesContext) {
	localctx = NewFuncionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SintaxParserRULE_funciones)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(SintaxParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Match(SintaxParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserGUION_BAJO || _la == SintaxParserID {
		{
			p.SetState(361)
			p.Parametros()
		}

	}
	{
		p.SetState(364)
		p.Match(SintaxParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserFLECHA {
		{
			p.SetState(365)
			p.Match(SintaxParserFLECHA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Tipo()
		}

	}
	{
		p.SetState(369)
		p.Bloque()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllDOSPUNTOS() []antlr.TerminalNode
	DOSPUNTOS(i int) antlr.TerminalNode
	AllExisteExInt() []IExisteExIntContext
	ExisteExInt(i int) IExisteExIntContext
	AllTipoinout() []ITipoinoutContext
	Tipoinout(i int) ITipoinoutContext
	AllTipofuncion() []ITipofuncionContext
	Tipofuncion(i int) ITipofuncionContext
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserID)
}

func (s *ParametrosContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserID, i)
}

func (s *ParametrosContext) AllDOSPUNTOS() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserDOSPUNTOS)
}

func (s *ParametrosContext) DOSPUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserDOSPUNTOS, i)
}

func (s *ParametrosContext) AllExisteExInt() []IExisteExIntContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExisteExIntContext); ok {
			len++
		}
	}

	tst := make([]IExisteExIntContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExisteExIntContext); ok {
			tst[i] = t.(IExisteExIntContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) ExisteExInt(i int) IExisteExIntContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExisteExIntContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExisteExIntContext)
}

func (s *ParametrosContext) AllTipoinout() []ITipoinoutContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITipoinoutContext); ok {
			len++
		}
	}

	tst := make([]ITipoinoutContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITipoinoutContext); ok {
			tst[i] = t.(ITipoinoutContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Tipoinout(i int) ITipoinoutContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoinoutContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoinoutContext)
}

func (s *ParametrosContext) AllTipofuncion() []ITipofuncionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITipofuncionContext); ok {
			len++
		}
	}

	tst := make([]ITipofuncionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITipofuncionContext); ok {
			tst[i] = t.(ITipofuncionContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Tipofuncion(i int) ITipofuncionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipofuncionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipofuncionContext)
}

func (s *ParametrosContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserCOMA)
}

func (s *ParametrosContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserCOMA, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitParametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SintaxParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(372)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(371)
			p.ExisteExInt()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(374)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(375)
		p.Match(SintaxParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserINOUT {
		{
			p.SetState(376)
			p.Tipoinout()
		}

	}
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16446) != 0 {
		{
			p.SetState(379)
			p.Tipofuncion()
		}

	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SintaxParserCOMA {
		{
			p.SetState(382)
			p.Match(SintaxParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(383)
				p.ExisteExInt()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(386)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.Match(SintaxParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SintaxParserINOUT {
			{
				p.SetState(388)
				p.Tipoinout()
			}

		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16446) != 0 {
			{
				p.SetState(391)
				p.Tipofuncion()
			}

		}

		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExisteExIntContext is an interface to support dynamic dispatch.
type IExisteExIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	GUION_BAJO() antlr.TerminalNode

	// IsExisteExIntContext differentiates from other interfaces.
	IsExisteExIntContext()
}

type ExisteExIntContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExisteExIntContext() *ExisteExIntContext {
	var p = new(ExisteExIntContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_existeExInt
	return p
}

func InitEmptyExisteExIntContext(p *ExisteExIntContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_existeExInt
}

func (*ExisteExIntContext) IsExisteExIntContext() {}

func NewExisteExIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExisteExIntContext {
	var p = new(ExisteExIntContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_existeExInt

	return p
}

func (s *ExisteExIntContext) GetParser() antlr.Parser { return s.parser }

func (s *ExisteExIntContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *ExisteExIntContext) GUION_BAJO() antlr.TerminalNode {
	return s.GetToken(SintaxParserGUION_BAJO, 0)
}

func (s *ExisteExIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExisteExIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExisteExIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitExisteExInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) ExisteExInt() (localctx IExisteExIntContext) {
	localctx = NewExisteExIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SintaxParserRULE_existeExInt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SintaxParserGUION_BAJO || _la == SintaxParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipofuncionContext is an interface to support dynamic dispatch.
type ITipofuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo() ITipoContext
	OPEN_SQUARE_BRACKET() antlr.TerminalNode
	CLOSE_SQUARE_BRACKET() antlr.TerminalNode

	// IsTipofuncionContext differentiates from other interfaces.
	IsTipofuncionContext()
}

type TipofuncionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipofuncionContext() *TipofuncionContext {
	var p = new(TipofuncionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_tipofuncion
	return p
}

func InitEmptyTipofuncionContext(p *TipofuncionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_tipofuncion
}

func (*TipofuncionContext) IsTipofuncionContext() {}

func NewTipofuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipofuncionContext {
	var p = new(TipofuncionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_tipofuncion

	return p
}

func (s *TipofuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *TipofuncionContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *TipofuncionContext) OPEN_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_SQUARE_BRACKET, 0)
}

func (s *TipofuncionContext) CLOSE_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_SQUARE_BRACKET, 0)
}

func (s *TipofuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipofuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipofuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitTipofuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Tipofuncion() (localctx ITipofuncionContext) {
	localctx = NewTipofuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SintaxParserRULE_tipofuncion)
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SintaxParserINT, SintaxParserFLOAT, SintaxParserSTRING, SintaxParserBOOL, SintaxParserCHARACTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.Tipo()
		}

	case SintaxParserOPEN_SQUARE_BRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.Match(SintaxParserOPEN_SQUARE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Tipo()
		}
		{
			p.SetState(404)
			p.Match(SintaxParserCLOSE_SQUARE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoinoutContext is an interface to support dynamic dispatch.
type ITipoinoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INOUT() antlr.TerminalNode

	// IsTipoinoutContext differentiates from other interfaces.
	IsTipoinoutContext()
}

type TipoinoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoinoutContext() *TipoinoutContext {
	var p = new(TipoinoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_tipoinout
	return p
}

func InitEmptyTipoinoutContext(p *TipoinoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_tipoinout
}

func (*TipoinoutContext) IsTipoinoutContext() {}

func NewTipoinoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoinoutContext {
	var p = new(TipoinoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_tipoinout

	return p
}

func (s *TipoinoutContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoinoutContext) INOUT() antlr.TerminalNode {
	return s.GetToken(SintaxParserINOUT, 0)
}

func (s *TipoinoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoinoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoinoutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitTipoinout(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Tipoinout() (localctx ITipoinoutContext) {
	localctx = NewTipoinoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SintaxParserRULE_tipoinout)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(408)
		p.Match(SintaxParserINOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncLLamadaContext is an interface to support dynamic dispatch.
type IFuncLLamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	ParametrosLLamada() IParametrosLLamadaContext

	// IsFuncLLamadaContext differentiates from other interfaces.
	IsFuncLLamadaContext()
}

type FuncLLamadaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncLLamadaContext() *FuncLLamadaContext {
	var p = new(FuncLLamadaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_funcLLamada
	return p
}

func InitEmptyFuncLLamadaContext(p *FuncLLamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_funcLLamada
}

func (*FuncLLamadaContext) IsFuncLLamadaContext() {}

func NewFuncLLamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncLLamadaContext {
	var p = new(FuncLLamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_funcLLamada

	return p
}

func (s *FuncLLamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncLLamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *FuncLLamadaContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_PAREN, 0)
}

func (s *FuncLLamadaContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_PAREN, 0)
}

func (s *FuncLLamadaContext) ParametrosLLamada() IParametrosLLamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosLLamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosLLamadaContext)
}

func (s *FuncLLamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncLLamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncLLamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitFuncLLamada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) FuncLLamada() (localctx IFuncLLamadaContext) {
	localctx = NewFuncLLamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SintaxParserRULE_funcLLamada)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(411)
		p.Match(SintaxParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2017612628767014594) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&51) != 0) {
		{
			p.SetState(412)
			p.ParametrosLLamada()
		}

	}
	{
		p.SetState(415)
		p.Match(SintaxParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametrosLLamadaContext is an interface to support dynamic dispatch.
type IParametrosLLamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllDOSPUNTOS() []antlr.TerminalNode
	DOSPUNTOS(i int) antlr.TerminalNode
	AllREFERENCIA() []antlr.TerminalNode
	REFERENCIA(i int) antlr.TerminalNode
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsParametrosLLamadaContext differentiates from other interfaces.
	IsParametrosLLamadaContext()
}

type ParametrosLLamadaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosLLamadaContext() *ParametrosLLamadaContext {
	var p = new(ParametrosLLamadaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_parametrosLLamada
	return p
}

func InitEmptyParametrosLLamadaContext(p *ParametrosLLamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_parametrosLLamada
}

func (*ParametrosLLamadaContext) IsParametrosLLamadaContext() {}

func NewParametrosLLamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosLLamadaContext {
	var p = new(ParametrosLLamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_parametrosLLamada

	return p
}

func (s *ParametrosLLamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosLLamadaContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosLLamadaContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ParametrosLLamadaContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserID)
}

func (s *ParametrosLLamadaContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserID, i)
}

func (s *ParametrosLLamadaContext) AllDOSPUNTOS() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserDOSPUNTOS)
}

func (s *ParametrosLLamadaContext) DOSPUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserDOSPUNTOS, i)
}

func (s *ParametrosLLamadaContext) AllREFERENCIA() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserREFERENCIA)
}

func (s *ParametrosLLamadaContext) REFERENCIA(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserREFERENCIA, i)
}

func (s *ParametrosLLamadaContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SintaxParserCOMA)
}

func (s *ParametrosLLamadaContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SintaxParserCOMA, i)
}

func (s *ParametrosLLamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosLLamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosLLamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitParametrosLLamada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) ParametrosLLamada() (localctx IParametrosLLamadaContext) {
	localctx = NewParametrosLLamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SintaxParserRULE_parametrosLLamada)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(419)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(417)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(SintaxParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SintaxParserREFERENCIA {
		{
			p.SetState(421)
			p.Match(SintaxParserREFERENCIA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(424)
		p.expresion(0)
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SintaxParserCOMA {
		{
			p.SetState(425)
			p.Match(SintaxParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(426)
				p.Match(SintaxParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(427)
				p.Match(SintaxParserDOSPUNTOS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SintaxParserREFERENCIA {
			{
				p.SetState(430)
				p.Match(SintaxParserREFERENCIA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(433)
			p.expresion(0)
		}

		p.SetState(438)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEstructsContext is an interface to support dynamic dispatch.
type IEstructsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	AllLista_struct() []ILista_structContext
	Lista_struct(i int) ILista_structContext

	// IsEstructsContext differentiates from other interfaces.
	IsEstructsContext()
}

type EstructsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEstructsContext() *EstructsContext {
	var p = new(EstructsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_estructs
	return p
}

func InitEmptyEstructsContext(p *EstructsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_estructs
}

func (*EstructsContext) IsEstructsContext() {}

func NewEstructsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EstructsContext {
	var p = new(EstructsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_estructs

	return p
}

func (s *EstructsContext) GetParser() antlr.Parser { return s.parser }

func (s *EstructsContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SintaxParserSTRUCT, 0)
}

func (s *EstructsContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *EstructsContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_BRACKET, 0)
}

func (s *EstructsContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_BRACKET, 0)
}

func (s *EstructsContext) AllLista_struct() []ILista_structContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_structContext); ok {
			len++
		}
	}

	tst := make([]ILista_structContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_structContext); ok {
			tst[i] = t.(ILista_structContext)
			i++
		}
	}

	return tst
}

func (s *EstructsContext) Lista_struct(i int) ILista_structContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_structContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_structContext)
}

func (s *EstructsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EstructsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EstructsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitEstructs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Estructs() (localctx IEstructsContext) {
	localctx = NewEstructsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SintaxParserRULE_estructs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.Match(SintaxParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Match(SintaxParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1188950327395614720) != 0 {
		{
			p.SetState(442)
			p.Lista_struct()
		}

		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(448)
		p.Match(SintaxParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_structContext is an interface to support dynamic dispatch.
type ILista_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LET() antlr.TerminalNode
	VAR() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Tipo() ITipoContext
	IGUAL() antlr.TerminalNode
	Expresion() IExpresionContext
	Funcionestruct() IFuncionestructContext
	MUTATING() antlr.TerminalNode

	// IsLista_structContext differentiates from other interfaces.
	IsLista_structContext()
}

type Lista_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_structContext() *Lista_structContext {
	var p = new(Lista_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_lista_struct
	return p
}

func InitEmptyLista_structContext(p *Lista_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_lista_struct
}

func (*Lista_structContext) IsLista_structContext() {}

func NewLista_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_structContext {
	var p = new(Lista_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_lista_struct

	return p
}

func (s *Lista_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_structContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *Lista_structContext) LET() antlr.TerminalNode {
	return s.GetToken(SintaxParserLET, 0)
}

func (s *Lista_structContext) VAR() antlr.TerminalNode {
	return s.GetToken(SintaxParserVAR, 0)
}

func (s *Lista_structContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SintaxParserDOSPUNTOS, 0)
}

func (s *Lista_structContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Lista_structContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintaxParserIGUAL, 0)
}

func (s *Lista_structContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Lista_structContext) Funcionestruct() IFuncionestructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionestructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionestructContext)
}

func (s *Lista_structContext) MUTATING() antlr.TerminalNode {
	return s.GetToken(SintaxParserMUTATING, 0)
}

func (s *Lista_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_structContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitLista_struct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Lista_struct() (localctx ILista_structContext) {
	localctx = NewLista_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SintaxParserRULE_lista_struct)
	var _la int

	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SintaxParserVAR, SintaxParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(450)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SintaxParserVAR || _la == SintaxParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(451)
			p.Match(SintaxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SintaxParserDOSPUNTOS {
			{
				p.SetState(452)
				p.Match(SintaxParserDOSPUNTOS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(453)
				p.Tipo()
			}

		}
		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SintaxParserIGUAL {
			{
				p.SetState(456)
				p.Match(SintaxParserIGUAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(457)
				p.expresion(0)
			}

		}

	case SintaxParserFUNC, SintaxParserMUTATING:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(461)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SintaxParserMUTATING {
			{
				p.SetState(460)
				p.Match(SintaxParserMUTATING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(463)
			p.Funcionestruct()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncionestructContext is an interface to support dynamic dispatch.
type IFuncionestructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	AllLista_proceso() []ILista_procesoContext
	Lista_proceso(i int) ILista_procesoContext

	// IsFuncionestructContext differentiates from other interfaces.
	IsFuncionestructContext()
}

type FuncionestructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncionestructContext() *FuncionestructContext {
	var p = new(FuncionestructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_funcionestruct
	return p
}

func InitEmptyFuncionestructContext(p *FuncionestructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SintaxParserRULE_funcionestruct
}

func (*FuncionestructContext) IsFuncionestructContext() {}

func NewFuncionestructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionestructContext {
	var p = new(FuncionestructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintaxParserRULE_funcionestruct

	return p
}

func (s *FuncionestructContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionestructContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SintaxParserFUNC, 0)
}

func (s *FuncionestructContext) ID() antlr.TerminalNode {
	return s.GetToken(SintaxParserID, 0)
}

func (s *FuncionestructContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_PAREN, 0)
}

func (s *FuncionestructContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_PAREN, 0)
}

func (s *FuncionestructContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserOPEN_BRACKET, 0)
}

func (s *FuncionestructContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SintaxParserCLOSE_BRACKET, 0)
}

func (s *FuncionestructContext) AllLista_proceso() []ILista_procesoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_procesoContext); ok {
			len++
		}
	}

	tst := make([]ILista_procesoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_procesoContext); ok {
			tst[i] = t.(ILista_procesoContext)
			i++
		}
	}

	return tst
}

func (s *FuncionestructContext) Lista_proceso(i int) ILista_procesoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_procesoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_procesoContext)
}

func (s *FuncionestructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionestructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionestructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SintaxVisitor:
		return t.VisitFuncionestruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SintaxParser) Funcionestruct() (localctx IFuncionestructContext) {
	localctx = NewFuncionestructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SintaxParserRULE_funcionestruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Match(SintaxParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(467)
		p.Match(SintaxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(468)
		p.Match(SintaxParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
		p.Match(SintaxParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(470)
		p.Match(SintaxParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&481103493935) != 0 {
		{
			p.SetState(471)
			p.Lista_proceso()
		}

		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(477)
		p.Match(SintaxParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SintaxParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 30:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SintaxParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
