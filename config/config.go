package config

import (
	"github.com/joho/godotenv"
	"os"
)

//AppConfig ...
type AppConfig struct {
	JwtSecret string
	RabbitMQURL string
}

//LoadEnv ...
func LoadEnv() (AppConfig, error) {
  err := godotenv.Load()

  if err != nil {
	  return AppConfig{}, err
  }

  return AppConfig{
	  JwtSecret: os.Getenv("JWT_SECRET"),
	  RabbitMQURL: os.Getenv("RABBITMQ_URL"),
  }, nil

}

