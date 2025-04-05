package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/devasherr/lambda/lexer"
	"github.com/devasherr/lambda/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		tok := l.NextToken()
		for tok.Type != token.EOF {
			fmt.Printf("%+v\n", tok)
			tok = l.NextToken()
		}
	}
}
