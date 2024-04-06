package data

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/pritesh-mantri/sailor/config"
)

type Models struct {
	Users interface {
		Create(ctx context.Context, user User) error
		GetByID(ctx context.Context, id int64) (*User, error)
		GetAuthInfo(ctx context.Context, id int64) (*AuthInfo, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context) *User
	}
}

func New(cfg config.Config) Models {
	db := connectDB(cfg)

	return Models{
		Users: UserModel{db},
	}
}

func connectDB(cfg config.Config) *sqlx.DB {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", cfg.DBConfig.User, cfg.DBConfig.DBname, cfg.DBConfig.Password)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %e", err)
	}

	log.Println("connected to database")
	return db
}
