package Reportes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CreateCST(input string) string {

	// replace the " with \"
	input = strings.ReplaceAll(input, "\"", "\\\"")

	url := "http://lab.antlr.org/parse/"

	// payload
	payload := []byte(`{"grammar":"parser grammar ExprParser;\noptions { tokenVocab=ExprLexer; }\n\ninicio: lista EOF;\n\nlista: (lista_proceso | funciones )*;\n\nlista_proceso: declaracion\n              | asignacion\n              | print\n              | if\n              | switch\n              | while\n              | for\n              | guard\n              | funcvectorList\n              | funcLLamada \n              | retornar\n             ;  \nretornar : RETURN (expresion)? \n         | BREAK\n         | CONTINUE\n         ;\n\ndeclaracion: dec_tipo_valor\n          |  dec_valor\n          | dec_vector\n          | dec_vector_V_C\n          ;\ndec_tipo_valor: op=(VAR|LET) ID  DOSPUNTOS tipo (IGUAL)?  expresion # asigtipo1;\ndec_valor: op=(VAR|LET) ID IGUAL expresion #asigtipo2;\n\nasignacion: asignacionVariable\n          | asignacionVector\n          ;\n\nasignacionVariable: ID op=(INCREMENTO | DECREMENTO |IGUAL)  expresion ;\nasignacionVector: ID OPEN_SQUARE_BRACKET NUMBER CLOSE_SQUARE_BRACKET IGUAL expresion;\n\n\nprint: PRINT OPEN_PAREN expresion (COMA expresion)* CLOSE_PAREN;\n\n\n// -------------------------------------------------------------------------------------------------\n// -------------------------------------------------------------------------------------------------\n// ------------------------------- Setencias de control ---------------------------------------------\n// IF\nif: IF expresion bloque (elseif_)* (else)? ;\nelseif_: ELSE IF expresion bloque;\nelse: ELSE bloque;\n\n// -------------------------------------------------------------------------------------------------\n// SWITCH\nswitch: SWITCH expresion OPEN_BRACKET (caso)+ (default)? CLOSE_BRACKET;\ncaso: CASE expresion DOSPUNTOS (lista_proceso)*;\ndefault: DEFAULT DOSPUNTOS (lista_proceso)*;\n\n// -------------------------------------------------------------------------------------------------\n// while\nwhile: WHILE expresion bloque;\n// -------------------------------------------------------------------------------------------------\n// for\nfor: FOR (ID|GUION_BAJO) IN left=expresion (TRESPUNTOS right=expresion)? bloque;\n//--------------------------------------------------------------------------------------------------\n// guard\nguard: GUARD expresion ELSE OPEN_BRACKET (lista_proceso)*  CLOSE_BRACKET;\n\n//--------------------------------------------------------------------------------------------------\n// --- vectores\ndec_vector: op=(VAR|LET) ID DOSPUNTOS OPEN_SQUARE_BRACKET tipo CLOSE_SQUARE_BRACKET IGUAL  OPEN_SQUARE_BRACKET expresion (COMA expresion)* CLOSE_SQUARE_BRACKET;\ndec_vector_V_C :  op=(VAR|LET) ID DOSPUNTOS OPEN_SQUARE_BRACKET tipo CLOSE_SQUARE_BRACKET IGUAL subVC;\nsubVC: OPEN_SQUARE_BRACKET CLOSE_SQUARE_BRACKET\n    | ID\n    ;\n// funcion vectores\n\nfuncvectorList: ID PUNTO tipevector OPEN_PAREN (AT)? (expresion)? CLOSE_PAREN; \n\ntipevector:   APPEND\n            | REMOVELAST\n            | REMOVE\n            ;\n\n\n\n\nbloque: OPEN_BRACKET (lista_proceso)* CLOSE_BRACKET;\n\ncontrol : RETURN\n        | BREAK\n        | CONTINUE\n        ;\n\ntipo: INT\n    | FLOAT\n    | STRING\n    | BOOL\n    | CHARACTER\n    ;\n    \nexpresion:\n      left=expresion op=(MUL|DIV|MOD)   right=expresion         # OpAritmetico\n    | left=expresion op=(MAS|MENOS) right=expresion             # OpAritmetico\n    | left=expresion op=(IGUAL_IGUAL|DIFERENTE) right=expresion # OpComparativo\n    | left=expresion op=(MAYOR|MENOR|MAYOR_IGUAL|MENOR_IGUAL) right=expresion # OpRelacional       \n    | left=expresion op=(AND|OR) right=expresion                # OpLogico\n    | NOT expresion                                             # OpLogicoNot               \n    | OPEN_PAREN expresion CLOSE_PAREN                      # parentexpr \n    |  ID PUNTO ISEMPTY # VectorVacio\n    | ID PUNTO COUNT   # VectorCount\n    | ID OPEN_SQUARE_BRACKET expresion CLOSE_SQUARE_BRACKET # VectorAsignacion\n    | funcLLamada   # printLLamada   \n    | NULO            # valorsimple    \n    | NUMBER          # NumberValor\n    | STRING_SINTAX   # StringValor\n    | VALORBOOL       # BoolValor\n    | SINVALOR        # SinValor\n    | ID              # IdValor\n    | tipo OPEN_PAREN expresion CLOSE_PAREN # Casteo\n    \n\n    ;\n\n    \n\nfunciones: FUNC ID OPEN_PAREN (parametros)? CLOSE_PAREN (FLECHA tipo)? bloque;\n\nparametros: (existeExInt)? ID DOSPUNTOS (INOUT)?  tipo (COMA (existeExInt)? ID DOSPUNTOS (INOUT)?  tipo)*;\n\nexisteExInt: ID\n            | GUION_BAJO\n            ;\n\nfuncLLamada: ID OPEN_PAREN parametrosLLamada CLOSE_PAREN;\nparametrosLLamada: (ID DOSPUNTOS)? SINVALOR? expresion (COMA (ID DOSPUNTOS)? SINVALOR? expresion)*;",

	"lexgrammar": "// DELETE THIS CONTENT IF YOU PUT COMBINED GRAMMAR IN Parser TAB\nlexer grammar ExprLexer;\n\n// Datos primitivos\nINT: 'Int';\nFLOAT: 'Float';\nSTRING: 'String';\nBOOL: 'Bool';\nCHARACTER: 'Character';\n\n// ---------------------- Symbolos\nIGUAL : '=' ;\nDOSPUNTOS: ':';\nOPEN_PAREN: '(';\nCLOSE_PAREN: ')';\nINCREMENTO: '+=';\nDECREMENTO: '-=';\nOPEN_BRACKET: '{';\nCLOSE_BRACKET: '}';\nOPEN_SQUARE_BRACKET: '[';\nCLOSE_SQUARE_BRACKET: ']';\nCOMA: ',';\nPUNTO: '.';\nGUION_BAJO: '_';\n// --------------------------- Aritmeticos\n\nMAS: '+';\nMENOS: '-';\nMUL: '*';\nDIV: '/';\nMOD: '%';\n// --------------------------- relacionales\nMAYOR: '>';\nMENOR: '<';\nMAYOR_IGUAL: '>=';\nMENOR_IGUAL: '<=';\n// --------------------------- Comparacion\nIGUAL_IGUAL: '==';\nDIFERENTE: '!=';\n// ---------------------------- Logicos\nAND: '&&';\nOR: '||';\nNOT: '!';\n\n// Reservadas\nVAR: 'var';\nLET: 'let';\nPRINT: 'print';\n\n// setencias\nIF: 'if';\nELSE: 'else';\nSWITCH: 'switch';\nCASE: 'case';\nDEFAULT: 'default';\nWHILE: 'while';\nFOR: 'for';\nIN: 'in';\nBREAK: 'break';\nCONTINUE: 'continue';\nRETURN: 'return';\nTRESPUNTOS: '...';\nGUARD: 'guard';\n\n// FUNCIONES VECTORES\nAPPEND: 'append';\nCOUNT: 'count';\nREMOVE: 'remove';\nREMOVELAST: 'removeLast';\nAT: 'at:';\nISEMPTY: 'IsEmpty';\n\n// funciones\nFUNC: 'func';\nINOUT: 'inout';\nFLECHA: '->';\n\n//Casteos\n\n\n//  valores\nNUMBER: '-'? [0-9]+ ('.' [0-9]+)?;\nVALORBOOL: 'true'| 'false';\nNULO: 'nil';\n\nSTRING_SINTAX: '\"' ( ESCAPE_SEQUENCE | ~['\\r\\n\"] )* '\"';\nESCAPE_SEQUENCE: '\\\\'('n' | 'r' | 't' ); \n\nCARACTER: '\"'~[\"]'\"';\nSINVALOR: '?';\n\nID: [a-zA-Z_][a-zA-Z_0-9]* ;\nCOMMENT: '//' ~[\\r\\n]* -> skip;\nCOMMENT_MULT: '/*' .*? '*/' -> skip;\nWS: [ \\t\\r\\n\\f]+ -> skip;",
    "input": "` + input + `",

	"start": "inicio"}`)

	// request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json") // Set the appropriate content type

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	// parse the response body to json
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return ""
	}

	// create a map to store the json
	var data map[string]interface{}

	// unmarshal the json
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return ""
	}

	// fmt.Println("Response Body:", data)

	result := data["result"].(map[string]interface{})

	// fmt.Println("Response Body:", result["svgtree"])

	// create the file
	err = os.WriteFile("cst.svg", []byte(result["svgtree"].(string)), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}

	fmt.Println("File created successfully")

	return result["svgtree"].(string)

}
