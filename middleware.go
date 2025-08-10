package main

import (
	"context"
	"fmt"

	"github.com/mentalcaries/go-gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
    return func(s *state, cmd command) error{
        currentUser, err := s.db.GetUserByName(context.Background(), s.config.CurrentUserName)

        if err != nil{
            return fmt.Errorf("could not get user")
        }
        return handler(s, cmd, currentUser)
    }
}
