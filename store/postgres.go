package store

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	Product_Postgres = "postgres"
)

func OpenPostgres(conf *Configure) *sql.DB {
	dsn := func() string {
		host := func() string {
			if conf.Host == "" {
				return "localhost"
			}
			return conf.Host
		}
		port := func() int {
			if conf.Port <= 0 || conf.Port >= 65535 {
				return 5432
			}
			return conf.Port
		}
		user := func() string {
			if conf.User == "" {
				return "postgres"
			}
			return conf.User
		}
		dbname := func() string {
			if conf.Database == "" {
				return "postgres"
			}
			return conf.Database
		}
		// "postgres://postgres:@localhost:5432/test?sslmode=disable"
		return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", Product_Postgres, user(), conf.Password, host(), port(), dbname())
	}
	return sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn())))
}
