package main

import (
	"context"
	"errors"
	"fmt"
	"horsy_api/config"
	"horsy_api/internal/handler"
	"horsy_api/internal/repository"
	"horsy_api/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	defer zapSync()

	cfg, err := config.LoadConfig()
	if err != nil {
		zap.L().Fatal("failed to load configs", zap.Error(err))
	}

	db, conn, err := repository.NewPostgresDB(&cfg.Database)
	if err != nil {
		zap.L().Fatal("failed to connect to PostgresDB", zap.Error(err))
	}

	r := repository.NewTransactionalRepository(db)
	s := service.NewService(cfg, r)
	sagaRunner := service.NewSagaRunner(s, r)
	h := handler.NewHandler(sagaRunner)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	router := h.GetRouter(&cfg.Server)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: router,
	}

	go func() {
		zap.L().Info("starting HTTP server", zap.String("port", cfg.Server.Port))
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("failed to start server", zap.Error(err))
		}
	}()

	<-stop
	zap.L().Info("shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn.Close()

	if err = srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("server forced to shutdown", zap.Error(err))
	}

	zap.L().Info("server exited properly")
}

func zapSync() {
	if err := zap.L().Sync(); err != nil {
		zap.L().Error("failed to sync logger", zap.Error(err))
	}
}
