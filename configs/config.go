package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	DBDriver          string
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	WebServerPort     string
	GRPCServerPort    string
	GraphQLServerPort string
	RabbitMQUser      string
	RabbitMQPass      string
	RabbitMQHost      string
	RabbitMQPort      string
}

func LoadConfig() (*Conf, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file loaded")
	}

	var cfg Conf

	cfg.DBDriver = os.Getenv("DB_DRIVER")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.WebServerPort = os.Getenv("WEB_SERVER_PORT")
	cfg.GRPCServerPort = os.Getenv("GRPC_SERVER_PORT")
	cfg.GraphQLServerPort = os.Getenv("GRAPHQL_SERVER_PORT")
	cfg.RabbitMQUser = os.Getenv("RABBITMQ_USER")
	cfg.RabbitMQPass = os.Getenv("RABBITMQ_PASS")
	cfg.RabbitMQHost = os.Getenv("RABBITMQ_HOST")
	cfg.RabbitMQPort = os.Getenv("RABBITMQ_PORT")

	return &cfg, nil
}
