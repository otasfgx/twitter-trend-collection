package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo"
	"github.com/mmcdole/gofeed"
)

const GOOGLE_NEWS_RSS_URL = "https://news.google.com/rss/search"

type Feed struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Word   string `json:"word"`
	Source string `json:"source"`
}

func main() {
	e := echo.New()
	e.GET("/", rss)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func rss(c echo.Context) error {
	url := fmt.Sprintf("%s?q=%s&hl=ja&gl=JP&ceid=JP:ja", GOOGLE_NEWS_RSS_URL, url.QueryEscape("脆弱性"))
	feed, err := gofeed.NewParser().ParseURL(url)
	feeds := []Feed{}

	if err != nil {
		return c.String(http.StatusExpectationFailed, "error")
	} else {
		for idx, item := range feed.Items {
			if idx > 2 {
				break
			}
			feeds = append(feeds, Feed{
				Title:  item.Title,
				Link:   item.Link,
				Word:   "脆弱性に関するニュース",
				Source: "Google News",
			})
		}
	}
	return c.JSON(http.StatusOK, feeds)
}
