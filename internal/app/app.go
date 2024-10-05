package app

import (
	postgres "blog/internal/databases"
	"blog/internal/repositories"
	"blog/internal/services"
	http "blog/internal/transport"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	gin *gin.Engine
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}
}

func New() (*App, error) {
	a := &App{}
	gin.SetMode(os.Getenv("GIN_MODE"))
	a.gin = gin.Default()

	a.gin.Use(
		gin.Recovery(),
	)
	postgresDBConnection := postgres.NewPostgresDBConnection()

	repos := repositories.NewUserRepositories(postgresDBConnection)
	services := services.NewAuthorizationServices(repos)

	handlers := http.NewHandler(services)

	userGroup := a.gin.Group("/api")
	handlers.RegisterUserAPI(userGroup)
	return a, nil
}

func (a *App) Run() error {
	log.Println("Blog service launched successfully")
	err := a.gin.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return err
	}
	return nil
}
