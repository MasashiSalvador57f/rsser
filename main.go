package main

import (
	"flag"
	"log"
	"os"
)

const (
	registerRSS = "register_rss"
)

func main() {
	var (
		commandName string
	)

	flag.StringVar(&commandName, "command", "", "command name you want to execute")
	flag.StringVar(&commandName, "c", "", "command name you want to execute")
	flag.Parse()

	if len(commandName) <= 0 {
		log.Fatal("command name is required")
	}

	log.Printf("command is %s", commandName)

	switch commandName {
	case registerRSS:
		log.Println("register rss start")
		log.Println("register rss end")
		os.Exit(0)
	}

	log.Fatalf("no command corresponding to given command name : %s", commandName)
}
