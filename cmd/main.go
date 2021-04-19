package main

import (
	"github/com/jorgeAM/goFireAuth/internal/platform/http"
	"github/com/jorgeAM/goFireAuth/internal/platform/http/handler"
	"github/com/jorgeAM/goFireAuth/internal/platform/postgres"
	todoApplication "github/com/jorgeAM/goFireAuth/internal/todo/application"
	userApplication "github/com/jorgeAM/goFireAuth/internal/user/application"
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

	todoRepository := postgres.NewTodoRepository(db)
	userRepository := postgres.NewUserRepository(db)

	todoCreator := todoApplication.NewTodoCreator(todoRepository)
	todoReader := todoApplication.NewTodoReader(todoRepository)
	userCreator := userApplication.NewUserCreator(userRepository)
	userFinder := userApplication.NewUserFinder(userRepository)

	handler := handler.NewHandler(todoCreator, todoReader, userCreator, userFinder)

	srv := http.NewServer(handler)

	if err := srv.Run(); err != nil {
		log.Fatalf("Something got wrong trying to run http serve: %x", err)
	}
}
