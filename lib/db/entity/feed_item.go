package entity

import (
	"fmt"
	"strings"

	"github.com/mmcdole/gofeed"
)

// FeedItem is corresponding to item in feeds.
type FeedItem struct {
	ID uint64 `json:"id"`
	*gofeed.Item
}

// ToString gives string format of feed item
func (f *FeedItem) ToString() string {
	return strings.Join([]string{
		fmt.Sprintf("%d", f.ID),
		f.Title,
		f.Description,
	}, ",")
}
