package configs

import (
	"monitoring/app/configs/app_config"
	"monitoring/app/configs/aws_config"
	"monitoring/app/configs/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	aws_config.InitAwsConfig()
	// log_config.DefaultLogging()
	db_config.InitDatabaseConfig()
}
