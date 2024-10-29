package main

import (
	"fmt"
	"os"

	"github.com/fakhriaunur/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Read Config:\n%+v\n", cfg)

	state := NewState(&cfg)
	commands := NewCommands()
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "not enough arguments\n")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	if err := commands.run(state, cmd); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Read Config Again:\n%+v\n", cfg)
}
