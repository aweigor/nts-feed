package heartbeat

import (
	"net/http"

	"github.com/aweigor/nts-feed-generator/config"
	"github.com/aweigor/nts-feed-generator/pkg/res"
)

type HeartbeatHandler struct{}

func NewHeartbeatHandler(router *http.ServeMux, conf *config.Config) {
	handler := &HeartbeatHandler{}
	router.HandleFunc("/explain-config", handler.PrintConfig(conf))
	router.HandleFunc("/heartbeat", handler.Hello())
}

func (handler *HeartbeatHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("OK"))
	}
}

func (handler *HeartbeatHandler) PrintConfig(conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		res.Json(w, conf, http.StatusOK)
	}
}
