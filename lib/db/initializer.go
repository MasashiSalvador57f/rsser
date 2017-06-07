package db

import (
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

// Initialize is to initialize databse buckets
func init() {
	log.Println("[LOG] initializing DB start")
	d, err := bolt.Open(dbFileName, 0600, nil)
	db = d

	if err != nil {
		log.Fatalf("db file %s is not opened %v", dbFileName, err)
	}
	log.Printf("db is opend, %v", db.Stats())

	createBucketsIfNotExists()
	log.Println("[LOG] DB initialized")
}

func createBucketsIfNotExists() {
	log.Println("[LOG] bucket initialization started")
	createBucketWithName(bucketNameregisteredFeedURL)
	createBucketWithName(bucketNameFeedItem)
	createBucketWithName(bucketNameKeyword)
	log.Println("[LOG] bucket initialization finished")
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
