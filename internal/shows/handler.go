package shows

import (
	"net/http"

	"github.com/aweigor/nts-feed-generator/config"
	"github.com/aweigor/nts-feed-generator/pkg/logger"
	"github.com/aweigor/nts-feed-generator/pkg/ntsclient"
	"github.com/aweigor/nts-feed-generator/pkg/res"
	"github.com/google/uuid"
)

type ShowsHandler struct {
	*ntsclient.NTSClient
	*config.ShowsConfig
	*logger.Logger
}

type ShowsHandlerDeps struct {
	*ntsclient.NTSClient
	*config.ShowsConfig
	*logger.Logger
}

func NewShowsHandler(router *http.ServeMux, deps ShowsHandlerDeps) {
	handler := &ShowsHandler{
		NTSClient:   deps.NTSClient,
		ShowsConfig: deps.ShowsConfig,
		Logger:      deps.Logger,
	}
	router.HandleFunc("/shows/{show_id}/episodes/{episode_id}", handler.HandleEpisode())
	router.HandleFunc("/shows/{show_id}/episodes/{episode_id}/tracklist", handler.HandleTracklist())
}

func (handler *ShowsHandler) HandleEpisode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := logger.WithRequestID(r.Context(), uuid.NewString())

		showId := r.PathValue("show_id")
		episodeId := r.PathValue("episode_id")

		if showId == "" || episodeId == "" {
			res.Error(w, http.StatusBadRequest, "show_id and episode_id must be set")
		}

		episode, err := handler.NTSClient.FetchEpisodeInfo(ctx, showId, episodeId)
		if err != nil {
			handler.Logger.Error(ctx, err.Error())
			res.Error(w, 500, err.Error())
		}

		data := EpisodeInfoResponse{
			Status:          episode.Status,
			Updated:         episode.Updated,
			Name:            episode.Name,
			Description:     episode.Description,
			DescriptionHTML: episode.DescriptionHTML,
			ExternalLinks:   episode.ExternalLinks,
			Moods:           episode.Moods,
			Genres:          episode.Genres,
			LocationShort:   episode.LocationShort,
			LocationLong:    episode.LocationLong,
			Intensity:       episode.Intensity,
			Media:           Media(episode.Media),
			EpisodeAlias:    episode.EpisodeAlias,
			ShowAlias:       episode.ShowAlias,
			Broadcast:       episode.Broadcast,
			Mixcloud:        episode.Mixcloud,
			AudioSources:    episode.AudioSources,
			Brand:           episode.Brand,
			Tracklist:       episode.Embeds.Tracklist.Results,
		}
		res.Json(w, data, 200)
	}
}

func (handler *ShowsHandler) HandleTracklist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := logger.WithRequestID(r.Context(), uuid.NewString())

		showId := r.PathValue("show_id")
		episodeId := r.PathValue("episode_id")

		if showId == "" || episodeId == "" {
			res.Error(w, http.StatusBadRequest, "show_id and episode_id must be set")
		}

		episode, err := handler.NTSClient.FetchEpisodeInfo(ctx, showId, episodeId)
		if err != nil {
			handler.Logger.Error(ctx, err.Error())
			res.Error(w, 500, err.Error())
		}

		data := EpisodeTracklistResponse{
			Tracklist: episode.Embeds.Tracklist.Results,
		}

		res.Json(w, data, 200)
	}
}
