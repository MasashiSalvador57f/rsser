package entity

import (
	"fmt"
	"strings"
)

// Feed is a struct for feed entity.
type Feed struct {
	ID            uint64 `json:"id"`
	URL           string `json:"url"`
	Title         string `json:"title"`
	UpdatedAt     string `json:"updated_at"`
	LastCheckedAt string `json:"last_checked_at"`
}

// ToString is to give string format of f.
func (f *Feed) ToString() string {
	return strings.Join([]string{
		fmt.Sprintf("%d", f.ID),
		f.URL,
		f.Title,
	}, ",")
}
