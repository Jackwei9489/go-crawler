package main

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/persist"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/zhenai/parser"
)

func main() {
	itemSaver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    itemSaver,
	}
	e.Run(engine.Request{
		Url:       "http://localhost:1234/mock/www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
