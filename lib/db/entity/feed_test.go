package entity

import (
	"testing"
)

func Test_getFormattedLastCheckedAt(t *testing.T) {
	cases := []struct {
		RawLastCheckedAt string
		Want             string
	}{
		{"2017-06-07 16:20:29.758723887 +0900 JST", "2017-06-07 16:20:29 +0900 JST"},
		{"2017-06-07", "2017-06-07 00:00:00 +0900 JST"},
	}
	for _, tc := range cases {
		if got := getFormattedLastCheckedAt(tc.RawLastCheckedAt); got != tc.Want {
			t.Errorf("raw=%s, got=%s", tc.RawLastCheckedAt, got)
		}
	}
}

func Test_GetLastCheckedAt(t *testing.T) {
	cases := []struct {
		RawLastCheckedAt string
		Want             string
	}{
		{"2017-06-07 16:20:29.758723887 +0900 JST", "2017-06-07 16:20:29 +0900 JST"},
		{"2017-06-07", "2017-06-07 00:00:00 +0900 JST"},
	}

	for _, tc := range cases {
		f := new(Feed)
		f.LastCheckedAt = tc.RawLastCheckedAt
		if got, err := f.GetLastCheckedAt(); got.String() != tc.Want || err != nil {
			t.Errorf("raw=%s got=%s err=%v", tc.RawLastCheckedAt, got, err)
		}
	}
}
