package api

import (
	"barber-server/internal/api/config"
	"barber-server/internal/api/database"
	"barber-server/internal/api/handler/route"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// StartServer init and start api server
func StartServer() {
	const ConfigFileName string = "config"
	const ConfigFileType string = "yaml"
	const ConfigFilePath string = "./config"
	// init config file
	err := config.InitConfig(ConfigFileName, ConfigFileType, ConfigFilePath)
	if err != nil {
		log.Fatal("Fail to read config. Error: ", err)
	}

	// init database
	databaseType := viper.GetString("database.type")
	switch databaseType {
	case "mysql":
		database.InitDatabase(database.MysqlDatabase{
			UserName: viper.GetString("mysql.username"),
			Password: viper.GetString("mysql.password"),
			Host:     viper.GetString("mysql.host"),
			Port:     viper.GetString("mysql.port"),
			Database: viper.GetString("mysql.database"),
			CharSet:  viper.GetString("mysql.charset"),
		})
	case "sqlite":
		database.InitDatabase(database.SqliteDatabase{
			FilePath: viper.GetString("sqlite.path"),
		})
	default:
		log.Fatal("Unsupported database type: ", databaseType)
	}

	// start gin service
	release := viper.GetBool("common.release")
	if release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	route.AllRouteCollection(r)
	port := viper.GetString("common.port")
	panic(r.Run(":" + port))
}
