package wager

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type wagerService struct {
	wager           *Wager
	client          *ethclient.Client
	contractAddress common.Address
	log             *zap.Logger
	done            chan struct{}
}

type WagerService interface {
	Run() (
		chan *WagerBetCanceled,
		chan *WagerWithdrawn,
		chan *WagerBetMade,
		chan *WagerJoinBet,
		error,
	)
	Close()
}

func New(cfg *Config, log *zap.Logger) (WagerService, error) {
	client, err := ethclient.Dial(cfg.ProviderURI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial client")
	}

	wagerInst, err := NewWager(common.HexToAddress(cfg.Address), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create stable instance")
	}

	contractAddress := common.HexToAddress(cfg.Address)

	return &wagerService{
		wager:           wagerInst,
		client:          client,
		log:             log,
		contractAddress: contractAddress,
		done:            make(chan struct{}),
	}, nil
}

func (s *wagerService) Run() (
	chan *WagerBetCanceled,
	chan *WagerWithdrawn,
	chan *WagerBetMade,
	chan *WagerJoinBet,
	error,
) {
	canceledChan := make(chan *WagerBetCanceled)
	withdrawnChan := make(chan *WagerWithdrawn)
	betMadeChan := make(chan *WagerBetMade)
	betJoinChan := make(chan *WagerJoinBet)

	betCanceledSub, err := s.wager.WatchBetCanceled(nil, canceledChan, nil)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed to subscribe to BetCanceled event")
	}

	withdrawnSub, err := s.wager.WatchWithdrawn(nil, withdrawnChan, nil)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed to subscribe to Withdrawn event")
	}

	betMadeSub, err := s.wager.WatchBetMade(nil, betMadeChan, nil)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed to subscribe to BetMade event")
	}

	betJoinedSub, err := s.wager.WatchJoinBet(nil, betJoinChan, nil, nil)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed to subscribe to BetJoined event")
	}

	go func() {
		for {
			select {
			case <-s.done:
				betCanceledSub.Unsubscribe()
				betCanceledSub.Unsubscribe()
				betMadeSub.Unsubscribe()
				betJoinedSub.Unsubscribe()

				return

			case err := <-betJoinedSub.Err():
				s.log.Error(err.Error())

			case err := <-betMadeSub.Err():
				s.log.Error(err.Error())

			case err := <-withdrawnSub.Err():
				s.log.Error(err.Error())

			case err := <-betCanceledSub.Err():
				s.log.Error(err.Error())
			}
		}
	}()

	return canceledChan, withdrawnChan, betMadeChan, betJoinChan, err
}

func (s *wagerService) Close() {
	s.log.Info("closing wager service...")
	s.done <- struct{}{}
}
