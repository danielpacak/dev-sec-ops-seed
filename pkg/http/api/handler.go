package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

type requestHandler struct {
	buildInfo BuildInfo
}

func NewAPIHandler(buildInfo BuildInfo) http.Handler {
	handler := &requestHandler{
		buildInfo: buildInfo,
	}
	router := mux.NewRouter()
	v1Router := router.PathPrefix("/api").Subrouter()

	v1Router.Methods(http.MethodGet).Path("/health").HandlerFunc(handler.GetHealth)
	v1Router.Methods(http.MethodGet).Path("/info").HandlerFunc(handler.GetInfo)
	return router
}

func (h *requestHandler) GetHealth(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}

func (h *requestHandler) GetInfo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	err := json.NewEncoder(res).Encode(h.buildInfo)
	if err != nil {
		log.WithError(err).Error("Error while writing JSON")
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
