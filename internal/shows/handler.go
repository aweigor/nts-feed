package shows

import (
	"net/http"

	"github.com/aweigor/nts-feed-generator/config"
	"github.com/aweigor/nts-feed-generator/pkg/logger"
	"github.com/aweigor/nts-feed-generator/pkg/ntsclient"
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
	}
}

func (handler *ShowsHandler) HandleTracklist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
