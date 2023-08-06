package main

import (
	"github.com/ssengalanto/runic/services/account/application"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/ssengalanto/runic/services/account/docs"
)

func main() {
	a, cleanup := application.Init()
	a.Run()
	defer cleanup()
}
