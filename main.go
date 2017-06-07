package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MasashiSalvador57f/rsser/lib/db"
	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

const (
	registerRSS     = "register_rss"
	listFeeds       = "list_feeds"
	registerKeyword = "register_keyword"
	listKeywords    = "list_keywords"
)

func main() {
	var (
		commandName string
		keyword     string
		feedID      uint64
		url         string
	)

	flag.StringVar(&commandName, "command", "", "command name you want to execute")
	flag.StringVar(&commandName, "c", "", "command name you want to execute")
	flag.StringVar(&url, "url", "", "feed url")
	flag.StringVar(&keyword, "keyword", "", "keyword for feed")
	flag.Uint64Var(&feedID, "feed_id", 0, "feed_id of keyword")

	flag.Parse()

	if len(commandName) <= 0 {
		log.Fatal("command name is required")
	}

	log.Printf("command is %s", commandName)

	switch commandName {
	case registerRSS:
		log.Println("register rss start")
		feed, err := parseURL(url)
		if err != nil {
			log.Fatalf("no feed %v", err)
		}

		feedDB := new(db.Feed)
		f := &entity.Feed{
			Title:     feed.Title,
			URL:       url,
			UpdatedAt: feed.Updated,
		}
		f, err = feedDB.Create(f)
		if err != nil {
			log.Fatalf("create failed %v", err)
		}

		os.Exit(0)
	case listFeeds:
		feedDB := new(db.Feed)
		fs, err := feedDB.GetAll()
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range fs {
			log.Println(f.ToString())
		}
		os.Exit(0)
	case registerKeyword:
		if len(keyword) <= 0 || feedID <= 0 {
			log.Fatalf("keyword & feed_id is required keyword: %v, feed_id:; %v", keyword, feedID)
		}
		keywordDB := new(db.Keyword)
		k := new(entity.Keyword)

		k, err := keywordDB.GetOne(keyword)
		if err != nil {
			log.Fatalf("error in getting keyword %v", err)
		}
		if k == nil {
			k = &entity.Keyword{
				Title:   keyword,
				FeedIDs: []uint64{feedID},
			}
		} else {
			k.FeedIDs = append(k.FeedIDs, feedID)
		}

		// Create or Update
		k, err = keywordDB.Create(k)
		if err != nil {
			log.Fatalf("error when creating keyword %v", err)
		}

		os.Exit(0)
	case listKeywords:
		keywordDB := new(db.Keyword)
		ks, err := keywordDB.GetAll()
		if err != nil {
			log.Fatalf("error in get data from db %v", err)
		}
		for _, k := range ks {
			fmt.Println(k.ToString())
		}
		os.Exit(0)
	}

	log.Fatalf("no command corresponding to given command name : %s", commandName)
}

func parseURL(url string) (*gofeed.Feed, error) {
	if len(url) <= 0 {
		return nil, errors.New("no url given")
	}
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)

	if err != nil {
		return nil, errors.Wrap(err, "parse failed")
	}
	return feed, nil
}
