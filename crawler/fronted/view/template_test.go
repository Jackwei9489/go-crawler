package view

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/fronted/model"
	model2 "go-crawler/crawler/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	view := template.Must(template.ParseFiles("template.html"))
	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "http://album.zhenai.com/u/108906739",
		Id:  "108906739",
		Payload: model2.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Execute(out, page)
	if err != nil {
		t.Error(err)
	}
}
