package main

import (
	"log/slog"
	"os"
	reportHandler "smart-waste/apis/report/handlers"
	userHandler "smart-waste/apis/user/handlers"
	wastebinHandler "smart-waste/apis/wastebin/handlers"

	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	reportRoutes "smart-waste/apis/report/routes"
	userRoutes "smart-waste/apis/user/routes"
	wastebinRoutes "smart-waste/apis/wastebin/routes"

	reportUsecase "smart-waste/domain/report/usecase"
	userUsecase "smart-waste/domain/user/usecase"
	wastebinUsecase "smart-waste/domain/wastebin/usecase"

	reportEntity "smart-waste/domain/report/entity"
	userEntity "smart-waste/domain/user/entity"
	wasteBinEntity "smart-waste/domain/wastebin/entity"

	userRepo "smart-waste/domain/user/repository"
	wasteBinRepo "smart-waste/domain/wastebin/repository"

	"smart-waste/pkgs/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var enableMigration = false

// @title Smart Waste Management API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email nguyenmanh180102@gmail.com
// @license.name Nginx 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
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
	userRepo := userRepo.NewUserRepo(db)
	userHandler := userHandler.UserHandler{
		CreateUserUsecase:     userUsecase.NewCreateUserUsecase(db),
		GetUserByPhoneUsecase: userUsecase.NewGetUserByPhoneUsecase(db),
		UpdateUserUsecase:     userUsecase.NewUpdateUserUsecase(db),
		DeleteUserUsecase:     userUsecase.NewDeleteUserUsecase(db),
		FindUserByIDUsecase:   userUsecase.NewFindUserByIDUsecase(db),
		FindAllUserUsecase:    userUsecase.NewFindAllUserUsecase(db),
	}

	wastebinRepo := wasteBinRepo.NewWasteBinRepo(db)
	wastebinHandler := wastebinHandler.WasteBinHandler{
		CreateWasteBinUsecase:  wastebinUsecase.NewCreateWasteBinUsecase(db),
		UpdateWasteBinUsecase:  wastebinUsecase.NewUpdateWasteBinUsecase(db),
		DeleteWasteBinUsecase:  wastebinUsecase.NewDeleteUserUsecase(db),
		ReadWasteBinUsecase:    wastebinUsecase.NewReadWasteBinUsecase(db),
		ReadAllWasteBinUsecase: wastebinUsecase.NewReadAllWasteBinUsecase(db),
	}

	reportHandler := reportHandler.ReportHandler{
		CreateReportUsecase:           reportUsecase.NewCreateReportUsecase(db),
		DeleteReportUsecase:           reportUsecase.NewDeleteReportUsecase(db),
		GetAllReportsUsecase:          reportUsecase.NewGetAllReportsUsecase(db),
		GetLast:                       reportUsecase.NewGetLatestByWasteBinID(db),
		GetReportByIDUsecase:          reportUsecase.NewGetReportByIDUsecase(db),
		GetReportsByDateUsecase:       reportUsecase.NewGetReportsByDateUsecase(db),
		GetReportsByUserIDUsecase:     reportUsecase.NewGetReportsByUserIDUsecase(db),
		GetReportsByWasteBinIDUsecase: reportUsecase.NewGetReportsByWasteBinIDUsecase(db),
		WasteBinRepo:                  wastebinRepo,
		UserRepo:                      userRepo,
	}

	userRoutes.SetupUserRoutes(app, userHandler)
	wastebinRoutes.SetupWasteBinRoutes(app, wastebinHandler)
	reportRoutes.SetupReportRoutes(app, reportHandler)

	wastebinRoutes.SetupExponentialSmoothingRoutes(app, wastebinHandler)

	if err := app.Listen(":3000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		panic(err)
	}
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

	// Kiểm tra xem các biến môi trường đã được load chưa
	if conn.Host == "" || conn.User == "" || conn.DBName == "" {
		slog.Error("Missing environment variables for DB connection")
		panic("Missing environment variables for DB connection")
	}

	// Kết nối tới cơ sở dữ liệu
	gormDB, err := db.New(conn)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		panic(err)
	}

	// Kiểm tra xem DB có được khởi tạo không
	if gormDB == nil {
		panic("Database connection is nil")
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

// docker push manh18/final:latest^C
// docker tag final manh18/final2:latest
// docker build -t final2 .
