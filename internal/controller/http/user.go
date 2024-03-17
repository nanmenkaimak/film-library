package http

import (
	"encoding/json"
	"fmt"
	"github.com/nanmenkaimak/film_library/internal/entity"
	"net/http"
)

func (f *EndpointHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request entity.User

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		f.logger.Errorf("failed to unmarshall body err: %v", err)
		handleError(w, fmt.Sprintf("failed to unmarshall body err: %v", err), http.StatusBadRequest)
		return
	}
	userID, err := f.service.RegisterUser(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to Register err: %v", err)
		handleError(w, fmt.Sprintf("failed to Register err: %v", err), http.StatusInternalServerError)
		return
	}
	renderJSON(w, userID)
}
