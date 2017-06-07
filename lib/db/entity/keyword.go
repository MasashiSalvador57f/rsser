package entity

import (
	"fmt"
	"strings"
)

// Keyword is ...
type Keyword struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	FeedID int64  `json:"feed_id"`
}

// ToString gives string format of Keyword.
func (k *Keyword) ToString() string {
	return strings.Join([]string{
		fmt.Sprintf("%d", k.ID),
		k.Title,
		fmt.Sprintf("%d", k.FeedID),
	}, ",")
}
