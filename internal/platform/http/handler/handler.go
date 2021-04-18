package handler

import (
	todoApplication "github/com/jorgeAM/goFireAuth/internal/todo/application"
	userApplication "github/com/jorgeAM/goFireAuth/internal/user/application"
)

type Handler struct {
	todoCreator todoApplication.TodoCreator
	todoReader  todoApplication.TodoReader
	userCreator userApplication.UserCreator
}

func NewHandler(todoCreator todoApplication.TodoCreator, todoReader todoApplication.TodoReader, userCreator userApplication.UserCreator) *Handler {
	return &Handler{
		todoCreator: todoCreator,
		todoReader:  todoReader,
		userCreator: userCreator,
	}
}
