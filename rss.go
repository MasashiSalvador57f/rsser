package main

import (
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

func parseURL(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, errors.Wrap(err, "parse failed")
	}

	return feed, nil
}
