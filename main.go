package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	http.HandleFunc("/trend", trendHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func trendHandler(w http.ResponseWriter, r *http.Request) {
	trend, err := qiitaTrend()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, trend)
}

func qiitaTrend() (string, error) {
	doc, err := goquery.NewDocument("https://qiita.com/")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	trend := doc.Find("div[data-hyperapp-app='Trend']")
	val, exists := trend.Attr("data-hyperapp-props")
	if !exists {
		return "", errors.New("data-hyperapp-props does not exist")
	}
	return val, nil
}
