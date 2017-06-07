package db

import (
	"encoding/json"

	"github.com/MasashiSalvador57f/rsser/lib/db/entity"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// Keyword is to handle database for Keyword.
type Keyword struct{}

// Create is to save new entity of Keyword in DB.
func (k *Keyword) Create(ek *entity.Keyword) (*entity.Keyword, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameKeyword)

		id, _ := b.NextSequence()
		ek.ID = id

		buf, err := json.Marshal(ek)
		if err != nil {
			return errors.Wrap(err, "error in marshaling data")
		}

		return b.Put([]byte(ek.GetKey()), buf)
	})

	return ek, err
}

// GetOne returns one entity that has key : title
func (k *Keyword) GetOne(title string) (*entity.Keyword, error) {
	ek := new(entity.Keyword)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameKeyword)
		v := b.Get([]byte(title))

		if v == nil {
			return nil
		}

		return json.Unmarshal(v, ek)
	})

	return ek, err
}

// DeleteAll is ...
func (k *Keyword) DeleteAll() error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameKeyword)

		return b.ForEach(func(k, v []byte) error {
			return b.Delete(k)
		})
	})

	return err
}

// GetAll returns all the keywords in datastore.
func (k *Keyword) GetAll() ([]*entity.Keyword, error) {
	ks := make([]*entity.Keyword, 0, defaultCapSlice)

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketNameKeyword)

		err := b.ForEach(func(k, v []byte) error {
			keyword := new(entity.Keyword)
			err := json.Unmarshal(v, keyword)

			if err != nil {
				return errors.Wrap(err, "feed unmarshal error")
			}

			ks = append(ks, keyword)
			return nil
		})
		return err
	})

	return ks, err
}
