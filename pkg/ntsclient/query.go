package ntsclient

import (
	"net/url"
	"strconv"
)

const (
	DefaultOffset = 0
	DefaultLimit  = 20
)

// PageQuery describes common page query parameters
type PageQuery struct {
	Offset int
	Limit  int
}

// EpisodesQueryParams describes query parameters of Episodes search request
type EpisodesSearchParams struct {
	Page PageQuery
}

// buildEpisodesQuery builds episodes search query string
func buildEpisodesSearchQuery(params EpisodesSearchParams) string {
	offset := params.Page.Offset
	limit := params.Page.Limit

	if offset < 0 {
		offset = DefaultOffset
	}

	if limit <= 0 {
		limit = DefaultLimit
	}

	q := url.Values{}
	q.Set("offset", strconv.Itoa(offset))
	q.Set("limit", strconv.Itoa(limit))

	return q.Encode()
}
