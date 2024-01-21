package main

import (
	"atm/config"
	"atm/internal/handler"
	"atm/internal/server"
	"atm/internal/service"
	"atm/logger"
	"log"
	"os"
)

var (
	AuthURL string = "http://localhost:8181/auth"
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

	s := service.NewService(AuthURL)

	h := handler.NewHandler(s, infoLog, errLog)

	errLog.Fatal(server.RunServer(cfg, h.Routes(), infoLog))
}
