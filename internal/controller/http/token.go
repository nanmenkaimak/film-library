package http

import (
	"encoding/json"
	"fmt"
	"github.com/nanmenkaimak/film_library/internal/usecase"
	"net/http"
)

func (f *EndpointHandler) Login(w http.ResponseWriter, r *http.Request) {
	var generateTokenRequest usecase.GenerateTokenRequest

	err := json.NewDecoder(r.Body).Decode(&generateTokenRequest)
	if err != nil {
		f.logger.Errorf("failed to unmarshall body err: %v", err)
		handleError(w, fmt.Sprintf("failed to unmarshall body err: %v", err), http.StatusBadRequest)
		return
	}

	userToken, err := f.service.GenerateToken(r.Context(), generateTokenRequest)
	if err != nil {
		f.logger.Errorf("failed to GenerateToken err: %v", err)
		handleError(w, fmt.Sprintf("failed to GenerateToken err: %v", err), http.StatusInternalServerError)
		return
	}

	response := usecase.GenerateTokenResponse{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	renderJSON(w, response)
}

func (f *EndpointHandler) RenewToken(w http.ResponseWriter, r *http.Request) {
	var request usecase.RenewTokenRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		f.logger.Errorf("failed to unmarshall body err: %v", err)
		handleError(w, fmt.Sprintf("failed to unmarshall body err: %v", err), http.StatusBadRequest)
		return
	}

	userToken, err := f.service.RenewToken(r.Context(), request.RefreshToken)
	if err != nil {
		f.logger.Errorf("failed to RenewToken err: %v", err)
		handleError(w, fmt.Sprintf("failed to RenewToken err: %v", err), http.StatusInternalServerError)
		return
	}

	response := usecase.RenewTokenResponse{
		Token: userToken.Token,
	}

	renderJSON(w, response)
}