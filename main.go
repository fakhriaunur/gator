package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/fakhriaunur/gator/internal/config"
	"github.com/fakhriaunur/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read Config:\n%+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := NewState(&cfg, dbQueries)
	cmds := NewCommands()
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
