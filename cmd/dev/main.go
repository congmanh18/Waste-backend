package main

import (
	"log/slog"
	userHandler "smart-waste/apis/user/handlers"
	wastebinHandler "smart-waste/apis/wastebin/handler"

	user "smart-waste/apis/user/routes"
	wastebin "smart-waste/apis/wastebin/router"

	userEntity "smart-waste/domain/user/entity"
	wastebinEntity "smart-waste/domain/wastebin/entity"

	userUsecase "smart-waste/domain/user/usecase"
	wasteBinUsecase "smart-waste/domain/wastebin/usecase"
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
	var userHander = userHandler.UserHandler{
		CreateUserUsecase: userUsecase.NewCreateUserUsecase(db),
	}

	var wastebinHandler = wastebinHandler.WasteBinHandler{
		CreateWasteBinUsecase: wasteBinUsecase.NewCreateWasteBinUsecase(db),
	}
	user.SetupUserRoutes(app, userHander)
	wastebin.SetupWasteBinRoutes(app, wastebinHandler)

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
		err := gormDB.AutoMigrate(&userEntity.User{})
		if err != nil {
			slog.Error("failed to migrate User database", "error", err)
			return nil
		}
	}

	if enableMigration {
		err := gormDB.AutoMigrate(&wastebinEntity.WasteBin{})
		if err != nil {
			slog.Error("failed to migrate WasteBin database", "error", err)
			return nil
		}
	}
	return gormDB
}
