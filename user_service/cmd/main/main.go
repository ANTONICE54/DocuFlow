package main

import (
	"auth_service/internal/database"
	"auth_service/internal/server"
	"auth_service/internal/server/handlers"
	"auth_service/internal/usecases"
	"auth_service/pkg/hasher"
	"auth_service/pkg/token"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Failed to read config file:", err)
	}
	dbSource := viper.GetString("DB_SOURCE")
	db := database.Init(dbSource)

	migrationPath := viper.GetString("MIGRATION_PATH")
	database.RunDBMigration(migrationPath, dbSource)

	userRepo := database.NewUserRepo(db)

	secret := viper.GetString("TOKEN_SECRET")
	jwtMaker := token.NewJWTMaker(secret)

	passwordHasher := hasher.NewPasswordHasher()

	httpClient := http.Client{}

	categoryUC := usecases.NewCategoryUC(&httpClient, "http://category_service:8080")

	userUC := usecases.NewUserUC(userRepo, jwtMaker, passwordHasher, categoryUC)
	userHandler := handlers.NewUserHandler(userUC)

	app := server.New(userHandler)

	port := viper.GetString("SERVER_PORT")
	app.Run(port)
}
