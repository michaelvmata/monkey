package main

import (
	"fmt"
	"github.com/michaelvmata/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is an implmentation of the Monkey programming language!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
