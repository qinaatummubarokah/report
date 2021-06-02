package service

import (
	"context"
	"fmt"
	"log"

	"report/repository"
)

type service struct {
	logger log.Logger
	db     repository.Repository
}

// Service interface
type Service interface {
	GetReport(ctx context.Context) (interface{}, error)
}

// NewService func initializes a service
func NewService(logger log.Logger, repo repository.Repository) Service {
	return &service{
		logger: logger,
		db:     repo,
	}
}

func (s *service) GetReport(ctx context.Context) (interface{}, error) {
	transactionEcv, err := s.db.GetReport(ctx)
	if err != nil {
		return "", err
	}

	fmt.Println("transaction mybb : ", transactionEcv)
	return "success", nil
}
