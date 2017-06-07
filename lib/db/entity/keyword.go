package entity

import (
	"fmt"
	"strings"
)

// Keyword is ...
type Keyword struct {
	ID     int64
	Title  string
	FeedID int64
}

// ToString gives string format of Keyword.
func (k *Keyword) ToString() string {
	return strings.Join([]string{
		fmt.Sprintf("%d", k.ID),
		k.Title,
		fmt.Sprintf("%d", k.FeedID),
	}, ",")
}
