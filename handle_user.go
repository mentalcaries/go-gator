package main

import (
	"errors"
	"fmt"
	"os"
)

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