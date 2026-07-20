package configs

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Name string `mapstructure:"name"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"app"`
	Mongo struct {
		Url      string `mapstructure:"url"`
		Name     string `mapstructure:"name"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
}

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config") // Config file name without extension
	viper.SetConfigType("yaml")   // Config file type
	viper.AddConfigPath(".")      // Look for the config file in the current directory

	// Читаем файл (если его нет, Viper всё равно сможет читать из ENV)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("ошибка чтения файла конфигурации: %w", err)
		}
		// Если файла нет, просто продолжаем — возможно, всё настроено через ENV
		fmt.Println("Файл конфигурации не найден, используются переменные окружения")
	}

	viper.AutomaticEnv()
	// Разделение в YAML (server.port) заменяется на подчеркивание в ENV (SERVER_PORT)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Create an instance of AppConfig
	var config AppConfig
	// Unmarshal the config file into the AppConfig struct
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	fmt.Printf("%+v\n", config)

	return &config, nil
}
