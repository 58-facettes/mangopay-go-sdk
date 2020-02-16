package model

import (
	"fmt"
	"net/url"
)

// Query is used to hold query params.
type Query struct {
	Page    int
	PerPage int
	Filter  map[string]string
}

// NewPagination creates query params in URL
// page: index of the page (start to 1) Default value: 1
// perPage: number of items returned. Default value: 10 Max: 100
// Here is an example:
// 		users/154876/bank-details?page=2&per_page=10
// If you miss out any parameter in the query the default value is used.
func NewPagination(page, perPage int) *Query {
	page, perPage = defaultValuesForPagnination(page, perPage)
	return &Query{
		Page:    page,
		PerPage: perPage,
	}
}

// NewFilter creates a Query with filters.
func NewFilter(key, value string) *Query {
	return &Query{
		Filter: map[string]string{
			key: value,
		},
	}
}

// AddFilter is adding a filter key value pair to an existing Query.
func (query *Query) AddFilter(key, value string) {
	if query.Filter == nil {
		query.Filter = make(map[string]string)
	}
	query.Filter[key] = value
}

// AddPagination is creating a pagination from an existing Query.
// if this Query has allready a pagination, this will overwrite on it.
func (query *Query) AddPagination(page, perPage int) {
	query.Page, query.PerPage = defaultValuesForPagnination(page, perPage)
}

// URI retive all filters and pagination data to use it with the URI.
func (query *Query) URI(uri string) string {
	u, _ := url.Parse(uri)
	q := u.Query()
	q.Add("page", fmt.Sprint(query.Page))
	q.Add("per_page", fmt.Sprint(query.PerPage))
	if query.Filter != nil {
		for k, val := range query.Filter {
			if len(k) == 0 || len(val) == 0 {
				continue
			}
			q.Add(url.QueryEscape(k), url.QueryEscape(val))
		}
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func defaultValuesForPagnination(page, perPage int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 || perPage > 100 {
		perPage = 10
	}
	return page, perPage
}
