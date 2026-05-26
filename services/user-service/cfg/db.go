package cfg

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewDbConn() *bun.DB {
	godotenv.Load(".env")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("UDBNAME"))
	sqldb := sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithDSN(dsn),
		),
	)
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}
