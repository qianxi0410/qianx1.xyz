package utils

import "testing"

func TestRss(t *testing.T) {
	if err := Rss(); err != nil {
		t.Fail()
	}
}
