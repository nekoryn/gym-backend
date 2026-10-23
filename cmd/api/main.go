package main

import (
	"log"

	"gym-backend/internal/config"
	"gym-backend/internal/database"
	httpDelivery "gym-backend/internal/delivery/http"
	"gym-backend/internal/repository"
	"gym-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"

	_ "gym-backend/docs"
)

// @title Gym Tracker API
// @version 1.0
// @description Backend API для трекинга тренировок и оборудования в зале
// @host localhost:8080
// @BasePath /
func main() {
	zapLogger, _ := zap.NewDevelopment()
	defer zapLogger.Sync()

	cfg := config.LoadConfig()

	zapLogger.Info("Запуск сервера...")

	db := database.InitDB(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Инициализация слоев Clean Architecture
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := httpDelivery.NewAuthHandler(authService)

	app := fiber.New(fiber.Config{
		AppName: "Gym Tracker v1.0",
	})

	app.Use(logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "pong", "message": "Сервер работает отлично!"})
	})

	log.Fatal(app.Listen(":8080"))
}