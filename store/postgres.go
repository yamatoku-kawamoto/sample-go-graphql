package store

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	Product_Postgres = "postgres"
)

func OpenPostgres(conf *Configure) (*sql.DB, error) {
	sslmode := func(sslmode bool) string {
		if sslmode {
			return "sslmode=enable"
		}
		return "sslmode=disable"
	}
	f := func(key, defaultvalue string, value ...string) string {
		if len(value) == 0 {
			return fmt.Sprintf("%s=%s", key, defaultvalue)
		}
		return fmt.Sprintf("%s=%s", key, value[0])
	}
	formatdsn := func() (string, error) {
		values := []string{
			f("host", "localhost", conf.Host),
			f("dbname", "postgres", conf.Database),
			f("user", "postgres", conf.User),
			f("password", conf.Password),
			sslmode(conf.SSLMode),
		}
		return strings.Join(values, " "), nil
	}
	dsn, err := formatdsn()
	if err != nil {
		return nil, err
	}
	return sql.Open(Product_Postgres, dsn)
}
