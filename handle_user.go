package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mentalcaries/go-gator/internal/database"
)

func handleLogin(s *state, cmd command) error {

	if len(cmd.args) < 1 {
		fmt.Println("username is required")
		os.Exit(1)
	}

	userName := cmd.args[0]

	_, err := s.db.GetUserByName(context.Background(), userName)
	if err != nil {
		fmt.Println("You must be registered to login")
		os.Exit(1)
	}

	err = s.config.SetUser(userName)
	if err != nil {
		return errors.New("could not set username")
	}

	fmt.Println("username has been set to", userName)
	return nil
}

func handleRegister(s *state, cmd command) error {

	if len(cmd.args) < 1 {
		fmt.Println("A name is required")
		os.Exit(1)
	}

	name := cmd.args[0]

	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
        ID: uuid.New(), 
        Name: name,
        CreatedAt: time.Now().UTC(), 
        UpdatedAt: time.Now().UTC(),
    })
	if err != nil {
		fmt.Println("User already exists")
		os.Exit(1)
	}

	s.config.SetUser(newUser.Name)
	fmt.Printf("User %s has been created\n", s.config.CurrentUserName)

	return nil
}

func handleResetUsers(s *state, cmd command) error{
    err := s.db.DeleteAllUsers(context.Background())
    if err != nil{
        return fmt.Errorf("could not delete users: %v", err)
    }

    fmt.Println("Users deleted successfully")
    return nil
}

func handleGetAllUsers(s *state, cmd command) error {
    allUsers, err := s.db.GetUsers(context.Background())
    if err != nil {
        return fmt.Errorf("could not get users: %v", err)
    }

    if len(allUsers) < 1{
        fmt.Println("There are no registered users")
    }

    for _, user := range allUsers {
        userItem := "* " + user.Name 
        if s.config.CurrentUserName == user.Name{
            userItem += " (current)"
        }
        fmt.Println(user)
    }
    return nil
}