package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/mentalcaries/go-gator/internal/config"
	"github.com/mentalcaries/go-gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState := &state{
		config: &cfg,
	}

    db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Println("error connecting to DB")
		os.Exit(1)
	}
	dbQueries := database.New(db)

	appState.db = dbQueries
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleResetUsers)
    cmds.register("users", handleGetAllUsers)
    cmds.register("agg", handleFetchFeed)
    cmds.register("addfeed", handleAddFeed)
    cmds.register("feeds", handleAllFeeds)

	commandArgs := os.Args
	if len(commandArgs) < 2 {
		log.Fatal("invalid command")
	}

	cmd := command{name: commandArgs[1], args: commandArgs[2:]}

	err = cmds.run(appState, cmd)
	if err != nil {
		log.Fatal(err)
	}

	
}
