package helpers

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/obidovsamandar/go-task-auth/config"
)

var DBClient *sqlx.DB

func ConnectionDB() {
	cfg := config.Load()

	var dbConnectionString = "host=" + cfg.PG_HOST + " port=" + cfg.PG_PORT + " user=" + cfg.PG_USER + " password=" + cfg.PG_PASSWORD + " dbname=" + cfg.PG_DB + " sslmode=" + cfg.SSlMode
	db, err := sqlx.Open("postgres", dbConnectionString)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}

	DBClient = db

}
