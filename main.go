package main

import (
	"log"
	"os"
	"rss/infrastructure/application/basic"
	"rss/infrastructure/datasource/trends"
	"rss/usecase"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] failed to load .env file: %w", err)
	}
}

func main() {
	var (
		TWITTER_API_KEY       = os.Getenv("TWITTER_API_KEY")
		TWITTER_API_SECRET    = os.Getenv("TWITTER_API_SECRET")
		TWITTER_ACCESS_TOKEN  = os.Getenv("TWITTER_ACCESS_TOKEN")
		TWITTER_ACCESS_SECRET = os.Getenv("TWITTER_ACCESS_SECRET")
		port, _               = strconv.Atoi(os.Getenv("PORT"))
	)
	twitter := trends.NewTwitter(TWITTER_API_KEY, TWITTER_API_SECRET, TWITTER_ACCESS_TOKEN, TWITTER_ACCESS_SECRET)
	usecase := usecase.NewUseCase(twitter)
	application := basic.NewApplication(usecase)
	application.Run(port)
}
