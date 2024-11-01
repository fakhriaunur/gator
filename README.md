# A Blog Aggregator in Go

A multi-player command-line tool for aggregating RSS Feeds and viewing the posts.

## Installation

Make sure to have the latest Go Toolchain installed, as well as Postgres database. You can then install `gator` with:

```bash
go install ...
```

## Config

Create a `.gatorconfig.json` file in your `$HOME` directory with the following structure:

```json
{
    "db_url": "postgres://username:password@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a user:

```go
gator register <username>
```

Add a feed:

```go
gator addfeed <url>
```

Start the aggregator:

```go
gator agg <time_between_requests>
```

View posts (limit defaulted to 2):

```go
gator browse [limit]
```

There are few other commands you'll need as well:

- `gator login <name>` - Login as an existing user
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow an existing feed in the database
- `gator unfollow <url>` - Unfollow an existing feed in the database