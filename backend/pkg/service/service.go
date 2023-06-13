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
	"github.com/pkg/errors"
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
				if err := s.handleBetCanceled(ctx, canceled); err != nil {
					s.logger.Error("failed to handle bet joined:", zap.Error(err))
				}

			case withdrawn := <-withdrawnChan:
				if err := s.handleBetWithdrawn(ctx, withdrawn); err != nil {
					s.logger.Error("failed to handle bet made:", zap.Error(err))
				}

			case betMade := <-betMadeChan:
				if err := s.handleBetMade(ctx, betMade); err != nil {
					s.logger.Error("failed to handle bet made:", zap.Error(err))
				}

			case betJoined := <-betJoinChan:
				if err := s.handleBetJoined(ctx, betJoined); err != nil {
					s.logger.Error("failed to handle bet joined:", zap.Error(err))
				}

			}
		}
	}()
}

func (s *service) handleBetMade(ctx context.Context, betMade *wager.WagerBetMade) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

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
		zap.String("creator", betMade.Initiator.String()),
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

func (s *service) handleBetJoined(ctx context.Context, joinBet *wager.WagerJoinBet) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	bet := &models.Bet{
		ID: joinBet.BetId.Int64(),
	}

	if err := s.repo.GetById(ctx, bet); err != nil {
		return errors.Wrap(err, "failed to get a bet by id")
	}

	s.logger.Info("A bet has been joined",
		zap.Int64("betId", joinBet.BetId.Int64()),
		zap.String("joiner", joinBet.Joiner.String()),
	)

	if bet.LongAddress == "" {
		bet.LongAddress = joinBet.Joiner.String()
	}

	if bet.ShortAddress == "" {
		bet.ShortAddress = joinBet.Joiner.String()
	}

	bet.IsActive = true

	err := s.repo.Update(ctx, bet)
	return err
}

func (s *service) handleBetCanceled(ctx context.Context, betCanceled *wager.WagerBetCanceled) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	bet := &models.Bet{
		ID: betCanceled.BetId.Int64(),
	}

	if err := s.repo.GetById(ctx, bet); err != nil {
		return errors.Wrap(err, "failed to get a bet by id")
	}

	s.logger.Info("A bet has been canceled",
		zap.Int64("betId", betCanceled.BetId.Int64()),
	)

	bet.Canceled = true

	return s.repo.Update(ctx, bet)
}

func (s *service) handleBetWithdrawn(ctx context.Context, betWithdrawn *wager.WagerWithdrawn) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	bet := &models.Bet{
		ID: betWithdrawn.BetId.Int64(),
	}

	if err := s.repo.GetById(ctx, bet); err != nil {
		return errors.Wrap(err, "failed to get a bet by id")
	}

	s.logger.Info("A bet has been withdrawn",
		zap.Int64("betId", betWithdrawn.BetId.Int64()),
		zap.String("winner", betWithdrawn.Winner.String()),
	)

	bet.IsActive = false
	bet.Winner = betWithdrawn.Winner.String()
	bet.Withdrawn = true

	return s.repo.Update(ctx, bet)
}
