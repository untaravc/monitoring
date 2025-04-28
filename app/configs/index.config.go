package configs

import (
	"monitoringapp/app/configs/app_config"
	"monitoringapp/app/configs/aws_config"
	"monitoringapp/app/configs/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	aws_config.InitAwsConfig()
	// log_config.DefaultLogging()
	db_config.InitDatabaseConfig()
}
