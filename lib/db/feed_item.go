package db

import (
	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/boltdb/bolt"
)

// FeedItem is database accessor for FeedItem.
type FeedItem struct {
}

// Create register new feed item in datastore.
func (fi *FeedItem) Create(efi *entity.FeedItem) {
	err := db.Update(func(tx *bolt.Tx) error {

	})
}
