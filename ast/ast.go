package ast

import "github.com/devasherr/lambda/token"

type Node interface {
	TokenLiteral() string
}

type Statment interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statments []Statment
}

func (p *Program) TokenLiteral() string {
	if len(p.Statments) == 0 {
		return ""
	}

	return p.Statments[0].TokenLiteral()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatment) statementNode()       {}
func (rs *ReturnStatment) TokenLiteral() string { return rs.Token.Literal }
