package db

import (
	"encoding/json"
	"time"

	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// Feed is a store handler for Feed.
type Feed struct {
}

// Create is to create new feed url in bucket.
func (fdb *Feed) Create(f *entity.Feed) (*entity.Feed, error) {

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameregisteredFeedURL)

		id, _ := b.NextSequence()
		f.ID = id
		f.LastCheckedAt = time.Now().String()

		buf, err := json.Marshal(f)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "faied to marshal json")
		}

		return b.Put(itob(f.ID), buf)
	})
	if err != nil {
		return nil, err
	}

	return f, err
}
