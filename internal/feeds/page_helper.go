package feeds

import (
	"context"
	"time"

	"github.com/aweigor/nts-feed-generator/pkg/ntsclient"
)

const (
	MaxLimitLatest  = 60
	MaxLimitShow    = 12
	MaxOffsetLatest = 240
	MaxOffsetShow   = 1000
)

type PageHelper struct {
	*ntsclient.NTSClient
}

func NewPageHelper(client *ntsclient.NTSClient) (*PageHelper, error) {
	return &PageHelper{
		client,
	}, nil
}

func (helper *PageHelper) beforeToLimitLatest(before *time.Time, offset int) int {
	if before == nil {
		return -1
	}
	searchParams := ntsclient.EpisodesSearchParams{
		Page: ntsclient.PageQuery{
			Offset: offset,
			Limit:  MaxLimitLatest,
		},
	}

	ctx := context.Background()
	episodes, err := helper.NTSClient.FetchLatest(ctx, searchParams)
	if err != nil {
		return -1
	}

	limit := MaxLimitLatest
	for i, item := range episodes.Results {
		itemTime, err := time.Parse(time.RFC3339, item.LocalDate)
		if err != nil {
			continue
		}
		if !itemTime.Before(*before) {
			limit = i
		}
	}
	return limit
}

// always 12 since NTS limits shows by const
func (helper *PageHelper) beforeToLimitShows(showId string, before *time.Time, offset int) int {
	if before == nil {
		return -1
	}

	searchParams := ntsclient.EpisodesSearchParams{
		Page: ntsclient.PageQuery{
			Offset: offset,
			Limit:  MaxLimitShow,
		},
	}

	ctx := context.Background()
	episodes, err := helper.NTSClient.FetchShowEpisodes(ctx, showId, searchParams)
	if err != nil {
		return -1
	}

	limit := MaxLimitShow
	for i, item := range episodes.Results {
		itemTime, err := time.Parse(time.RFC3339, item.Updated)
		if err != nil {
			continue
		}
		if !itemTime.Before(*before) {
			limit = i
		}
	}
	return limit
}

func (helper *PageHelper) sinceToOffsetShows(showId string, since *time.Time) int {
	if since == nil {
		return 0
	}
	offset := 0

	for offset < MaxOffsetShow {
		searchParams := ntsclient.EpisodesSearchParams{
			Page: ntsclient.PageQuery{
				Offset: offset,
				Limit:  MaxLimitShow,
			},
		}

		ctx := context.Background()
		episodes, err := helper.NTSClient.FetchShowEpisodes(ctx, showId, searchParams)
		if err != nil {
			return 0
		}
		if len(episodes.Results) == 0 {
			return 0
		}

		for i, item := range episodes.Results {
			itemTime, err := time.Parse(time.RFC3339, item.Updated)
			if err != nil {
				continue
			}
			if itemTime.Before(*since) {
				continue
			}
			return offset + i
		}

		offset += MaxLimitShow
	}

	return offset
}

func (helper *PageHelper) sinceToOffsetLatest(since *time.Time) int {
	if since == nil {
		return 0
	}

	offset := 0

	for offset < MaxOffsetLatest {

		searchParams := ntsclient.EpisodesSearchParams{
			Page: ntsclient.PageQuery{
				Offset: offset,
				Limit:  MaxLimitLatest,
			},
		}

		ctx := context.Background()
		episodes, err := helper.NTSClient.FetchLatest(ctx, searchParams)
		if err != nil {
			return 0
		}
		if len(episodes.Results) == 0 {
			return 0
		}

		for i, item := range episodes.Results {
			itemTime, err := time.Parse(time.RFC3339, item.LocalDate)
			if err != nil {
				continue
			}
			if itemTime.Before(*since) {
				continue
			}
			return offset + i
		}

		offset += MaxLimitLatest
	}

	return offset
}
