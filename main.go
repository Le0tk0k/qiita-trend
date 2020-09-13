package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/trend", func(c *gin.Context) {
		trend, err := qiitaTrend()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.String(http.StatusOK, trend)
		return
	})

	r.Run(":8080")
}

func qiitaTrend() (string, error) {
	doc, err := goquery.NewDocument("https://qiita.com/")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}
	trend := doc.Find("div[data-hyperapp-app='Trend']")
	val, exists := trend.Attr("data-hyperapp-props")
	if !exists {
		return "", errors.New("Internal Server Error")
	}
	return val, nil
}
