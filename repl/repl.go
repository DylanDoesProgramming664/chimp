package repl

import (
	"chimp/lexer"
	"chimp/token"
	"fmt"

	"github.com/chzyer/readline"
)

const PROMPT = ">> "

func Start() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:            PROMPT,
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		input, err := rl.Readline()
		if err != nil {
			fmt.Println("Keyboard Interrupt")
			return
		}
		if len(input) == 0 {
			input = string(rune(0))
		}
		l := lexer.New(input, "stdin")
		var tok token.Token
		for tok.Type != token.EOF {
			tok = l.NextToken()
			fmt.Printf("Token%+v\n", tok)
		}
		for _, err := range l.Errors {
			fmt.Println(err)
		}
	}
}
