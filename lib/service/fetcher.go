package service

import (
	"github.com/MasashiSalvador57f/rsser/lib/db"
	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

// Fetcher is ...
type Fetcher struct {
	FeedItems []*entity.FeedItem
	feed      *entity.Feed
}

const feedItemCap = 128

// NewFetcher is ....
func NewFetcher(feedID uint64) *Fetcher {
	fDB := new(db.Feed)

	feed, err := fDB.GetOne(feedID)
	if err != nil || feed.ID <= 0 {
		return nil
	}

	return &Fetcher{
		FeedItems: make([]*entity.FeedItem, 0, feedItemCap),
		feed:      feed,
	}
}

// Fetch is ...
func (f *Fetcher) Fetch() error {
	gp := gofeed.NewParser()
	feed, err := gp.ParseURL(f.feed.URL)
	if err != nil {
		return errors.Wrap(err, "error in fetching feed")
	}

	for _, fi := range feed.Items {
		f.FeedItems = append(f.FeedItems, &entity.FeedItem{
			0, fi,
		})
	}

	return nil
}
