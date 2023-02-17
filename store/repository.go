package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/iwashi623/go_posts/clock"
	"github.com/iwashi623/go_posts/config"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword,
			cfg.DBHost, cfg.DBPort, cfg.DBName,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	//Openでは接続テストが実行されていないのでPingで接続テストを行う
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, err
	}
	xdb := sqlx.NewDb(db, "mysql")
	return xdb, func() { _ = db.Close() }, nil
}

type Beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Preparer interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type Queryer interface {
	Preparer
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

var (
	_ Beginner = (*sqlx.DB)(nil)
	_ Preparer = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.Tx)(nil)
	_ Queryer  = (*sqlx.DB)(nil)
)

type Repository struct {
	Clocker clock.Clocker
}
