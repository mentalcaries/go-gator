package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mentalcaries/go-gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) < 1 {
		fmt.Println("username is required")
        os.Exit(1)
	}

	userName := cmd.args[0]

	err := s.config.SetUser(userName)
	if err != nil {
		return errors.New("could not set username")
	}

	fmt.Println("username has been set to", userName)

	return nil
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.handlers[cmd.name]
	if !exists {
		return errors.New("invalid command")
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
    if len(name) < 1 || f == nil {
        fmt.Println("invalid name or function")
    }

    c.handlers[name] = f
}
