package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fakhriaunur/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read Config:\n%+v\n", cfg)

	programState := NewState(&cfg)
	cmds := NewCommands()
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalln("not enough arguments")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	if err := cmds.run(programState, cmd); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Read Config Again:\n%+v\n", cfg)
}
