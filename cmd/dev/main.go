package main

import (
	"log/slog"
	handler "smart-waste/apis/user/handlers"
	route "smart-waste/apis/user/routes"
	"smart-waste/domain/user/entity"

	"smart-waste/domain/user/usecase"
	"smart-waste/pkgs/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var enableMigration = true

func main() {
	slog.Info("service running on port 3000")
	// 1. init fiber instance
	app := fiber.New()

	// 2. connect to database
	var db = connectDB()

	// 3. init route
	var userHander = handler.UserHandler{
		CreateUserUsecase: usecase.NewCreateUserUsecase(db),
	}
	route.SetupUserRoutes(app, userHander)

	app.Listen(":3000")
}

func connectDB() *gorm.DB {
	var conn = db.Connection{
		Host:     "localhost",
		User:     "postgres",
		Password: "231002",
		DBName:   "postgres",
		Port:     "5432",
	}

	var gormDB, err = db.NewDB(conn)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		panic(err)
	}

	if enableMigration {
		err := gormDB.AutoMigrate(&entity.User{})
		if err != nil {
			slog.Error("failed to migrate database", "error", err)
			return nil
		}
	}

	return gormDB
}
