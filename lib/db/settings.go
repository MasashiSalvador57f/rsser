package db

import (
	"encoding/binary"
)

const dbFileName = "rsser.db"

const (
	// NameRegisteredFeedURL is the name for bucket for registered feed urls
	nameRegisteredFeedURL = "registered_feeds"
)

var (
	bucketNameregisteredFeedURL = []byte(nameRegisteredFeedURL)
)

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
