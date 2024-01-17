package main

import (
	"atm/config"
	"atm/internal/handler"
	"atm/internal/repository"
	"atm/internal/server"
	"atm/internal/service"
	"atm/logger"
	"fmt"
	"log"
	"os"
)

func main() {
	infoLog, errLog, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("NewLogger: %s", err)
	}

	var cfgPath string

	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	default:
		cfgPath = "./.env"
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		errLog.Fatalf("LoadConfig %s", err)
	}

	db, err := repository.NewDB(cfg)
	if err != nil {
		fmt.Println(err)
		errLog.Fatalf("NewDB %s", err)
	}

	r := repository.NewRepository(db)

	s := service.NewService(*r)

	h := handler.NewHandler(s, infoLog, errLog)

	errLog.Fatal(server.RunServer(cfg, h.Routes(), infoLog))
}
