package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"rss/model"

	"github.com/labstack/echo"
	"github.com/mmcdole/gofeed"
)

const GOOGLE_NEWS_RSS_URL = "https://news.google.com/rss/search"

func main() {
	e := echo.New()
	e.GET("/", rss)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func rss(c echo.Context) error {
	url := fmt.Sprintf("%s?q=%s&hl=ja&gl=JP&ceid=JP:ja", GOOGLE_NEWS_RSS_URL, url.QueryEscape("脆弱性"))
	feed, err := gofeed.NewParser().ParseURL(url)
	feeds := []model.Feed{}

	if err != nil {
		return c.String(http.StatusExpectationFailed, "error")
	} else {
		for idx, item := range feed.Items {
			if idx > 2 {
				break
			}
			feeds = append(feeds, model.Feed{
				Title:  item.Title,
				Link:   item.Link,
				Word:   "脆弱性に関するニュース",
				Source: "Google News",
			})
		}
	}
	return c.JSON(http.StatusOK, feeds)
}
