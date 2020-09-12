package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	hatebuTrend()
}

func qiitaTrend() {
	doc, err := goquery.NewDocument("https://qiita.com/")
	if err != nil {
		log.Fatal(err)
	}
	trend := doc.Find("div[data-hyperapp-app='Trend']")
	val, exists := trend.Attr("data-hyperapp-props")
	if !exists {
		errors.New("data-hyperapp-props does not exist")
	}
	fmt.Println(val)
}
