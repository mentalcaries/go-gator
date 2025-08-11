# GO Gator

A CLI application written in Go that allows users to subscribe to RSS feeds, save posts to a database for later perusal.

Once a feed is added, the application can check for posts according to a user-defined period.

The app also supports simple user login and registration, and allows to multiple users to be subscribed to the same feed.

## Stack
- Go
- PostgreSQL
- Goose
- SQLC

## Instructions for Running Locally

*Requirements*
- Go
- PostgresSQL 13+

1. Clone the repository locally and CD into that directory:
```
    git clone <repository_url>
    cd <your_project_directory_name>
```
2. Run `go install -o .`
3. Create the config file:
```
    touch ~/.gatorconfig.json
    echo {
    "db_url": "postgres://devin:@localhost:5432/gator?sslmode=disable",
    "current_user_name": "devin"
    } > ~/.gatorconfig.json
```
4. To use the app, you must register a user (and this automatically logs you in)
`go-gator register dev`
5. Subscribe to a feed by running
`go-gator addfeed <feed name> <feed url>`

