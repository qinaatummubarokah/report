package repository

import (
	"context"
)

type Data struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Repository interface {
	GetReport(ctx context.Context) ([]Data, error)
}
