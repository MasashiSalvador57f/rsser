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

const defaultCapSlice = 255

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

// GetOne is ...
func (fdb *Feed) GetOne(feedID uint64) (*entity.Feed, error) {
	ef := new(entity.Feed)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameregisteredFeedURL)
		v := b.Get(itob(feedID))

		if v == nil {
			return nil
		}

		return json.Unmarshal(v, ef)
	})

	return ef, err
}

// GetAll is to get all data in bucket.
func (fdb *Feed) GetAll() ([]*entity.Feed, error) {
	fs := make([]*entity.Feed, 0, defaultCapSlice)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameregisteredFeedURL)

		err := b.ForEach(func(k, v []byte) error {
			f := new(entity.Feed)
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
