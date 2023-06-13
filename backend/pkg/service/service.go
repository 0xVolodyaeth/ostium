package service

import (
	"context"
	"log"
	"ostium/pkg/repository"
	"ostium/pkg/wager"
	"sync"

	"go.uber.org/zap"
)

type service struct {
	wagerSvc wager.WagerService
	repo     repository.BetRepository
	logger   *zap.Logger
}

type Service interface {
	Run(ctx context.Context, wg *sync.WaitGroup)
}

func New(wager wager.WagerService, repo repository.BetRepository, logger *zap.Logger) Service {
	return &service{
		wagerSvc: wager,
		repo:     repo,
		logger:   logger,
	}
}

func (s *service) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.logger.Info("starting the service...")

		canceledChan, withdrawnChan, betMadeChan, betJoinChan, err := s.wagerSvc.Run()
		if err != nil {
			log.Fatalln(err)
		}

		for {
			select {
			case <-ctx.Done():
				s.logger.Info("shutting the service...")
				s.wagerSvc.Close()
				return

			case canceled := <-canceledChan:
				s.logger.Info("canceled", zap.Any("canceled", canceled))
			case withdrawn := <-withdrawnChan:
				s.logger.Info("withdrawn", zap.Any("withdrawn", withdrawn))
			case betMade := <-betMadeChan:
				log.Println("betMade", betMade)
			case betJoin := <-betJoinChan:
				s.logger.Info("betJoin", zap.Any("betJoin", betJoin))
			}
		}
	}()
}
