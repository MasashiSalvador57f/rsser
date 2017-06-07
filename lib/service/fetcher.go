package service

import (
	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/mmcdole/gofeed"
)

// Fetcher is ...
type Fetcher struct {
	FeedItems []*gofeed.Item
	feed      *entity.Feed
}

// NewFetcher is ....
func NewFetcher(feedID uint64) *Fetcher {
	f := new(Fetcher)
}
