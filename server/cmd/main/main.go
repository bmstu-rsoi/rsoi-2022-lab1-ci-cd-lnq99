package main

import (
	"rsoi-1/internal/driver"
	"rsoi-1/internal/repository"
	"rsoi-1/internal/server"
	"rsoi-1/internal/service"

	"rsoi-1/config"
)

func main() {
	var err error

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := driver.NewSQLDatabase(&cfg.DbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	repo := repository.NewSqlRepository(db)

	services := service.InitServices(repo)

	s := server.NewEchoServer(services, &cfg.ServerConfig)
	if err = s.Run(); err != nil {
		panic(err)
	}
}
