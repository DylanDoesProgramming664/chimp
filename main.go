package main

import (
	"chimp/lexer"
	"chimp/repl"
	"chimp/token"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	argv := os.Args
	argc := len(argv)
	switch argc {
	case 1:
		fmt.Println("CLI: no arguments supplied")
		os.Exit(64)
	case 2:
		switch argv[1] {
		case "run":
			fmt.Println("CLI: no filename provided")
			os.Exit(64)
		case "play":
			fmt.Printf("Hello %s! This is the Chimp programming language!\nFeel free to type in commands\n", user.Name)
			repl.Start()
		default:
			fmt.Printf("CLI: unrecognized argument '%s'", argv[1])
			os.Exit(64)
		}
	case 3:
		switch argv[1] {
		case "run":
			var Filename string = argv[2]
			contents, err := os.ReadFile(Filename)
			if err != nil {
				os.Exit(74)
			}
			input := string(contents)
			l := lexer.New(input, Filename)
			var tok token.Token
			for {
				tok = l.NextToken()
				fmt.Printf("Token%+v\n", tok)
				if tok.Type == token.EOF {
					break
				}
			}
			for _, err := range l.Errors {
				fmt.Println(err)
			}
		case "play":
			fmt.Println("CLI: play takes no additional arguments. Moving along.")
			fmt.Printf("Hello %s! This is the Chimp programming language!\nFeel free to type in commands\n", user.Name)
			repl.Start()
		default:
			fmt.Println("CLI: why are we here? just to suffer?")
			os.Exit(64)
		}
	default:
		fmt.Println("CLI: too many arguments")
		os.Exit(64)
	}
}
