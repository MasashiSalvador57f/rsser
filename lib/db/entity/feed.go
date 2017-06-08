package entity

import (
	"fmt"
	"strings"
	"time"
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

func getFormattedLastCheckedAt(raw string) string {
	ts := strings.Split(raw, " ")
	if len(ts) <= 1 {
		return fmt.Sprintf("%s %s +0900 JST", ts[0], "00:00:00")
	}

	ymd := ts[0]
	dt := strings.Split(ts[1], ".")[0]

	return fmt.Sprintf("%s %s +0900 JST", ymd, dt)
}

// GetLastCheckedAt is ..
func (f *Feed) GetLastCheckedAt() (*time.Time, error) {
	s := getFormattedLastCheckedAt(f.LastCheckedAt)
	t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
