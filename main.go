package main

import (
	"dolphin/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to the Dolphin programming language, %s!\n",
		user.Username)
	fmt.Printf("Type commands to execute them\n")
	repl.Start(os.Stdin, os.Stdout)
}
