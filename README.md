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
```bash
    git clone <repository_url>
    cd <your_project_directory_name>
```
2. Run `go install -o .`
3. Create the config file:
```bash
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

## Other Commands

Start the aggregator (checks for updates every 30 s):

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```


- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
