package model

import (
	"testing"
)

func TestURI(t *testing.T) {
	data := []struct {
		title    string
		page     int
		perPage  int
		filter   map[string]string
		uri      string
		expected string
	}{
		{"a", 1, 10, nil, "users/", "users/?page=1&per_page=10"},
		{"b", 1, 10, nil, "users/?test=coucou", "users/?page=1&per_page=10&test=coucou"},
		{"c", 1, 10, map[string]string{
			"user_type": "fun",
		}, "users/?test=coucou", "users/?page=1&per_page=10&test=coucou&user_type=fun"},
	}
	for _, d := range data {
		pagination := NewPagination(d.page, d.perPage)
		if d.filter != nil {
			for key, val := range d.filter {
				pagination.AddFilter(key, val)
			}
		}
		if got := pagination.URI(d.uri); got != d.expected {
			t.Errorf("for test %v got %v expected %v", d.title, got, d.expected)
		}
	}
}
