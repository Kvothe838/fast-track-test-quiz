package main

import (
	"context"
	"flag"
	"github.com/Kvothe838/fast-track-test-quiz/internal/app/controller"
	"github.com/Kvothe838/fast-track-test-quiz/internal/app/server"
	"github.com/Kvothe838/fast-track-test-quiz/internal/config"
	"github.com/Kvothe838/fast-track-test-quiz/internal/database/memory"
	"github.com/Kvothe838/fast-track-test-quiz/internal/pkg/graceful"
	"github.com/Kvothe838/fast-track-test-quiz/internal/pkg/logger"
	"github.com/Kvothe838/fast-track-test-quiz/internal/services"
)

func main() {
	ctx := context.Background()

	filePath := flag.String("config", "", "path of configuration file")
	flag.Parse()

	var sources []config.Source
	conf := config.New(ctx, *filePath, sources...)

	repo := memory.NewRepository()
	interactor := services.NewInteractor(repo)

	setupRestAPI(ctx, interactor, conf.Port)

	if err := graceful.Wait(); err != nil {
		logger.CtxWarn(ctx, err)
	}
}

func setupRestAPI(ctx context.Context, interactor controller.Interactor, port string) {
	ctrl := controller.New(interactor)
	srv := server.New(port)
	srv.RegisterHandler(ctrl)
	srv.StartAsync()

	logger.CtxInfof(ctx, "listening on :%s", port)
}
