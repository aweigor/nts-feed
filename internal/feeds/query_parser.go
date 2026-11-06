package feeds

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	DEFAULT_FORMAT = "xml"
	DEFAULT_OFFSET = 0
	DEFAULT_LIMIT  = 20
)

type FeedsQueryParser struct{}

func NewFeedsQueryParser() *FeedsQueryParser {
	return &FeedsQueryParser{}
}

func (parser *FeedsQueryParser) parseQuery(q url.Values) (*FeedsQueryParams, error) {
	format := q.Get("format")

	if format == "" {
		format = DEFAULT_FORMAT
	}

	allowedFormats := map[string]bool{
		"json": true,
		"xml":  true,
		"atom": true,
	}

	if !allowedFormats[format] {
		return nil, fmt.Errorf("[FeedsQueryParser]: format must be 'json' or 'xml' or 'atom', got '%s'", format)
	}

	offset := DEFAULT_OFFSET

	queryPage, err := strconv.Atoi(q.Get("page"))
	if err == nil {
		offset = queryPage
	}

	queryOffset, err := strconv.Atoi(q.Get("offset"))
	if err == nil {
		offset = queryOffset
	}

	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		limit = DEFAULT_LIMIT
	}

	result := &FeedsQueryParams{
		format,
		offset,
		limit,
		nil,
		nil,
	}

	sinceStr := q.Get("since")
	if sinceStr != "" {
		since, err := time.Parse(time.RFC3339, sinceStr)
		if err != nil {
			result.since = &since
		}
	}

	beforeStr := q.Get("before")
	if beforeStr != "" {
		before, err := time.Parse(time.RFC3339, beforeStr)
		if err != nil {
			result.before = &before
		}
	}

	return result, nil
}
