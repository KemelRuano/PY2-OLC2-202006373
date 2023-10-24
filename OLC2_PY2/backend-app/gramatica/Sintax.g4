grammar Sintax;
import Lexer;

inicio: lista EOF;

lista: (lista_proceso | funciones )*;

lista_proceso: declaracion
              | asignacion
              | print
              | if
              | switch
              | while
              | for
              | guard
              | funcvectorList
              | funcLLamada 
              | retornar
              | estructs
              | comentarios
             ;  
comentarios: COMMENT
            | COMMENT_MULT
            ;
retornar : RETURN (expresion)? 
         | BREAK
         | CONTINUE
         ;

declaracion: dec_tipo_valor
          |  dec_valor
          | dec_vector
          | dec_vector_V_C
          ;
dec_tipo_valor: op=(VAR|LET) ID  DOSPUNTOS tipo (IGUAL)?  expresion # DeclaracionTipo;
dec_valor: op=(VAR|LET) ID IGUAL expresion #DeclaracionTipoImplicito;

asignacion: asignacionVariable
          | asignacionVector
          ;

asignacionVariable: ID op=(INCREMENTO | DECREMENTO |IGUAL)  expresion ;
asignacionVector: ID OPEN_SQUARE_BRACKET subasig CLOSE_SQUARE_BRACKET op=(IGUAL|INCREMENTO) expresion;
subasig: NUMBER
        | ID
        ;


print: PRINT OPEN_PAREN expresion (COMA expresion)* CLOSE_PAREN;


// -------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------
// ------------------------------- Setencias de control ---------------------------------------------
// IF
if: IF expresion bloque (elseif_)* (else)? ;
elseif_: ELSE IF expresion bloque;
else: ELSE bloque;

// -------------------------------------------------------------------------------------------------
// SWITCH
switch: SWITCH expresion OPEN_BRACKET (caso)+ (default)? CLOSE_BRACKET;
caso: CASE expresion DOSPUNTOS (lista_proceso)*;
default: DEFAULT DOSPUNTOS (lista_proceso)*;

// -------------------------------------------------------------------------------------------------
// while
while: WHILE expresion bloque;
// -------------------------------------------------------------------------------------------------
// for
for: FOR (ID|GUION_BAJO) IN left=expresion (TRESPUNTOS right=expresion)? bloque;
//--------------------------------------------------------------------------------------------------
// guard
guard: GUARD expresion ELSE OPEN_BRACKET (lista_proceso)*  CLOSE_BRACKET;

//--------------------------------------------------------------------------------------------------
// --- vectores
dec_vector: op=(VAR|LET) ID DOSPUNTOS OPEN_SQUARE_BRACKET tipo CLOSE_SQUARE_BRACKET IGUAL  OPEN_SQUARE_BRACKET expresion (COMA expresion)* CLOSE_SQUARE_BRACKET;
dec_vector_V_C :  op=(VAR|LET) ID DOSPUNTOS OPEN_SQUARE_BRACKET tipo CLOSE_SQUARE_BRACKET IGUAL subVC;
subVC: OPEN_SQUARE_BRACKET CLOSE_SQUARE_BRACKET
    | ID
    ;
// funcion vectores

funcvectorList: ID PUNTO tipevector OPEN_PAREN (AT)? (expresion)? CLOSE_PAREN; 

tipevector:   APPEND
            | REMOVELAST
            | REMOVE
            ;




bloque: OPEN_BRACKET (lista_proceso)* CLOSE_BRACKET;

control : RETURN
        | BREAK
        | CONTINUE
        ;

tipo: INT
    | FLOAT
    | STRING
    | BOOL
    | CHARACTER
    ;
    
expresion:
      left=expresion op=(MUL|DIV|MOD)   right=expresion         # OpAritmetico
    | left=expresion op=(MAS|MENOS) right=expresion             # OpAritmetico
    | left=expresion op=(IGUAL_IGUAL|DIFERENTE) right=expresion # OpComparativo
    | left=expresion op=(MAYOR|MENOR|MAYOR_IGUAL|MENOR_IGUAL) right=expresion # OpRelacional       
    | left=expresion op=(AND|OR) right=expresion                # OpLogico
    | NOT expresion                                             # OpLogicoNot               
    | OPEN_PAREN expresion CLOSE_PAREN                      # parentexpr 
    |  ID PUNTO ISEMPTY # VectorVacio
    | ID PUNTO COUNT   # VectorCount
    | ID OPEN_SQUARE_BRACKET expresion CLOSE_SQUARE_BRACKET # VectorAsignacion
    | funcLLamada   # printLLamada   
    | NULO            # valorsimple    
    | NUMBER          # NumberValor
    | STRING_SINTAX   # StringValor
    | VALORBOOL       # BoolValor
    | SINVALOR        # SinValor
    | ID              # IdValor
    | tipo OPEN_PAREN expresion CLOSE_PAREN # Casteo
    | SELF PUNTO ID IGUAL expresion #selfCALL

    ;


funciones: FUNC ID OPEN_PAREN (parametros)? CLOSE_PAREN (FLECHA tipo)? bloque;

parametros: (existeExInt)? ID DOSPUNTOS (tipoinout)? (tipofuncion)? (COMA (existeExInt)? ID DOSPUNTOS (tipoinout)?  (tipofuncion)?)*;

existeExInt: ID
            | GUION_BAJO
            ;
tipofuncion: tipo
            | OPEN_SQUARE_BRACKET tipo CLOSE_SQUARE_BRACKET
            ;
tipoinout : INOUT ;

funcLLamada: ID OPEN_PAREN parametrosLLamada CLOSE_PAREN;
parametrosLLamada: (ID DOSPUNTOS)? REFERENCIA? expresion (COMA (ID DOSPUNTOS)? REFERENCIA? expresion)*;


estructs: STRUCT ID OPEN_BRACKET lista_struct* CLOSE_BRACKET;
lista_struct: (LET|VAR) ID (DOSPUNTOS tipo)? (IGUAL expresion)? 
               | MUTATING? funcionestruct ;

funcionestruct: FUNC ID OPEN_PAREN CLOSE_PAREN OPEN_BRACKET lista_proceso* CLOSE_BRACKET;

                      