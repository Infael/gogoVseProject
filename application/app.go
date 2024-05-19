package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
)

type App struct {
	router       http.Handler
	db           *db.Database
	repositories *repository.Repositories
	services     *service.Services
}

func New() *App {
	app := &App{}

	if os.Getenv("PROD") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			panic(err)
		}
	}

	// init db
	database, err := db.NewDatabase(&db.InitDatabase{
		Host:          os.Getenv("POSTGRES_HOST"),
		Port:          os.Getenv("POSTGRES_PORT"),
		Username:      os.Getenv("POSTGRES_USER"),
		Password:      os.Getenv("POSTGRES_PASSWORD"),
		Dbname:        os.Getenv("POSTGRES_DB"),
		RunMigrations: os.Getenv("POSTGRES_RUN_MIGRATIONS"),
	})

	if err != nil {
		log.Fatal("failed to connect to database: %v", err)
		panic(err)
	}
	app.db = &database

	// init cache
	cache := cache.New(15*time.Minute, 20*time.Minute)

	// init mail dailer
	// provider := os.Getenv("STMP_PROVIDER")
	// port, err := strconv.Atoi(os.Getenv("STMP_PORT"))
	// if err != nil {
	// 	log.Fatal("failed to stmp server: %v", err)
	// 	panic(err)
	// }
	// user := os.Getenv("STMP_MAIL")
	// pwd := os.Getenv("STMP_PWD")
	// mailDialer := gomail.NewDialer(provider, port, user, pwd)

	app.repositories = repository.NewRepositories(app.db)
	app.services = service.NewServices(app.repositories, cache, nil) // TODO MS: uncomment code above and add mailDialer here
	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	defer a.db.Connection.Close()

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
