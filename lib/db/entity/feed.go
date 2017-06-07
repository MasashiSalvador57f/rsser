package entity

// Feed is a struct for feed entity.
type Feed struct {
	ID            uint64 `json:"id"`
	URL           string `json:"url"`
	Title         string `json:"title"`
	UpdatedAt     string `json:"updated_at"`
	LastCheckedAt string `json:"last_checked_at"`
}
