package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/devasherr/lambda/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("<%s> Welcome to lambda REPL. Type 'help' for mor information.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
