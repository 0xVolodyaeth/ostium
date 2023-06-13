package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"ostium/pkg/config"
	"ostium/pkg/repository"
	"ostium/pkg/service"
	"ostium/pkg/wager"
	"sync"
	"syscall"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}

	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")

	if err := v.ReadInConfig(); err != nil {
		logger.Sugar().Fatalln("Failed to connect to read in a config:", err)
	}

	cfg, err := config.New(v)
	if err != nil {
		logger.Sugar().Fatalln("Failed to connect to create a config:", err)
	}

	repo, err := repository.NewBetRepository(&cfg.DataBase)
	if err != nil {
		logger.Sugar().Fatalln("Failed to connect to db:", err)
	}

	wagerSvc, err := wager.New(&cfg.WagerConfig, logger)
	if err != nil {
		log.Fatalln(err)
	}

	svc := service.New(wagerSvc, repo, logger)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	svc.Run(ctx, &wg)

	<-sigint
	cancel()

	wg.Wait()
}
