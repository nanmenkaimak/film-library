package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nanmenkaimak/film_library/internal/config"
)

type Db struct {
	Db *sqlx.DB
}

type Config config.DbNode

func (c Config) dsn() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Name, c.User, c.Password, c.SslMode)
}

func New(cfg config.DbNode) (*Db, error) {
	conf := Config(cfg)
	conn, err := sqlx.Connect("postgres", conf.dsn())
	if err != nil {
		return nil, fmt.Errorf("db connection failed: %w", err)
	}

	return &Db{
		Db: conn,
	}, nil
}

//func (d *Db) Close() error {
//	sqlDB, _ := d.Db.DB()
//
//	return sqlDB.Close()
//}
