package ntsclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aweigor/nts-feed-generator/config"
)

type NTSClient struct {
	baseUrl    string
	httpClient *http.Client
}

func NewNTSClient(cfg *config.NtsAPIProperties) (*NTSClient, error) {
	if cfg.APIV2Url == "" {
		return nil, errors.New("ntsclient: api base required")
	}
	return &NTSClient{
		baseUrl:    cfg.APIV2Url,
		httpClient: &http.Client{},
	}, nil
}

func (client *NTSClient) FetchLatest(ctx context.Context, params EpisodesSearchParams) (*ArticlesResponse, error) {
	query := buildEpisodesSearchQuery(params)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.buildLatestEpisodesURL(query), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed", resp.Status)
	}

	var out ArticlesResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, err
}

func (client *NTSClient) FetchShowEpisodes(ctx context.Context, showId string, params EpisodesSearchParams) (*EpisodesResponse, error) {
	query := buildEpisodesSearchQuery(params)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.buildShowEpisodesURL(showId, query), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed", resp.Status)
	}

	var out EpisodesResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, err
}

func (client *NTSClient) FetchEpisodeInfo(ctx context.Context, showId string, episodeId string) (*EpisodeInfoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed", resp.Status)
	}

	var out EpisodeInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, err
}

func (client *NTSClient) buildShowEpisodesURL(showId string, urlQuery string) string {
	urlPath := strings.ReplaceAll(GetShowEpisodesPath, "{show_id}", showId)
	return fmt.Sprintf("%s%s?%s", client.baseUrl, urlPath, urlQuery)
}

func (client *NTSClient) buildLatestEpisodesURL(urlQuery string) string {
	return fmt.Sprintf("%s%s?%s", client.baseUrl, GetLatestPath, urlQuery)
}
