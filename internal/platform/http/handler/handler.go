package handler

import "github/com/jorgeAM/goFireAuth/internal/todo/application"

type Handler struct {
	todoCreator application.TodoCreator
	todoReader  application.TodoReader
}

func NewHandler(todoCreator application.TodoCreator, todoReader application.TodoReader) *Handler {
	return &Handler{
		todoCreator: todoCreator,
		todoReader:  todoReader,
	}
}
