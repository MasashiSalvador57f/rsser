package db

import (
	"encoding/json"

	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// FeedItem is database accessor for FeedItem.
type FeedItem struct {
}

// Batch is ...
func (fi *FeedItem) Batch(efis []*entity.FeedItem) ([]*entity.FeedItem, error) {
	err := db.Batch(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameFeedItem)
		for _, efi := range efis {
			id, _ := b.NextSequence()
			efi.ID = id

			buf, err := json.Marshal(efi)
			if err != nil {
				return errors.Wrap(err, "error in marshaling json")
			}

			b.Put(itob(id), buf)
		}

		return nil
	})

	return efis, err
}

// GetAll is to get all data in bucket.
func (fi FeedItem) GetAll() ([]*entity.FeedItem, error) {
	fs := make([]*entity.FeedItem, 0, defaultCapSlice)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameFeedItem)

		err := b.ForEach(func(k, v []byte) error {
			f := new(entity.FeedItem)
			err := json.Unmarshal(v, f)

			if err != nil {
				return errors.Wrap(err, "feed unmarshal error")
			}
			fs = append(fs, f)
			return nil
		})
		return err
	})
	return fs, err
}

// DeleteAll is ...
func (fi *FeedItem) DeleteAll() error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameFeedItem)

		return b.ForEach(func(k, v []byte) error {
			return b.Delete(k)
		})
	})

	return err
}
