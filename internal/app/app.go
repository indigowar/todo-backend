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
	"time"
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
	//log.Println(cfg)
	mongo, err := mongodb.NewClient(cfg.Mongo.URI)
	if err != nil {
		log.Fatal(err)
	}

	// close the connection to database
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := mongo.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	database := mongo.Database("todo")

	// init of all repositories
	userRepository, err := repository.NewUserMongoRepository(database)
	if err != nil {
		log.Fatal(err)
	}

	todoRepository, err := repository.NewTodoRepo(mongo)
	if err != nil {
		log.Fatal(err)
	}

	// init utilities
	tokenManager, err := auth.NewTokenManager(cfg.Auth.SigningKey, cfg.Auth.AccessTTL, cfg.Auth.RefreshTTL, cfg.Auth.Issuer)
	if err != nil {
		log.Fatal(err)
	}

	// init services

	_ = services.NewUserService(userRepository, todoRepository, tokenManager)
	_ = services.NewTodoService(userRepository, todoRepository, tokenManager)

	// init transport
	r := NewHandler(fmt.Sprint(cfg))

	// init server
	server := &http.Server{
		Handler:        r,
		Addr:           ":" + cfg.HTTP.Port,
		WriteTimeout:   cfg.HTTP.WriteTimeout,
		ReadTimeout:    cfg.HTTP.ReadTimeout,
		MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Graceful shutdown is starting...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.App.ShutdownTime)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("failed to shutdown: ", err)
	}

	log.Println("Server is stopped.")
}
