package service

import (
	"context"
	"log"
	"ostium/pkg/models"
	"ostium/pkg/repository"
	"ostium/pkg/wager"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
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

				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()

				if err := s.handleBetMade(ctx, betMade); err != nil {
					s.logger.Error("failed to handle bet made:", zap.Error(err))
				}

			case betJoin := <-betJoinChan:
				s.logger.Info("betJoin", zap.Any("betJoin", betJoin))
			}
		}
	}()
}

func (s *service) handleBetMade(ctx context.Context, betMade *wager.WagerBetMade) error {
	bet := &models.Bet{
		ID:           betMade.BetId.Int64(),
		Amount:       betMade.Amount.Int64(),
		Expiration:   betMade.Expiration.Int64(),
		CreatedAt:    betMade.CreatedAt.Int64(),
		OpeningPrice: betMade.OpeningPrice.Int64(),
		IsActive:     false,
		Withdrawn:    false,
		Winner:       common.Address{}.String(),
	}

	s.logger.Info("A new bet has been made",
		zap.Int64("betId", betMade.BetId.Int64()),
		zap.Bool("long", betMade.Long),
		zap.Int64("amount", betMade.Amount.Int64()),
		zap.Int64("expiration", betMade.Expiration.Int64()),
		zap.Int64("createdAt", betMade.CreatedAt.Int64()),
		zap.Int64("openingPrice", betMade.OpeningPrice.Int64()),
	)

	if betMade.Long {
		bet.LongAddress = betMade.Initiator.String()
	} else {
		bet.ShortAddress = betMade.Initiator.String()
	}

	err := s.repo.Create(ctx, bet)

	return err
}

func (s *service) handleBetJoined(ctx context.Context, betMade *wager.WagerJoinBet) error {
	panic("implement me")
}

func (s *service) handleBetCanceled(ctx context.Context, betMade *wager.WagerBetCanceled) error {
	panic("implement me")
}

func (s *service) handleBetWithdrawn(ctx context.Context, betMade *wager.WagerWithdrawn) error {
	panic("implement me")
}
