package http

import (
	"github/com/jorgeAM/goFireAuth/internal/platform/http/handler"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var port = os.Getenv("PORT")

type Server struct {
	engine *fiber.App
	port   string
}

func NewServer(handler *handler.Handler) *Server {
	srv := &Server{
		engine: fiber.New(),
		port:   ":" + port,
	}

	srv.engine.Use(logger.New())

	srv.registerRoutes(handler)

	return srv
}

func (s *Server) registerRoutes(handler *handler.Handler) {
	s.engine.Get("/", handler.SayHi)

	router := s.engine.Group("/todos")
	router.Post("", handler.CreateTodo)
	router.Get("", handler.GetTodos)

	s.engine.Post("/sign-up", handler.SignUp)
	s.engine.Post("/login", handler.Login)

}

func (s *Server) Run() error {
	return s.engine.Listen(s.port)
}
