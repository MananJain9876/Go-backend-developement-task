package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/example/user-age-api/internal/logger"
	"github.com/example/user-age-api/internal/middleware"
	"github.com/example/user-age-api/internal/routes"
	"github.com/example/user-age-api/internal/service"
	"github.com/example/user-age-api/internal/repository"
)

func main() {
	// Basic environment-based configuration
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to create pgx pool: %v", err)
	}
	defer pool.Close()

	zapLogger := logger.NewLogger()
	defer zapLogger.Sync()

	app := fiber.New()

	// Middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(zapLogger))

	// Repository and services
	userRepo := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepo)

	// Routes
	routes.RegisterUserRoutes(app, userService, zapLogger)

	// Change port here if 8080 is occupied
	if err := app.Listen(":3000"); err != nil {
		zapLogger.Fatal("failed to start server", zap.Error(err))
	}
}


