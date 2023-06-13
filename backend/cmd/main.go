package main

import (
	"log"
	"os"
	"os/signal"
	"ostium/pkg/wager"
	"sync"
	"syscall"

	"go.uber.org/zap"
)

func main() {

	cfg := wager.Config{
		ProviderURI: "wss://polygon-mumbai.g.alchemy.com/v2/A20qXmbk28bOB0C5P852lsiYhYkFQ_zx",
		Address:     "0xfA217E7f00FFB089Faf570BD2Beed3a0193D9bCd",
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}

	svc, err := wager.New(&cfg, logger)
	if err != nil {
		log.Fatalln(err)
	}

	sigint := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		logger.Info("starting the service...")

		canceledChan, withdrawnChan, betMadeChan, betJoinChan, err := svc.Run()
		if err != nil {
			log.Fatalln(err)
		}

		for {
			select {
			case <-done:
				logger.Info("shutting the service...")
				svc.Close()
				return

			case canceled := <-canceledChan:
				logger.Info("canceled", zap.Any("canceled", canceled))
			case withdrawn := <-withdrawnChan:
				logger.Info("withdrawn", zap.Any("withdrawn", withdrawn))
			case betMade := <-betMadeChan:
				log.Println("betMade", betMade)
			case betJoin := <-betJoinChan:
				logger.Info("betJoin", zap.Any("betJoin", betJoin))
			}
		}
	}()

	<-sigint
	done <- struct{}{}

	wg.Wait()
}
