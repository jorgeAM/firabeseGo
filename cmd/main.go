package main

import (
	"github/com/jorgeAM/goFireAuth/internal/platform/http"
	"github/com/jorgeAM/goFireAuth/internal/platform/http/handler"
	"github/com/jorgeAM/goFireAuth/internal/platform/postgres"
	"github/com/jorgeAM/goFireAuth/internal/todo/application"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	_ "github.com/joho/godotenv/autoload"
)

var dbURL = os.Getenv("DB_URL")

func main() {
	opt, err := pg.ParseURL(dbURL)

	if err != nil {
		log.Fatalf("Something got wrong trying to connect with postgres: %x", err)
	}

	db := pg.Connect(opt)

	todoRepository := postgres.NewRepository(db)

	todoCreator := application.NewTodoCreator(todoRepository)
	todoReader := application.NewTodoReader(todoRepository)

	handler := handler.NewHandler(todoCreator, todoReader)

	srv := http.NewServer(handler)

	if err := srv.Run(); err != nil {
		log.Fatalf("Something got wrong trying to run http serve: %x", err)
	}
}
