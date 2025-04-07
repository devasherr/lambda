package parser

import (
	"github.com/devasherr/lambda/ast"
	"github.com/devasherr/lambda/lexer"
	"github.com/devasherr/lambda/token"
)

type Parse struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parse {
	p := &Parse{l: l}

	// read two tokens, so curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parse) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parse) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parse) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parse) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parse) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parse) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parse) expectPeek(t token.TokenType) bool {
	if !p.peekTokenIs(t) {
		return false
	}

	p.nextToken()
	return true
}
