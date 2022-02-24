package main

import (
	"context"
	"fmt"
	"gshop/module/users/userhandler"
	"gshop/module/users/usermodel"
	"gshop/module/users/userrepo"
	"gshop/module/users/userstorage"
	"gshop/svcctx"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sc, err := svcctx.NewServiceContext(ctx)
	if err != nil {
		panic(err)
	}

	storage := userstorage.NewUserSQLStorage(sc.DB)
	repo := userrepo.NewUserRepo(storage)

	rdb := userstorage.NewTokenStore(sc.RdClient)
	tokenRepo := userrepo.NewTokenRepo(rdb)

	hdl := userhandler.NewLoginUserHdl(repo, tokenRepo)
	var input usermodel.UserLogin

	token, err := hdl.Response(ctx, &input)
	if err != nil {
		panic(err)
	}

	fmt.Println(token)
}
