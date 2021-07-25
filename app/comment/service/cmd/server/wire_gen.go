// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"base-service/app/comment/service/internal/biz"
	"base-service/app/comment/service/internal/conf"
	"base-service/app/comment/service/internal/data"
	"base-service/app/comment/service/internal/server"
	"base-service/app/comment/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	commentRepo := data.NewCommentRepo(dataData, logger)
	commentUsecase := biz.NewCommentUsecase(commentRepo, logger)
	commentService := service.NewCommentService(commentUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, commentService, logger)
	grpcServer := server.NewGRPCServer(confServer, commentService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
