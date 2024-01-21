package handler

import (
	"atm/internal/service"
	"log"
)

type Handler struct {
	service    *service.Service
	infoLogger *log.Logger
	errLogger  *log.Logger
}

func NewHandler(service *service.Service, infoLog, errLog *log.Logger) *Handler {
	return &Handler{
		service:    service,
		infoLogger: infoLog,
		errLogger:  errLog,
	}
}
