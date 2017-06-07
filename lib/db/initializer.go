package db

import (
	"log"

	"github.com/boltdb/bolt"
)

var db bolt.DB

// Initialize is to initialize databse buckets
func init() {
	db, err := bolt.Open(dbFileName, 0600, nil)
	if err != nil {
		log.Fatalf("db file %s is not opened %v", dbFileName, err)
	}
	log.Printf("db is opend, %v", db.Stats())

	createBucketsIfNotExists()
}
func createBucketsIfNotExists() {
	createBucketWithName(bucketNameregisteredFeedURL)
}

func createBucketWithName(bucketName []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			log.Fatalf("bucket name %s is not created with error: %v", bucketName, err)
		}
		return err
	})

	return err
}
