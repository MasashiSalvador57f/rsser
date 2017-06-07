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

		return b.Put(itob(id), buf)
	})

	return ek, err
}
