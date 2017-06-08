package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MasashiSalvador57f/rsser/lib/db"
	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/MasashiSalvador57f/rsser/lib/service"
	"github.com/davecgh/go-spew/spew"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

const (
	registerRSS     = "register_rss"
	listFeeds       = "list_feeds"
	showFeed        = "show_feed"
	registerKeyword = "register_keyword"
	listKeywords    = "list_keywords"
	clearKeywords   = "clear_keywords"
	fetchFeed       = "fetch_feed"
	listFeedItems   = "list_feed_items"
	clearFeedItems  = "clear_feed_items"
	filterFeedItems = "filter_feed_items"
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
	log.Printf("keyword is %s", keyword)

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
		if k.ID <= 0 {
			k = &entity.Keyword{
				Title:   keyword,
				FeedIDs: []uint64{feedID},
			}
		} else {
			if k.HasFeedID(feedID) {
				os.Exit(0)
			}
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
	case clearKeywords:
		keywordDB := new(db.Keyword)
		err := keywordDB.DeleteAll()
		if err != nil {
			log.Fatalf("error in deleting ... %v", err)
		}
		os.Exit(0)
	case showFeed:
		if feedID <= 0 {
			log.Fatal("feed_id is requied")
		}
		fDB := new(db.Feed)
		feed, err := fDB.GetOne(feedID)
		if err != nil {
			log.Fatalf("error in getting a feed %v", err)
		}
		spew.Dump(feed)
		os.Exit(0)
	case fetchFeed:
		if feedID <= 0 {
			log.Fatal("feed_id is required!")
		}

		fetcher := service.NewFetcher(feedID)
		if fetcher == nil {
			log.Fatalf("no feed found with feed_id = %d", feedID)
		}
		err := fetcher.Fetch()
		if err != nil {
			log.Fatalf("error when fetching items of feed with feed_id = %d, : %v", feedID, err)
		}
		fiDB := new(db.FeedItem)
		_, err = fiDB.Batch(fetcher.FeedItems)
		if err != nil {
			log.Fatalf("error in batch creating feed items %v", err)
		}
		os.Exit(0)
	case listFeedItems:
		fiDB := new(db.FeedItem)
		fis, err := fiDB.GetAll()
		if err != nil {
			log.Fatalf("error in GetAll feed_items %v", err)
		}
		for _, fi := range fis {
			log.Println(fi.ToString())
		}
		os.Exit(0)
	case clearFeedItems:
		fiDB := new(db.FeedItem)
		err := fiDB.DeleteAll()
		if err != nil {
			log.Fatalf("error in deleting : %v", err)
		}
		os.Exit(0)
	case filterFeedItems:
		if feedID <= 0 {
			log.Fatal("feed_id is required")
		}
		filter := service.NewFilter()
		feDB := new(db.Feed)
		feed, err := feDB.GetOne(feedID)
		if feed.ID <= 0 || err != nil {
			log.Fatalf("error in getting feed. error: %v", err)
		}
		filter.SetFeed(feed)

		fiDB := new(db.FeedItem)
		feedItems, err := fiDB.GetAll()
		if err != nil {
			log.Fatalf("error in get all feed item. error: %v", err)
		}

		filtered, err := filter.Do(feedItems)
		if err != nil {
			log.Fatalf("error in filtering. error : %v", err)
		}

		for _, item := range filtered {
			fmt.Println(item.ToString())
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
