  /* (F)lex definition for the simple language */

  /* DECLARATIONS SECTION */

Digit       [0-9]
Char        [a-zA-Z]

  /* 
    option noyywrap tells flex to not look for a yywrap function, avoiding
    gcc error when building with gcc lex.yy.c -lfl:

    ld: library not found for -lfl

    UPDATE: Turns out I don't need -lfl to compile and run this, but option noyywrap
    IS required.
  */
%option noyywrap

%%

  /* RULES SECTION */

  /*


    From http://dinosaur.compilertools.net/lex/index.html:

    Occasionally, a Lex action may decide that a rule has not recognized the correct
    span of characters. Two routines are provided to aid with this situation. First,
    yymore() can be called to indicate that the next input expression recognized is
    to be tacked on to the end of this input. Normally, the next input string would
    overwrite the current entry in yytext. Second, yyless (n) may be called to
    indicate that not all the characters matched by the currently successful
    expression are wanted right now. The argument n indicates the number of
    characters in yytext to be retained. Further characters previously matched are
    returned to the input.

    (See source URL for extremely useful tip for string literals parsing)


    Lex also permits access to the I/O routines it uses. They are:
    1) input() which returns the next input character;
    2) output(c) which writes the character c on the output; and
    3) unput(c) pushes the character c back onto the input stream to be read later
       by input().

    Another Lex library routine that the user will sometimes want to redefine is 
    yywrap() which is called whenever Lex reaches an end-of-file. If yywrap returns
    a 1, Lex continues with the normal wrapup on end of input. 0 instructs Lex to
    continue processing.

    When more than one expression can match the current input, Lex chooses as follows:
    1) The longest match is preferred.
    2) Among rules which matched the same number of characters, the rule given first
       is preferred.

    The action REJECT means ``go do the next alternative.'' It causes whatever rule
    was second choice after the current rule to be executed. The position of the
    input pointer is adjusted accordingly.

    [...] if Yacc is loaded, and its main program is used, Yacc will call yylex().
    In this case each Lex rule should end with return(token); where the appropriate
    token value is returned. An easy way to get access to Yacc's names for tokens
    is to compile the Lex output file as part of the Yacc output file by placing
    the line # include "lex.yy.c" in the last section of Yacc input.
  */

[ \t\n]                   ; /* Ignore whitespace */
fn                        { printf("FUNC\n"); }
var                       { printf("VAR\n"); }
print                     { printf("PRINT\n"); }
"("                       { printf("LPAREN\n"); }
")"                       { printf("RPAREN\n"); }
"{"                       { printf("LBRACE\n"); }
"}"                       { printf("RBRACE\n"); }
","                       { printf("COMMA\n"); }
":"                       { printf("COLON\n"); }
"+"                       { printf("ADD\n"); }
"-"                       { printf("SUB\n"); }
"*"                       { printf("MUL\n"); }
{Digit}+                  { printf("NUM: %s\n", yytext); }
{Char}({Char}|{Digit}|_)* { printf("IDENT: %s, length: %d\n", yytext, yyleng); }
.                         { printf("ILLEGAL: %s\n", yytext); }

%%

  /* USER ROUTINES SECTION */

main()
{
  yylex();
  printf("done\n");
}
