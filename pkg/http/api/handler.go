package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type requestHandler struct {
}

func NewAPIHandler() http.Handler {
	handler := &requestHandler{}
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.Methods(http.MethodGet).Path("/health").HandlerFunc(handler.GetInfo)
	apiRouter.Methods(http.MethodGet).Path("/changes").HandlerFunc(handler.GetChanges)
	return router
}

func (h *requestHandler) GetInfo(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}

func (h *requestHandler) GetChanges(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	data := struct {
		Changes []string `json:"changes"`
	}{
		Changes: []string{"A", "B"},
	}

	err := json.NewEncoder(res).Encode(data)
	if err != nil {
		log.WithError(err).Error("Error while writing JSON")
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
