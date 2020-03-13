# gcmm

Reserved words
print
if
else
while
Punctuation and operators
(	+	=	<
)	-	==	>
{	*	!=	<=
}	/	&&	>=
,	%	||	
Other lexical rules
Each number consists of one or more digits, and denotes a non-negative integer.

Each identifier consists of one or more letters that do not form a reserved word. Reserved words and identifiers are case sensitive. That is, if denotes a reserved word, but If and iF and IF are each distinct identifiers.

Whitespace characters include blanks, tabs, line feeds, and carriage returns.

Statements
A program is a sequence of statements. Each statement is one of the following:

statement type	
assignment	identifier = expression
print	print expression1 , expression2 ... , expressionN
selection	if ( expression ) statement1 else statement2
iteration	while ( expression ) statement
compound	{ statement1 statement2 ... statementN }
