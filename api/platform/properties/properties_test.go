package properties

import "testing"

func TestAdd(t *testing.T) {
	housing := New()
	housing.Add(Property{})
	if len(housing.Properties) != 1 {
		t.Errorf("Property was not added")
	}
}

func TestGetAll(t *testing.T) {
	housing := New()
	housing.Add(Property{})
	results := housing.GetAll()
	if len(results) != 1 {
		t.Errorf("Property was not added")
	}
}
