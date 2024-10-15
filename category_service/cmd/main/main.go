package main

import (
	"category_service/internal/database"
	"category_service/internal/server"
	"category_service/internal/server/handlers"
	"category_service/internal/usecases"
	"log"

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

	categoryRepo := database.NewCategoryRepo(db)
	subcategoryRepo := database.NewSubcategoryRepo(db)

	categoryUC := usecases.NewCategoryUC(categoryRepo)
	subcategoryUC := usecases.NewSubcategoryUC(subcategoryRepo)

	categoryHandler := handlers.NewCategoryHandler(categoryUC)
	subcategoryHandler := handlers.NewSubcategoryHandler(subcategoryUC)

	app := server.New(categoryHandler, subcategoryHandler)
	port := viper.GetString("SERVER_PORT")
	app.Run(port)
}
