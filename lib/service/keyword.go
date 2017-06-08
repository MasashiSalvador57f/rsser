package service

import (
	"github.com/MasashiSalvador57f/rsser/lib/db"
)

// Keyword is ..
type Keyword struct {
	feedIDKeywordsMap map[uint64][]string
	isInitialized     bool
}

var fk map[uint64][]string

// NewKeywordService is ...
func NewKeywordService() *Keyword {
	ks := new(Keyword)

	if len(fk) <= 0 {
		fk = initializeFeedIDKeywordMap()
	}
	ks.feedIDKeywordsMap = fk

	return ks
}

func initializeFeedIDKeywordMap() map[uint64][]string {
	fk := make(map[uint64][]string)
	keywordDB := new(db.Keyword)
	allKeywords, err := keywordDB.GetAll()
	if err != nil {
		return nil
	}

	for _, keyword := range allKeywords {
		for _, feedID := range keyword.FeedIDs {
			_, ok := fk[feedID]
			if ok {
				fk[feedID] = append(fk[feedID], keyword.Title)
			} else {
				fk[feedID] = []string{keyword.Title}
			}
		}
	}
	return fk
}
