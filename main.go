package main

import (
	"log/slog"
	"os"
	userHandler "smart-waste/apis/user/handlers"
	wastebinHandler "smart-waste/apis/wastebin/handlers"

	_ "smart-waste/docs"

	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	userRoutes "smart-waste/apis/user/routes"
	wastebinRoutes "smart-waste/apis/wastebin/routes"

	userUsecase "smart-waste/domain/user/usecase"
	wastebinUsecase "smart-waste/domain/wastebin/usecase"

	reportEntity "smart-waste/domain/report/entity"
	userEntity "smart-waste/domain/user/entity"
	wasteBinEntity "smart-waste/domain/wastebin/entity"

	"smart-waste/pkgs/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var enableMigration = true

// @title Waste Management API
// @version 1.0
// @description This is a waste management API server.
// @host localhost:3000
// @BasePath /
func main() {
	// Khởi tạo Fiber app
	slog.Info("Service running on port 3000")
	app := fiber.New()

	// Route Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // Mặc định Swagger handler

	// Kết nối đến cơ sở dữ liệu và migrate
	db := connectAndMigrateDB()

	// Khởi tạo các handler và route
	userHandler := userHandler.UserHandler{
		CreateUserUsecase: userUsecase.NewCreateUserUsecase(db),
	}

	wastebinHandler := wastebinHandler.WasteBinHandler{
		CreateWasteBinUsecase: wastebinUsecase.NewCreateWasteBinUsecase(db),
		UpdateWasteBinUsecase: wastebinUsecase.NewUpdateWasteBinUsecase(db),
		DeleteWasteBinUsecase: wastebinUsecase.NewDeleteUserUsecase(db),
		ReadWasteBinUsecase:   wastebinUsecase.NewReadWasteBinUsecase(db),
	}

	// Thiết lập route người dùng
	userRoutes.SetupUserRoutes(app, userHandler)

	// Thiết lập route wastebin
	wastebinRoutes.SetupWasteBinRoutes(app, wastebinHandler)

	// Chạy ứng dụng trên cổng 3000
	app.Listen(":3000")
}

func connectAndMigrateDB() *gorm.DB {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		panic("Failed to load .env file")
	}

	// Lấy thông tin kết nối từ biến môi trường
	conn := db.Connection{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	// Kết nối tới cơ sở dữ liệu
	gormDB, err := db.New(conn)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		panic(err)
	}

	if enableMigration {
		migrateDB(gormDB)
	}

	return gormDB
}
func migrateDB(db *gorm.DB) {
	entities := []interface{}{
		&userEntity.User{},
		&wasteBinEntity.WasteBin{},
		&reportEntity.Report{},
	}

	for _, entity := range entities {
		if err := db.AutoMigrate(entity); err != nil {
			slog.Error("failed to migrate database", "entity", entity, "error", err)
		}
	}
}
