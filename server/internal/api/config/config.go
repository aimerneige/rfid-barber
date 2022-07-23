package config

import "github.com/spf13/viper"

// InitConfig init viper config
func InitConfig(fileName, fileType, filePath string) error {
	// config file
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	// config file path
	viper.AddConfigPath(filePath)
	// try to load config file
	err := viper.ReadInConfig()
	return err
}
