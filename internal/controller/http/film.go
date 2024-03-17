package http

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/film_library/internal/entity"
	"github.com/nanmenkaimak/film_library/internal/usecase"
	"net/http"
)

func (f *EndpointHandler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	var request entity.FilmeWithActors
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		f.logger.Errorf("failed to unmarshall body err: %v", err)
		handleError(w, fmt.Sprintf("failed to unmarshall body err: %v", err), http.StatusBadRequest)
		return
	}
	filmID, err := f.service.CreateFilm(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to CreateFilm err: %v", err)
		handleError(w, fmt.Sprintf("failed to CreateFilm err: %v", err), http.StatusInternalServerError)
		return
	}
	renderJSON(w, filmID)
}

func (f *EndpointHandler) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	filmID, err := uuid.Parse(r.PathValue("film_id"))
	if err != nil {
		f.logger.Errorf("parsing value from url err: %v", err)
		handleError(w, fmt.Sprintf("parsing value from url err: %v", err), http.StatusBadRequest)
		return
	}
	var request entity.UpdateMap
	err = json.NewDecoder(r.Body).Decode(&request.Values)
	if err != nil {
		f.logger.Errorf("failed to unmarshall body err: %v", err)
		handleError(w, fmt.Sprintf("failed to unmarshall body err: %v", err), http.StatusBadRequest)
		return
	}

	request.ID = filmID

	err = f.service.UpdateFilm(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to UpdateFilm err: %v", err)
		handleError(w, fmt.Sprintf("failed to UpdateFilm err: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (f *EndpointHandler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	filmID, err := uuid.Parse(r.PathValue("film_id"))
	if err != nil {
		f.logger.Errorf("parsing value from url err: %v", err)
		handleError(w, fmt.Sprintf("parsing value from url err: %v", err), http.StatusBadRequest)
		return
	}

	request := usecase.DeleteFilmRequest{
		ID: filmID,
	}
	err = f.service.DeleteFilm(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to DeleteFilm err: %v", err)
		handleError(w, fmt.Sprintf("failed to DeleteFilm err: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (f *EndpointHandler) GetFilms(w http.ResponseWriter, r *http.Request) {
	request := usecase.GetFilmsRequest{
		Sorting: r.URL.Query().Get("sorting"),
		Name:    r.URL.Query().Get("name"),
	}
	films, err := f.service.GetFilms(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to GetFilms err: %v", err)
		handleError(w, fmt.Sprintf("failed to GetFilms err: %v", err), http.StatusInternalServerError)
		return
	}
	renderJSON(w, films)
}
