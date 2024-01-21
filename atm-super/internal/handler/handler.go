package handler

import "log"

type Handler struct {
	infoLogger *log.Logger
	errLogger  *log.Logger
}

func NewHandler(infoLog, errLog *log.Logger) *Handler {
	return &Handler{
		infoLogger: infoLog,
		errLogger:  errLog,
	}
}
