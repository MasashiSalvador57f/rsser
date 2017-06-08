package service

import (
	"strings"

	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
)

// Filter is ...
type Filter struct {
	ks   *Keyword
	feed *entity.Feed
}

const capFilter = 128

// NewFilter is to filter feed_items.
func NewFilter() *Filter {
	ks := NewKeywordService()
	f := new(Filter)
	f.ks = ks

	return f
}

// SetFeed is ...
func (f *Filter) SetFeed(fe *entity.Feed) {
	f.feed = fe
}

// Do is ...
func (f *Filter) Do(fis []*entity.FeedItem) ([]*entity.FeedItem, error) {
	filtered := make([]*entity.FeedItem, 0, capFilter)

	lct, err := f.feed.GetLastCheckedAt()
	if err != nil {
		return nil, err
	}

	feedID := f.feed.ID
	if len(f.ks.feedIDKeywordsMap) <= 0 {
		return nil, nil
	}
	keywords := f.ks.feedIDKeywordsMap[feedID]
	for _, item := range fis {
		if item.FeedID != feedID {
			continue
		}
		if item.PublishedParsed != nil && lct.Before(*item.PublishedParsed) {
			continue
		}
		for _, k := range keywords {
			if strings.Contains(item.Title, k) || strings.Contains(item.Description, k) || strings.Contains(item.Content, k) {
				filtered = append(filtered, item)
				continue
			}
		}
	}

	return filtered, nil
}
