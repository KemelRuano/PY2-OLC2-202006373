grammar Lexer;

// Datos primitivos
INT: 'Int';
FLOAT: 'Float';
STRING: 'String';
BOOL: 'Bool';
CHARACTER: 'Character';

// ---------------------- Symbolos
IGUAL : '=' ;
DOSPUNTOS: ':';
OPEN_PAREN: '(';
CLOSE_PAREN: ')';
INCREMENTO: '+=';
DECREMENTO: '-=';
OPEN_BRACKET: '{';
CLOSE_BRACKET: '}';
OPEN_SQUARE_BRACKET: '[';
CLOSE_SQUARE_BRACKET: ']';
COMA: ',';
PUNTO: '.';
GUION_BAJO: '_';
// --------------------------- Aritmeticos

MAS: '+';
MENOS: '-';
MUL: '*';
DIV: '/';
MOD: '%';
// --------------------------- relacionales
MAYOR: '>';
MENOR: '<';
MAYOR_IGUAL: '>=';
MENOR_IGUAL: '<=';
// --------------------------- Comparacion
IGUAL_IGUAL: '==';
DIFERENTE: '!=';
// ---------------------------- Logicos
AND: '&&';
OR: '||';
NOT: '!';

// Reservadas
VAR: 'var';
LET: 'let';
PRINT: 'print';

// setencias
IF: 'if';
ELSE: 'else';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
WHILE: 'while';
FOR: 'for';
IN: 'in';
BREAK: 'break';
CONTINUE: 'continue';
RETURN: 'return';
TRESPUNTOS: '...';
GUARD: 'guard';

// FUNCIONES VECTORES
APPEND: 'append';
COUNT: 'count';
REMOVE: 'remove';
REMOVELAST: 'removeLast';
AT: 'at:';
ISEMPTY: 'IsEmpty';

// funciones
FUNC: 'func';
INOUT: 'inout';
FLECHA: '->';
REFERENCIA: '&';
//Casteos


// struct
STRUCT: 'struct';
MUTATING: 'mutating';
SELF: 'self';

//  valores
NUMBER: '-'? [0-9]+ ('.' [0-9]+)?;
VALORBOOL: 'true'| 'false';
NULO: 'nil';

STRING_SINTAX: '"' ( ESCAPE_SEQUENCE | ~['\r\n"] )* '"';
ESCAPE_SEQUENCE: '\\'('n' | 'r' | 't' ); 

CARACTER: '"'~["]'"';
SINVALOR: '?';

ID: [a-zA-Z_][a-zA-Z_0-9]* ;
COMMENT: '//' ~[\r\n]* -> skip;
COMMENT_MULT: '/*' .*? '*/' -> skip;
WS: [ \t\r\n\f]+ -> skip;

