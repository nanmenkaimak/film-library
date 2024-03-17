package http

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/film_library/internal/entity"
	"github.com/nanmenkaimak/film_library/internal/usecase"
	"net/http"
)

func (f *EndpointHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	var request entity.Actor
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		f.logger.Errorf("failed to unmarshall body err: %v", err)
		handleError(w, fmt.Sprintf("failed to unmarshall body err: %v", err), http.StatusBadRequest)
		return
	}
	userID, err := f.service.CreateActor(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to CreateActor err: %v", err)
		handleError(w, fmt.Sprintf("failed to CreateActor err: %v", err), http.StatusInternalServerError)
		return
	}
	renderJSON(w, userID)
}

func (f *EndpointHandler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	actorID, err := uuid.Parse(r.PathValue("actor_id"))
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

	request.ID = actorID

	err = f.service.UpdateActor(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to UpdateActor err: %v", err)
		handleError(w, fmt.Sprintf("failed to UpdateActor err: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (f *EndpointHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	actorID, err := uuid.Parse(r.PathValue("actor_id"))
	if err != nil {
		f.logger.Errorf("parsing value from url err: %v", err)
		handleError(w, fmt.Sprintf("parsing value from url err: %v", err), http.StatusBadRequest)
		return
	}
	request := usecase.DeleteActorRequest{
		ID: actorID,
	}
	err = f.service.DeleteActor(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to DeleteActor err: %v", err)
		handleError(w, fmt.Sprintf("failed to DeleteActor err: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (f *EndpointHandler) GetActors(w http.ResponseWriter, r *http.Request) {
	request := usecase.GetActorRequest{
		Name:    r.URL.Query().Get("name"),
	}
	actors, err := f.service.GetActors(r.Context(), request)
	if err != nil {
		f.logger.Errorf("failed to GetActors err: %v", err)
		handleError(w, fmt.Sprintf("failed to GetActors err: %v", err), http.StatusInternalServerError)
		return
	}
	renderJSON(w, actors)
}