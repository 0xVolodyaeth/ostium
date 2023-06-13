package repository

import (
	"context"
	"fmt"
	"ostium/pkg/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type repository struct {
	db *bun.DB
}

type BetRepository interface {
	Create(ctx context.Context, bet *models.Bet) error
	Update(ctx context.Context, bet *models.Bet) error
}

func NewBetRepository(cfg *Config) (BetRepository, error) {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	connConfig, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse db url")
	}
	pgxConn := stdlib.OpenDB(*connConfig)
	pingCtx, pingCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer pingCancel()
	if err := pgxConn.PingContext(pingCtx); err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	dbconn := bun.NewDB(pgxConn, pgdialect.New())

	return &repository{db: dbconn}, nil
}

func (r *repository) Create(ctx context.Context, bet *models.Bet) error {
	_, err := r.db.NewInsert().Model(bet).Returning("*").Exec(ctx)
	return err
}

func (r *repository) Update(ctx context.Context, bet *models.Bet) error {
	_, err := r.db.NewUpdate().Model(bet).WherePK().Exec(ctx)
	return err
}
