package ast 

import {
	"/home/silliwilli/Desktop/interpreter/Interpreter_GO/token"
}

type Node interface {
	TokenLiteral() string
}

type Statement interface{
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct{
	Token token.Token 
	name *Identifier
	Value Expression
}

func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0 {
		return p.Statement[0].TokenLiteral()
	}else{
		return ""
	}
}

func (ls *LetStatement) statementNode()		{}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

type Identifier struct {
	Token token.Token
	Value string
}

func(i *Identifier) expressionNode()		{}
func(i *Identifier) TokenLiteral() string{ return i.Token.Literal}
