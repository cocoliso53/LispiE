grammar lispie;

// Parser rules
program : sexpr+ EOF;

sexpr
    : '(' (defvar | defconst | let | defun | logical_op | ifExpr | map | filter | event | emit | forLoop | whileLoop | definterface | function_call) ')'
    ;

defvar : 'defvar' IDENTIFIER type ;
defconst : 'defconst' IDENTIFIER expression ;
let : 'let' '(' IDENTIFIER type expression ')' expression ;
defun : 'defun' ':' visibility IDENTIFIER '(' (IDENTIFIER type)* ')' '->' type expression ;

logical_op : and | or | not ;
and : 'and' expression expression ;
or : 'or' expression expression ;
not : 'not' expression ;

ifExpr : 'if' expression expression expression ;

map : 'map' sexpr expression ;
filter : 'filter' sexpr expression ;

event : 'event' IDENTIFIER '(' (IDENTIFIER type)* ')';
emit : 'emit' IDENTIFIER (expression)* ;

forLoop : 'for' '(' IDENTIFIER type expression ')' expression expression expression ;
whileLoop : 'while' expression expression ;

definterface : 'definterface' IDENTIFIER (function_signature)*;
function_signature : 'function' IDENTIFIER '(' (IDENTIFIER type)* ')' '->' type ;

function_call : IDENTIFIER (expression)* ;

type : 'uint256' | 'int256' | 'address' | 'bool' | 'bytes32' | 'string';
visibility : 'external' | 'internal' | 'payable';

expression
    : sexpr
    | IDENTIFIER
    | NUMBER
    | STRING
    ;

// Lexer rules
IDENTIFIER : [a-zA-Z_] [a-zA-Z0-9_]*;
NUMBER : [0-9]+ ('.' [0-9]+)?;
STRING : '"' (~["\r\n])* '"';
WHITESPACE : [ \t\r\n]+ -> skip;
