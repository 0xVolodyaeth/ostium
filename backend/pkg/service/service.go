package service

import (
	"ostium/pkg/wager"
)

type service struct {
	wagerSvc wager.WagerService
}

type Service interface {
	Run() error
}

func New(wager wager.WagerService) Service {
	return &service{
		wagerSvc: wager,
	}
}

func (s *service) Run() error {
	return nil
}
