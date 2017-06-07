package entity

import (
	"fmt"
	"strings"
)

// Keyword is ...
type Keyword struct {
	ID      uint64   `json:"id"`
	Title   string   `json:"title"`
	FeedIDs []uint64 `json:"feed_ids"`
}

// ToString gives string format of Keyword.
func (k *Keyword) ToString() string {
	return strings.Join([]string{
		fmt.Sprintf("%d", k.ID),
		k.Title,
		fmt.Sprintf("%v", k.FeedIDs),
	}, ",")
}

// GetKey returns key for kvs.
func (k *Keyword) GetKey() []byte {
	return []byte(k.Title)
}
