package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mentalcaries/go-gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {

	cfg, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }
    
    appState := state{
        config: &cfg,
    }

    cmds := commands{
        handlers: make(map[string]func(*state, command) error),
    }

    cmds.register("login", handlerLogin)

    commandArgs := os.Args
    if len (commandArgs)  < 2{
        fmt.Println("invalid command")
        os.Exit(1)
    }


    cmd := command{ name: commandArgs[1], args: commandArgs[2:] }
    
    err = cmds.run(&appState, cmd)
    if err != nil{
        log.Fatal(err)
    }
}
