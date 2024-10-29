package main

import (
	"fmt"
	"log"

	"github.com/fakhriaunur/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read Config:\n%+v\n", cfg)

	err = cfg.SetUser("pahri")
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(cfg.DbURL)
	// fmt.Println(cfg.CurrentUserName)
	fmt.Printf("Read Config Again:\n%+v\n", cfg)
}
