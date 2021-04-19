package handler

import (
	todoApplication "github/com/jorgeAM/goFireAuth/internal/todo/application"
	userApplication "github/com/jorgeAM/goFireAuth/internal/user/application"

	"gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	todoCreator todoApplication.TodoCreator
	todoReader  todoApplication.TodoReader
	userCreator userApplication.UserCreator
	userFinder  userApplication.UserFinder
}

func NewHandler(todoCreator todoApplication.TodoCreator, todoReader todoApplication.TodoReader, userCreator userApplication.UserCreator, userFinder userApplication.UserFinder) *Handler {
	return &Handler{
		todoCreator: todoCreator,
		todoReader:  todoReader,
		userCreator: userCreator,
		userFinder:  userFinder,
	}
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (h *Handler) ValidateRequest(req interface{}) []*ErrorResponse {
	var errors []*ErrorResponse

	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
