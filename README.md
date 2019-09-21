# Highscore server

Server to store highscore data for my game Go/Android game [Splendid](https://github.com/kyeett/splendid.git).

Uses:

- Go
- Heroku

## Usage

Run with:

```sh
PORT=8080 make run
```

## Misc

### Install `migrate`

```sh
brew install golang-migrate
```

### Setup heroku

```sh
heroku apps:create highscore-backend --region eu
heroku git:remote -a highscores-backend
```

### Links

- [`migrate` Postgres tutorial](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)
- [Heroku and `migrate`](https://devcenter.heroku.com/articles/go-support)
