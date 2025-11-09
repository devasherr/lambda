package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/devasherr/lambda/evaluator"
	"github.com/devasherr/lambda/lexer"
	"github.com/devasherr/lambda/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErros(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErros(out io.Writer, errors []string) {
	if len(errors) > 0 {
		io.WriteString(out, "parser error:\n")
	}

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
