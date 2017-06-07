package db

import (
	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/boltdb/bolt"
)

// FeedItem is database accessor for FeedItem.
type FeedItem struct {
}

// Create register new feed item in datastore.
func (fi *FeedItem) Create(efi *entity.FeedItem) (*entity.FeedItem, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		return nil
	})

	return nil, err
}
