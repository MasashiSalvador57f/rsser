package service

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/mmcdole/gofeed"
)

// Filter is ...
type Filter struct {
	keywords []string
}

const capFilter = 128
const layout = time.RFC3339

// FilterByLastFetched is ...
func (f *Filter) FilterByLastFetched(is []*gofeed.Item, ld string) ([]*gofeed.Item, error) {
	fis := make([]*gofeed.Item, 0, capFilter)
	lastFetchedTime, err := time.Parse(layout, ld)
	if err != nil {
		return nil, err
	}

	spew.Dump(ld)
	spew.Dump(lastFetchedTime)

	for _, fi := range fis {
		fis = append(fis, fi)
	}
	return fis, nil
}
