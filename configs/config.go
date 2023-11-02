package configs

import (
	"github.com/spf13/viper"
)

// Config represents the configuration structure for the application.
type Config struct {
	DBUser     string `json:"dbUser"`     // DBUser is the username for the database connection.
	DBName     string `json:"dbName"`     // DBName is the name of the database.
	DBPassword string `json:"dbPassword"` // DBPassword is the password for the database connection.
	Port       string `json:"port"`       // Port on which the server should run.
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
