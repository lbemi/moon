// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"prometheus-manager/apps/master/internal/biz"
	"prometheus-manager/apps/master/internal/conf"
	"prometheus-manager/apps/master/internal/data"
	"prometheus-manager/apps/master/internal/server"
	"prometheus-manager/apps/master/internal/service"
	"prometheus-manager/pkg/conn"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	env := bootstrap.Env
	confServer := bootstrap.Server
	trace := bootstrap.Trace
	tracerProvider := conn.NewTracerProvider(trace, env)
	confData := bootstrap.Data
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	pingRepo := data.NewPingRepo(dataData, logger)
	pingLogic := biz.NewPingLogic(pingRepo, logger)
	pingService := service.NewPingService(pingLogic, logger)
	promV1Repo := data.NewPromV1Repo(dataData, logger)
	promLogic := biz.NewPromLogic(promV1Repo, logger)
	promV1Service := service.NewPromService(promLogic, logger)
	dictRepo := data.NewDictRepo(dataData, logger)
	dictLogic := biz.NewDictLogic(dictRepo, logger)
	dictV1Service := service.NewDictService(dictLogic, logger)
	alarmPageRepo := data.NewAlarmPageRepo(dataData, logger)
	alarmPageLogic := biz.NewAlarmPageLogic(alarmPageRepo, logger)
	alarmPageV1Service := service.NewAlarmPageService(alarmPageLogic, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, pingService, promV1Service, dictV1Service, alarmPageV1Service)
	httpServer := server.NewHTTPServer(confServer, logger, tracerProvider, pingService, promV1Service, dictV1Service, alarmPageV1Service)
	pushStrategy := bootstrap.PushStrategy
	pushRepo := data.NewPushRepo(dataData, logger)
	pushLogic := biz.NewPushLogic(pushRepo, logger)
	pushService := service.NewPushService(pushLogic, logger)
	timer := server.NewTimer(pushStrategy, logger, pushService)
	app := newApp(env, logger, grpcServer, httpServer, timer)
	return app, func() {
		cleanup()
	}, nil
}
