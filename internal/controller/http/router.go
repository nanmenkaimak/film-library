package http

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type router struct {
	logger *zap.SugaredLogger
}

func NewRouter(logger *zap.SugaredLogger) *router {
	return &router{
		logger: logger,
	}
}

func (s *router) GetHandler(eh *EndpointHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			eh.Register(w, r)
		default:
			handleError(w, fmt.Sprintf("Method not allowed"), http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			eh.Login(w, r)
		default:
			handleError(w, fmt.Sprintf("Method not allowed"), http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/renew-token", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			eh.RenewToken(w, r)
		default:
			handleError(w, fmt.Sprintf("Method not allowed"), http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/actor", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			eh.CreateActor(w, r)
		case http.MethodGet:
			eh.GetActors(w, r)
		default:
			handleError(w, fmt.Sprintf("Method not allowed"), http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/actor/{actor_id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			eh.UpdateActor(w, r)
		case http.MethodDelete:
			eh.DeleteActor(w, r)
		default:
			handleError(w, fmt.Sprintf("Method not allowed"), http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/film", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			eh.CreateFilm(w, r)
		case http.MethodGet:
			eh.GetFilms(w, r)
		default:
			handleError(w, fmt.Sprintf("Method not allowed"), http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/film/{film_id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			eh.UpdateFilm(w, r)
		case http.MethodDelete:
			eh.DeleteFilm(w, r)
		default:
			handleError(w, fmt.Sprintf("Method not allowed"), http.StatusMethodNotAllowed)
		}
	})
	return mux
}
