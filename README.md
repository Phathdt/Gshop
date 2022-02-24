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

**Project Structure**

```bash
.
├── README.md
├── backend
│   ├── Makefile
│   ├── cmd
│   │   ├── command # chay trong terminal
│   │   │   └── main.go
│   │   ├── migrate # migrate db
│   │   │   ├── Dockerfile
│   │   │   ├── main.go
│   │   │   └── run.sh
│   │   └── server # server api
│   │       ├── Dockerfile
│   │       ├── internal # rieng cua app sv api thoi
│   │       │   ├── middleware
│   │       │   └── server
│   │       ├── main.go
│   │       └── run.sh
│   ├── docs.yaml # file open api
│   ├── go.mod # deps
│   ├── go.sum # deps lock
│   ├── integration # client tuong tac voi ben ngoai, vd lazada, shopee, ...
│   ├── main_test.go
│   ├── migrations # folder migrations
│   ├── module
│   │   └── users # module user
│   │       ├── userhandler # handler
│   │       │   ├── create_user.go
│   │       │   ├── get_user.go
│   │       │   ├── login_user.go
│   │       │   └── logout_user.go
│   │       ├── usermodel #model
│   │       │   ├── user.go
│   │       │   ├── user_create.go
│   │       │   └── user_login.go
│   │       ├── userrepo #repo
│   │       │   ├── repo.go
│   │       │   └── token_repo.go
│   │       ├── userstorage # tuong tac voi db
│   │       │   ├── storage.go
│   │       │   └── token.go
│   │       └── usertransport #  tầng transport
│   │           └── userfiber # fiber transport, nếu có socket thì user socket, usergprc
│   ├── pkg # chứa những package xài chung, common
│   │   ├── common
│   │   ├── config
│   │   ├── httpserver
│   │   ├── logger
│   │   └── sdkcm
│   ├── service # nếu logic phức tạp thì để trong này, vd call api order của lazada -> order service, trong đó sẽ gọi thằng integration/lazada chẳng hạn
│   ├── svcctx
├── docker-compose.yml
└── nginx
    ├── Dockerfile
    └── nginx.conf
```

**Viết 1 api mới** 
- đầu tiên suy nghĩ về đó là domain/module mới hay gì, nếu là mới thì tạo trong module 
- suy nghĩ có những field gì cần chứa -> tạo migration
- xong viết tầng model của module đó, vd usermodel 
- từ đây có 2 đường là bottom down hoặc bottom up

*bottom down*
- viết tầng handler, có interface gọi đến tầng model 
- viết tầng model + interface cho nó, implement interface vừa viết ở transport 
- viết tầng storage, implement interface ở tầng repo, tương tác với db qua GORM hoặc tự viết raw sql 
- tạo endpoint ở server.go, xong trong transport -> xyzfiber tạo 1 file chứa code transport, ở đây call đến thằng handler, cung cấp đủ arg cho nó 
- nếu có cache thì thêm con redis repo với redis storage vào, default là sql ( postgres ) 

*bottom up* 
- ngược lại, viết từ storage lên repo, handler và transport 

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
