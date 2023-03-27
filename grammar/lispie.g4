grammar lispie;

program: contractName sexpr* EOF;

contractName: '#' ID;

sexpr
  : '(' (operator | keyword) args ')' #sexprOp
  | atom                              #sexprAtom
  ;

operator: OP;
keyword: KEYWORD;

args: (atom | sexpr)*;

atom
  : INT
  | FLOAT
  | STRING
  | IDENTIFIER
  | TYPE
  | ID
  | self
  ;

self: 'self';

OP: ('+' | '-' | '*' | '/' | '<' | '>' | '=' | 'and' | 'or' | 'not' | 'if' | 'let' | 'defun' | 'defvar' | 'defconst' | 'definterface' | 'function' | 'event' | 'emit' | 'map' | 'filter' | 'for' | 'while' | ':' KEYWORD);

KEYWORD: ('external' | 'internal' | 'payable');

INT: [0-9]+;
ID: [a-zA-Z_] [a-zA-Z_0-9]*;
FLOAT: [0-9]+ '.' [0-9]+;
STRING: '"' (~["\n\r\t])* '"';
IDENTIFIER: [a-zA-Z_$][a-zA-Z_$0-9]*;
TYPE: ('uint256' | 'int256' | 'address' | 'bool' | 'bytes32' | 'string');
COMMENT: ';;' .*? '\n' -> skip;

WS: [ \t\r\n]+ -> skip;
