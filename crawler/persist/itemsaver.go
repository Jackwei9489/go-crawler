package persist

import (
	"errors"
	"github.com/olivere/elastic/v7"
	"go-crawler/crawler/engine"
	"golang.org/x/net/context"
	"log"
)

var (
	IdNotPresent = errors.New("document id not present")
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	itemChan := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-itemChan
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
			err := save(client, index, item)
			if err != nil {
				log.Printf("ItemSaver: error saving item %v %v", item, err)
			}
		}
	}()
	return itemChan, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	if item.Id == "" {
		return IdNotPresent
	}
	_, err := client.
		Index().
		Index(index).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
