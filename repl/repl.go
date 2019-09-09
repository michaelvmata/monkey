package repl

import (
	"bufio"
	"fmt"
	"github.com/michaelvmata/monkey/lexer"
	"github.com/michaelvmata/monkey/token"
	"io"
)

// PROMPT text for input
const PROMPT = ">>> "

// Start the repl
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
