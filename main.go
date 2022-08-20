package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"rss/infrastructure/application/basic"
	"rss/infrastructure/datasource/trends"
	"rss/model"
	"rss/usecase"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/mmcdole/gofeed"
)

const GOOGLE_NEWS_RSS_URL = "https://news.google.com/rss/search"

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

func rss(c echo.Context) error {
	word := c.FormValue("word")
	url := fmt.Sprintf("%s?q=%s&hl=ja&gl=JP&ceid=JP:ja", GOOGLE_NEWS_RSS_URL, url.QueryEscape(word))
	feed, err := gofeed.NewParser().ParseURL(url)
	feeds := []model.Feed{}

	if err != nil {
		return c.String(http.StatusOK, "error")
	} else {
		for idx, item := range feed.Items {
			if idx > 0 {
				break
			}
			feeds = append(feeds, model.Feed{
				Title:  item.Title,
				Link:   item.Link,
				Word:   word,
				Source: "Google News",
			})
		}
	}
	return c.JSON(http.StatusOK, feeds)
}
