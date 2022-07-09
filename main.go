package main

import (
	"fmt"
	"os"
	"os/user"
	"simplelang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, this is the SimpleLang programming language.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
