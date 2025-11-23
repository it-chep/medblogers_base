package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . AppConfig

type AppConfig interface {
	GetCreateNotificationChatID() int64
	GetSubscribersHost() string
	GetSubscribersPort() string
	GetUserPhotosBucket() string
	GetFreelancersPhotosBucket() string
	GetBlogsPhotosBucket() string
	GetSalebotHost() string
	GetS3Region() string
	GetS3Endpoint() string
	GetS3SecretKey() string
	GetS3AccessKey() string
	GetS3Config() S3Config
	GetAllowedHosts() []string
	GetJWTRefreshSecret() string
	GetJWTSecret() string
}

type Config struct {
	Storage           Storage            `mapstructure:"db"`
	Server            Server             `mapstructure:"server"`
	SubscribersClient SubscribersClient  `mapstructure:"subscribers"`
	SalebotClient     SalebotClient      `mapstructure:"salebot"`
	S3Client          S3Config           `mapstructure:"s3"`
	Notification      NotificationConfig `mapstructure:"notification"`
	AllowedHosts      []string           `mapstructure:"allowed_hosts"`
	JWTConfig         JWTConfig          `mapstructure:"jwt"`
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
	Freelancers string `mapstructure:"freelancers"`
	Blogs       string `mapstructure:"blogs"`
}

type SalebotClient struct {
	Host string `mapstructure:"host"`
}

type Server struct {
	Address     string        `mapstructure:"address"`
	GrpcAddress string        `mapstructure:"grpc_address"`
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

type NotificationConfig struct {
	NotificreateAdminID int64 `mapstructure:"notificreate_admin_id"`
}

type JWTConfig struct {
	Secret        string `mapstructure:"secret"`
	RefreshSecret string `mapstructure:"refresh_secret"`
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

func (c *Config) GetCreateNotificationChatID() int64 {
	return c.Notification.NotificreateAdminID
}

func (c *Config) GetSubscribersHost() string {
	return c.SubscribersClient.Host
}

func (c *Config) GetSubscribersPort() string {
	return c.SubscribersClient.Port
}

func (c *Config) GetUserPhotosBucket() string {
	return c.S3Client.Bucket.UsersPhotos
}

func (c *Config) GetSalebotHost() string {
	return c.SalebotClient.Host
}

func (c *Config) GetS3Region() string {
	return c.S3Client.Region
}

func (c *Config) GetS3Endpoint() string {
	return c.S3Client.Endpoint
}

func (c *Config) GetS3SecretKey() string {
	return c.S3Client.SecretKey
}

func (c *Config) GetS3AccessKey() string {
	return c.S3Client.AccessKey
}

func (c *Config) GetS3Config() S3Config {
	return c.S3Client
}

func (c *Config) GetAllowedHosts() []string {
	return c.AllowedHosts
}

func (c *Config) GetFreelancersPhotosBucket() string {
	return c.S3Client.Bucket.Freelancers
}

func (c *Config) GetBlogsPhotosBucket() string {
	return c.S3Client.Bucket.Blogs
}

func (c *Config) GetJWTRefreshSecret() string {
	return c.JWTConfig.RefreshSecret
}

func (c *Config) GetJWTSecret() string {
	return c.JWTConfig.Secret
}
