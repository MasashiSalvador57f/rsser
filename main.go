package main

import (
	"flag"
	"log"
	"os"

	"github.com/MasashiSalvador57f/rsser/lib/db"
	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

const (
	registerRSS = "register_rss"
	listFeeds   = "list_feeds"
)

func main() {
	var (
		commandName string
		url         string
	)

	flag.StringVar(&commandName, "command", "", "command name you want to execute")
	flag.StringVar(&commandName, "c", "", "command name you want to execute")
	flag.StringVar(&url, "url", "", "feed url")
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
