package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ssengalanto/runic/pkg/http/mux"
	"github.com/ssengalanto/runic/pkg/pgsql"
	"github.com/ssengalanto/runic/services/account/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	pg := cfg.PGSQL
	db, err := pgsql.NewConnection(
		pg.Username,
		pg.Password,
		pg.Host,
		strconv.Itoa(pg.Port),
		pg.DBName,
		pg.QueryParams,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.New()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Runic Account Service!")) //nolint:errcheck //unnecessary
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTP.Port), r)) //nolint:gosec //todo
}
