# Writing an Interpreter in GO 

## Chapter One

*Defining the lexer*
    Position and readposition in lexer are used to access the characted in input
    by using them as an index. 

    These 'pointers' allow us to look further into the input and see what is 
    coming next so we can look after the current value which is ch.

    readposition always points to whats next, and position always points to
    the character in the into that corresponds with the ch byte, the current 
    character.

### readChar Function 

    Takes an instance of lexer as an arg. Firstly checks the string isn't 
    empty / if all args have been checked. It then assigns l.ch as the next
    position in the string and moves across the string reading each element 
    of the string.

    0 is the ASCII code for null.

    l.position is always assigned the readposition and then readposition is 
    is added to, this ensures that readposition always points to the next 
    position and position always represents the last read position. 

...

### Identifiers and Keyworkds

The lexer needs to recognize whether a character is a letter or not and if so needs to read the rest of the keywork until it encounters a non-letter character. 

Once we've read the identified/keywork we need to find out if its an indintifier or a keywork so we can use the correct token.TokenType. 

The first step is to extent the switch statement within the lexer. 

We extended it with a default branch, this will check and return anytime
the stirng contains a character that isn't a recognised character.
The token.ILLEGAL token was also added which will be returned when we don't
know how to handle the token at all. 

The *isLetter* helper function checks whether a ch is a letter or not and 
provides the basis for what we can or cannot read within our lang. It 
includes all of the letters from a-z and A-Z as well as _. '_' is an 
important addition as it will allow 'variable_names' to be assigned like that 
without throwing an error that it isn't a letter.

The *readidentified* function will keep calling the read char function on
the string being passed to the lexer until a non letter-character has 
been found.

In the *default* section of the switch case we check for the Literals of 
the current token but not the type. Because we now have built-in keywords
like let, foobar, and fn we need to be able to differentiate this between 
a user-defined indentifier. We can do this by determining the type of token. 

*LookUpIdent* checks the keyworks within the token const table to see whether 
the given indenifies is already a keywork and will return the TokenType constant if it is. Otherwise it will return token.IDENT whcich is default TokenType for all user defined identifiers. 

*SkipWhitespace* function tells the lexer that if the token is a whitespace or new line etc that it should just call the readchar function and move onto testing if the next ch is a letter or not. 

Next we added function to check for whether the token was a digit using the 
**isDigit** function. This is almost identical to the **isLetter** function and simply returns whether the character found is between 0 and 9 and if it 
is will return the character, position and call the nextchar function. 



## 1.4 Extending Our Token Set and Lexer

Adding more tokens for extra characters, and the keyworks true, false, if, else
and return 

These are defined as single character tokens, double character tokens and keywords. 

Adding new cases to the test doesn't need to make sense. Because all we want at the moment is the lexer to translate the ch into tokens and not actually parse the code. 

The **peekChar** method looks one ch ahead in the input to test whether it satisfys one of our double character tokens such as ==  and !=. 

## 1.5 Start of a REPL 

REPL stands for Read Eval Print Loop. REPL is present in a lot of languages like Python and Ruby and is sometimes known as 'console' or interactive mode. 

The purpose of REPL is to read the input, interpret and evaluate it and print the results/output. 

# Chapter Two - Parsing

## 2.1 Parsers

    A parser is a software component that takes input data and builds 
    a data structure to give structural representation to an input, 
    check for syntax in the proces, and is often preceded by a lexical
    analyser which creates tokens from a sequence of character inputs. 


