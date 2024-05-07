# GOGO Dancers Project for course 4IT428 - Vývoj mikroslužeb v jazyce Go

The goal of the project is still unknown to me, I read the assigment long time ago. But this is the base for it.

## Assignment

You can see the assignment [here](assignment.pdf).

## Run the code

```bash
  go run main.go
```

## Stack

routing - [chi](https://github.com/go-chi/chi)

## Contributing

Please follow rules in [CONTRIBUTING.md](./CONTRIBUTING.md)

## Let's GO!

![MeAndGo](logo.jpg)

## Links

[Lecture repo](https://github.com/strvcom/backend-go-vse-microservice-development)

## Create db locally

I'm doing something soooo wrong... it can't be that complicated

- install postgres

```bash
# mac
brew install postgresql@16

# run postgres
brew services start postgresql@16
```

- create postgres user

```bash
# install psql
brew install psql
# open postgres
psql postgres
# in psql create user
CREATE USER postgres_user WITH PASSWORD 'user';
# and create database
CREATE DATABASE godb WITH OWNER postgres_user ENCODING='UTF8';
```

- run migrations

```bash
# install tool for go migrations
brew install golang-migrate
export POSTGRESQL_URL="postgres://postgres_user:user@localhost:5432/godb?sslmode=disable"
migrate -database {$POSTGRESQL_URL} -path db/migrations up
```


export POSTGRESQL_URL="host=aws-0-eu-central-1.pooler.supabase.com port=5432 user=postgres.itzckmthyazjhgqkptum password=jIj22fbwM1OqRBAw dbname=postgres sslmode=disable"