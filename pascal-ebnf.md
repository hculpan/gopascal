# Pascal EBNF

## Reference
[EBNF Definition of the Pascal Programming Language](http://www.cs.kent.edu/~durand/CS43101Fall2004/resources/Pascal-EBNF.html)

## Implemented
program =
    program-heading block "." . 
program-heading =
    "program" identifier ";" . 
block =
    statement-part .
statement-part =
    "begin" statement-sequence "end" . 
statement-sequence =
    statement { ";" statement } . 
statement =
    simple-statement . 
simple-statement =
    procedure-statement . 
procedure-statement =
    procedure-identifier [ actual-parameter-list ] . 
procedure-identifier =
    identifier . 
identifier =
    letter { letter | digit } . 
actual-parameter-list =
    "(" actual-parameter { "," actual-parameter } ")" . 
actual-parameter =
    actual-value . 
actual-value =
    expression . 
expression =
    simple-expression . 
simple-expression =
    term  . 
term =
    factor .   
    string  . 
string =
    "'" string-character { string-character } "'" . 
string-character =
    any-character-except-quote | "''" . 