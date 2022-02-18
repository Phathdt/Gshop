## Requirements
1. [X] Implement a Rest API with CRUD functionality.
2. [X] Database: PostgreSQL.
3. [X] Set up service with docker compose.
4. [X] Authen with jwt, whitelist token
5. [X] generate query with gorm
6. [X] Private and public routes
7. [X] Split logic into transport, handler, repo and storage
8. [X] API Document
9. [X] Redis cache
10. [X] Unit test
11. [X] Hot Reload with Air

## Technology Stack

- **Go 1.17**: *Leverage the standard libraries as much as possible*
- **PostgreSQL**: *RDBMS of choice because of faster read due to its indexing model and safer transaction with better isolation levels handling*
- **Fiber**: *Fast and have respect for native net/http API*
- **Gorm**: *The fantastic ORM library for Golang*
- **JWT Token**: *Also implemented to demonstrate the decoupility*
- **Goose**: *Efficient schema generating, up/down migrating*
- **Docker** + **Docker-Compose**: *Containerization, what else to say ...*
- **Env**: *Add robustness to configurations*
- **Github Actions CI**: *Make sure we don't push trash code into the codebase*
- **Redis**: *caching database*
- **Air**: *hot reload*

## Booting Up

**Docker**
```bash
cp .env.sample .env

docker-compose build

docker-compose up

# docker-compose down
```

**Local**
```bash
# insert some data in env like database url, debug level

# migrate
make migrate args="up"

# run app
make server

# or hot reload running with air
# install air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
make air
```

## Migration

**Generate migration**
```bash
make migrate args="create create_xyz sql"
```

## Testing

Lib: https://github.com/smartystreets/goconvey

Test case format
```go
// Only pass t into top-level Convey calls
Convey("Given ...", t, func() {
  ...

  Convey("When ...", func() {
    ...

    Convey("Then ...", func() {
      ...
    })
  })
})
```


## Todos
- add manifest K8s
- deploy with Github Action and ArgoCD
- add replica postgres
- add more testable for repository, usecase, ...
