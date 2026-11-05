package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/aweigor/nts-feed-generator/config"
	"github.com/aweigor/nts-feed-generator/internal/feeds"
	"github.com/aweigor/nts-feed-generator/internal/heartbeat"
	"github.com/aweigor/nts-feed-generator/internal/shows"
	"github.com/aweigor/nts-feed-generator/pkg/logger"
	"github.com/aweigor/nts-feed-generator/pkg/middleware"
	"github.com/aweigor/nts-feed-generator/pkg/ntsclient"
)

func NewApp() http.Handler {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger := logger.NewLogger(slogLogger)

	router := http.NewServeMux()
	heartbeat.NewHeartbeatHandler(router, conf)

	ntsClient, err := ntsclient.NewNTSClient(&conf.Nts)
	if err != nil {
		log.Fatalf("failed to initialize NTS client", err)
	}

	feeds.NewFeedsHandler(router, feeds.FeedsHandlerDeps{
		NTSClient: ntsClient,
		FeedsConfig: &config.FeedsConfig{
			Nts:       conf.Feeds.Nts,
			Tracklist: conf.Feeds.Tracklist,
		},
		Logger: logger,
	})

	shows.NewShowsHandler(router, shows.ShowsHandlerDeps{
		NTSClient: ntsClient,
		ShowsConfig: &config.ShowsConfig{
			Nts: conf.Nts,
		},
		Logger: logger,
	})

	mwStack := middleware.Chain(middleware.CORS, middleware.Logging)

	return mwStack(router)
}

func main() {
	app := NewApp()
	server := http.Server{Addr: ":9000", Handler: app}
	fmt.Println("Server is listening on port 9000")
	server.ListenAndServe()
}
