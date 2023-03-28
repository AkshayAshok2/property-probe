package search

import "testing"

func TestSearchAdd(t *testing.T) {
	history := New()
	history.Add(Search{})
	if len(history.Searches) != 1 {
		t.Errorf("Search was not added")
	}
}

func TestSearchGetAll(t *testing.T) {
	history := New()
	history.Add(Search{})
	results := history.GetAll()
	if len(results) != 1 {
		t.Errorf("Search was not added")
	}
}
