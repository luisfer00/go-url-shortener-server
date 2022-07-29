package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func init() {
	ENV := os.Getenv("ENV")
	if ENV == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

var client *redis.Client

func GetRedisClient() *redis.Client {
	DB_URL := os.Getenv("DB_URL")
	DB_PASSWD := os.Getenv("DB_PASSWD")

	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr: DB_URL,
			Password: DB_PASSWD,
		})

		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Connected to redis")
	}
	return client
}