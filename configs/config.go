package configs

// Config represents the configuration structure for the application.

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string // DBUser is the username for the database connection.
	DBName     string // DBName is the name of the database.
	DBPassword string // DBPassword is the password for the database connection.
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")      // specify the file type
	viper.AddConfigPath("./configs") // adjust path if the config file is inside the 'configs' directory

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
