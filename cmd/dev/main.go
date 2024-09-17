package main

import (
	"log/slog"
	userHandler "smart-waste/apis/user/handlers"
	wastebinHandler "smart-waste/apis/wastebin/handlers"

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

func main() {
	slog.Info("service running on port 3000")
	app := fiber.New()

	// Connect to database
	db := connectAndMigrateDB()

	// Initialize routes
	initRoutes(app, db)

	app.Listen(":3000")
}

func connectAndMigrateDB() *gorm.DB {
	conn := db.Connection{
		Host:     "localhost",
		User:     "postgres",
		Password: "231002",
		DBName:   "postgres",
		Port:     "5432",
	}

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

func initRoutes(app *fiber.App, db *gorm.DB) {
	// Init handlers and routes with alias names to avoid conflict
	userHandler := userHandler.UserHandler{
		CreateUserUsecase: userUsecase.NewCreateUserUsecase(db),
	}

	wastebinHandler := wastebinHandler.WasteBinHandler{
		CreateWasteBinUsecase: wastebinUsecase.NewCreateWasteBinUsecase(db),
		UpdateWasteBinUsecase: wastebinUsecase.NewUpdateWasteBinUsecase(db),
		DeleteWasteBinUsecase: wastebinUsecase.NewDeleteUserUsecase(db),
		ReadWasteBinUsecase:   wastebinUsecase.NewReadWasteBinUsecase(db),
	}

	// Register user and wastebin routes
	userRoutes.SetupUserRoutes(app, userHandler)
	wastebinRoutes.SetupWasteBinRoutes(app, wastebinHandler)
}
