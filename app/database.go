package app

import (
	"database/sql"
	"fmt"

	"github.com/api-abc/internal-api/configuration"
	"github.com/api-abc/internal-api/helper"

	_ "github.com/lib/pq"
)

func NewDB(di *configuration.DI) *sql.DB {
	config := di.GetConfig()
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Dbname,
		config.Database.SSLMode,
	)

	db, err := sql.Open("postgres", psqlconn)
	helper.HandlePanic(err)

	return db
}
