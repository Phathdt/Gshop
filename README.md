## Requirements
1. [X] Implement a Rest API with CRUD functionality.
2. [X] Database: PostgreSQL.
3. [X] Set up service with docker compse.
4. [X] Authen with jwt
5. [X] generate query with gorm
6. [X] Private and public routes
7. [ ] API Document
8. [ ] Redis cache


## Technology Stack

- **Go 1.17**: *Leverage the standard libraries as much as possible*
- **PostgreSQL**: *RDBMS of choice because of faster read due to its indexing model and safer transaction with better isolation levels handling*
- **Fiber**: *Fast and have respect for native net/http API*
- **Gorm**: *The fantastic ORM library for Golang*
- **JWT Token**: *Also implemented to demonstrate the decoupility*
- **Golang-Migrate**: *Efficient schema generating, up/down migrating*
- **Docker** + **Docker-Compose**: *Containerization, what else to say ...*
- **Viper**: *Add robustness to configurations*
- **Github Actions CI**: *Make sure we don't push trash code into the codebase*

## Booting Up

```bash
docker-compose build

docker-compose up

# docker-compose down
```

## Todos
- add manifest K8s
- deploy with Github Action and ArgoCD
- add replica postgres
- add cache DB ( Redis )
- add more testable for repository, usecase, ...
- add pagination and meta data
