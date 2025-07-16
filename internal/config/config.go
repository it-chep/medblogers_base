package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Storage           Storage           `mapstructure:"db"`
	Server            Server            `mapstructure:"server"`
	SubscribersClient SubscribersClient `mapstructure:"subscribers"`
	S3Client          S3Config          `mapstructure:"s3"`
}

type S3Config struct {
	Region    string   `mapstructure:"region"`
	Endpoint  string   `mapstructure:"endpoint"`
	AccessKey string   `mapstructure:"access_key"`
	SecretKey string   `mapstructure:"secret_key"`
	Bucket    S3Bucket `mapstructure:"bucket"`
}

type S3Bucket struct {
	UsersPhotos string `mapstructure:"photos"`
}

type Server struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

type Storage struct {
	Host         string        `mapstructure:"host"`
	Port         string        `mapstructure:"port"`
	Database     string        `mapstructure:"database"`
	User         string        `mapstructure:"user"`
	Password     string        `mapstructure:"password"`
	MaxRetry     int           `mapstructure:"max_retry"`
	MaxConnects  int           `mapstructure:"max_connects"`
	RetryTimeout time.Duration `mapstructure:"retry_timeout"`
}

// SubscribersClient todo тк нет serivce discovery и сервис 1 то делаем пока хардкод
type SubscribersClient struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 1. Загружаем .env файл (если есть)
	_ = viper.BindEnv("CONFIG_PATH") // Читаем из переменной окружения

	// 2. Определяем путь к конфигурационному файлу
	configPath := viper.GetString("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// 3. Настраиваем Viper
	viper.SetConfigFile(configPath) // Явно указываем файл конфига

	// 4. Читаем конфигурацию
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 5. Автоматически связываем переменные окружения
	viper.AutomaticEnv()

	// 6. Преобразуем конфиг в структуру
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &cfg
}
