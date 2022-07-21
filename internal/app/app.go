package app

import (
	"context"
	"fmt"
	"github.com/indigowar/todo-backend/internal/config"
	"github.com/indigowar/todo-backend/internal/repository"
	"github.com/indigowar/todo-backend/internal/services"
	"github.com/indigowar/todo-backend/pkg/auth"
	"github.com/indigowar/todo-backend/pkg/database/mongodb"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Handler struct {
	message string
}

func NewHandler(message string) *Handler {
	return &Handler{
		message: message,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.message))
}

func Run(cfg *config.Config) {
	log.Println(cfg)
}

func run(cfg *config.Config) {
	// Init resources
	mongoConnection, err := mongodb.NewClient("")

	if err != nil {
		log.Fatal(err)
	}

	defer func(ctx context.Context) {
		if err := mongoConnection.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}(context.Background())

	// Init repositories
	userRepo, err := repository.NewUserMongoRepository(mongoConnection)
	if err != nil {
		log.Fatal(err)
	}

	todoRepo, err := repository.NewTodoRepo(mongoConnection)
	if err != nil {
		log.Fatal(err)
	}

	// utilities
	tokenManager, err := auth.NewTokenManager(
		cfg.Auth.SigningKey,
		cfg.Auth.AccessTTL,
		cfg.Auth.RefreshTTL,
		cfg.Auth.Issuer)

	// Init services
	_ = services.NewTodoService(userRepo, todoRepo, tokenManager)
	_ = services.NewUserService(userRepo, todoRepo, tokenManager)

	// init handlers

	r := NewHandler(fmt.Sprint(cfg))

	// Init server
	server := &http.Server{
		Handler: r,
		Addr:    ":" + cfg.HTTP.Port,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Start the shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.App.ShutdownTime)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("failed to shutdown:", err)
	}

	log.Println("Server is stopped.")
}
