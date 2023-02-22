# The Complete Syntax of Lua

Here is the complete syntax of Lua in extended BNF. As usual in extended
BNF, {A} means 0 or more As, and \[A\] means an optional A. (For
operator precedences, see [Precedence](/03_the_language/ch04#precedence); for a description of the
terminals Name, Numeral, and LiteralString, see [Lexical Conventions](/03_the_language/ch01#lexical-conventions).)


        chunk ::= block

        block ::= {stat} [retstat]

        stat ::=  ‘;’ | 
             varlist ‘=’ explist | 
             functioncall | 
             label | 
             break | 
             goto Name | 
             do block end | 
             while exp do block end | 
             repeat block until exp | 
             if exp then block {elseif exp then block} [else block] end | 
             for Name ‘=’ exp ‘,’ exp [‘,’ exp] do block end | 
             for namelist in explist do block end | 
             function funcname funcbody | 
             local function Name funcbody | 
             local attnamelist [‘=’ explist] 

        attnamelist ::=  Name attrib {‘,’ Name attrib}

        attrib ::= [‘<’ Name ‘>’]

        retstat ::= return [explist] [‘;’]

        label ::= ‘::’ Name ‘::’

        funcname ::= Name {‘.’ Name} [‘:’ Name]

        varlist ::= var {‘,’ var}

        var ::=  Name | prefixexp ‘[’ exp ‘]’ | prefixexp ‘.’ Name 

        namelist ::= Name {‘,’ Name}

        explist ::= exp {‘,’ exp}

        exp ::=  nil | false | true | Numeral | LiteralString | ‘...’ | functiondef | 
             prefixexp | tableconstructor | exp binop exp | unop exp 

        prefixexp ::= var | functioncall | ‘(’ exp ‘)’

        functioncall ::=  prefixexp args | prefixexp ‘:’ Name args 

        args ::=  ‘(’ [explist] ‘)’ | tableconstructor | LiteralString 

        functiondef ::= function funcbody

        funcbody ::= ‘(’ [parlist] ‘)’ block end

        parlist ::= namelist [‘,’ ‘...’] | ‘...’

        tableconstructor ::= ‘{’ [fieldlist] ‘}’

        fieldlist ::= field {fieldsep field} [fieldsep]

        field ::= ‘[’ exp ‘]’ ‘=’ exp | Name ‘=’ exp | exp

        fieldsep ::= ‘,’ | ‘;’

        binop ::=  ‘+’ | ‘-’ | ‘*’ | ‘/’ | ‘//’ | ‘^’ | ‘%’ | 
             ‘&’ | ‘~’ | ‘|’ | ‘>>’ | ‘<<’ | ‘..’ | 
             ‘<’ | ‘<=’ | ‘>’ | ‘>=’ | ‘==’ | ‘~=’ | 
             and | or

        unop ::= ‘-’ | not | ‘#’ | ‘~’

Last update: Thu Jan 13 11:33:16 UTC 2022
