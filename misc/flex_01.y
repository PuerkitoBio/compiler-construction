%option noyywrap
%%
username    printf( "%s", getlogin() );

%%
int main()
{
  yylex();
  printf("done\n");
}
