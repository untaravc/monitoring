package bootstrap

import (
	// "context"
	"log"
	"monitoring/app/configs"
	"monitoring/app/configs/app_config"
	"monitoring/app/routes"
	"monitoring/app/utils/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	configs.InitConfig()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			log.Println("Running migrations...")
			// commands.RunMigrationPostgres()
			return
		}
	}

	// if app_config.GIN_MODE == "release" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }
	app := gin.Default()

	// app.Use(cors_config.CORSMiddleware())
	// app.Use(log_config.LoggerMiddleware())

	database.InitMysqlDatabase()
	defer database.DB.Close()

	routes.InitRoute(app)

	if err := app.Run(app_config.PORT); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
