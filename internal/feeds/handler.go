package feeds

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/aweigor/nts-feed-generator/config"
	"github.com/aweigor/nts-feed-generator/pkg/logger"
	"github.com/aweigor/nts-feed-generator/pkg/ntsclient"
	"github.com/aweigor/nts-feed-generator/pkg/res"
	"github.com/google/uuid"
)

type FeedsHandler struct {
	*ntsclient.NTSClient
	*config.FeedsConfig
	*logger.Logger
}

type FeedsHandlerDeps struct {
	*ntsclient.NTSClient
	*config.FeedsConfig
	*logger.Logger
}

func NewFeedsHandler(router *http.ServeMux, deps FeedsHandlerDeps) {
	handler := &FeedsHandler{
		NTSClient:   deps.NTSClient,
		FeedsConfig: deps.FeedsConfig,
		Logger:      deps.Logger,
	}
	router.HandleFunc("/feeds/latest", handler.HandleLatest())
	router.HandleFunc("/feeds/show/{show_id}", handler.HandleShow())
}

func (handler *FeedsHandler) HandleLatest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := logger.WithRequestID(r.Context(), uuid.NewString())

		queryParser := NewFeedsQueryParser()
		queryParams, err := queryParser.parseQuery(r.URL.Query())
		if err != nil {
			res.Error(w, http.StatusBadRequest, err.Error())
		}

		searchParams := handler.buildEpisodesSearchParams(queryParams)
		episodes, err := handler.NTSClient.FetchLatest(r.Context(), searchParams)
		if err != nil {
			handler.Logger.Error(ctx, err.Error())
			res.Error(w, 500, err.Error())
		}

		channelItems := make([]ChannelItem, len(episodes.Results))

		for i := range len(episodes.Results) {
			sourceData := episodes.Results[i]
			channelItems[i] = handler.channelItemFromArticle(&sourceData)
		}

		data := RSSResponse{
			Channel{
				Title: "Latest episodes",
				Items: channelItems,
			},
		}

		switch format := queryParams.format; format {
		case "xml":
			res.Xml(w, data, 200)
		case "json":
			res.Json(w, data, 200)
		}
	}
}

func (handler *FeedsHandler) HandleShow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := logger.WithRequestID(r.Context(), uuid.NewString())

		queryParser := NewFeedsQueryParser()
		queryParams, err := queryParser.parseQuery(r.URL.Query())
		if err != nil {
			res.Error(w, http.StatusBadRequest, err.Error())
		}

		showID := r.PathValue("show_id")

		if showID == "" {
			res.Error(w, http.StatusBadRequest, "Expected show id in route path")
		}

		searchParams := handler.buildEpisodesSearchParams(queryParams)
		episodes, err := handler.NTSClient.FetchShowEpisodes(r.Context(), showID, searchParams)
		if err != nil {
			handler.Logger.Error(ctx, err.Error())
			res.Error(w, 500, err.Error())
		}

		channelItems := make([]ChannelItem, len(episodes.Results))

		for i := range len(episodes.Results) {
			sourceData := episodes.Results[i]
			channelItems[i] = handler.channelItemFromEpisode(&sourceData)
		}

		data := RSSResponse{
			Channel: Channel{
				Title: fmt.Sprintf("%s %s", showID, "shows"),
				Items: channelItems,
			},
		}

		switch format := queryParams.format; format {
		case "xml":
			res.Xml(w, data, 200)
		case "json":
			res.Json(w, data, 200)
		}
	}
}

func (handler *FeedsHandler) buildEpisodeLinkFromArticle(articlePath string) string {
	return fmt.Sprintf("%s%s", handler.Nts.APIV2Url, articlePath)
}

func (handler *FeedsHandler) buildEpisodeLinkFromEpisode(showID string, episodeAlias string) string {
	urlPath := strings.ReplaceAll(ntsclient.GetShowEpisodesPath, "{show_id}", showID)
	return fmt.Sprintf("%s/%s/%s", handler.Nts.APIV2Url, urlPath, episodeAlias)
}

func (handler *FeedsHandler) getEpisodeIdFromArticlePath(articlePath string) string {
	return path.Base(articlePath)
}

func (handler *FeedsHandler) getShowIdFromArticlePath(articlePath string) string {
	parts := strings.Split(path.Dir(articlePath), "/")
	idx := len(parts) - 2
	if idx >= 0 {
		return parts[idx]
	}
	return ""
}

func (handler *FeedsHandler) channelItemFromArticle(articleInfo *ntsclient.ArticleInfo) ChannelItem {
	enclosure := Enclosure{}

	if len(articleInfo.AudioSources) > 0 {
		enclosure.URL = articleInfo.AudioSources[0].URL
		enclosure.Type = articleInfo.AudioSources[0].Source
	}

	return ChannelItem{
		Title:       articleInfo.Title,
		Link:        handler.buildEpisodeLinkFromArticle(articleInfo.Article.Path),
		Description: "",
		PubDate:     articleInfo.LocalDate,
		Enclosure:   enclosure,
		Duration:    nil,
		GUID:        "",
		AirDate:     articleInfo.LocalDate,
		Tracklist:   []Track{},
		ShowID:      handler.getShowIdFromArticlePath(articleInfo.ArticleType),
		EpisodeID:   handler.getEpisodeIdFromArticlePath(articleInfo.Article.Path),
	}
}

func (handler *FeedsHandler) channelItemFromEpisode(episodeInfo *ntsclient.EpisodeInfo) ChannelItem {
	enclosure := Enclosure{}

	if len(episodeInfo.AudioSources) > 0 {
		enclosure.URL = episodeInfo.AudioSources[0].URL
		enclosure.Type = episodeInfo.AudioSources[0].Source
	}

	return ChannelItem{
		Title:       episodeInfo.Name,
		Link:        handler.buildEpisodeLinkFromEpisode(episodeInfo.ShowAlias, episodeInfo.EpisodeAlias),
		Description: episodeInfo.Description,
		PubDate:     episodeInfo.Updated,
		Enclosure:   enclosure,
		Duration:    nil,
		GUID:        "",
		AirDate:     episodeInfo.Broadcast,
		Tracklist:   []Track{},
		ShowID:      episodeInfo.ShowAlias,
		EpisodeID:   episodeInfo.EpisodeAlias,
	}
}

func (handler *FeedsHandler) buildEpisodesSearchParams(query *FeedsQueryParams) ntsclient.EpisodesSearchParams {
	return ntsclient.EpisodesSearchParams{
		Page: ntsclient.PageQuery{
			Offset: query.offset,
			Limit:  query.limit,
		},
	}
}
