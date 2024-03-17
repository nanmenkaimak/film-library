package http

import (
	"encoding/json"
	"fmt"
	"github.com/nanmenkaimak/film_library/internal/usecase"
	"go.uber.org/zap"
	"net/http"
	"unicode"
)

type EndpointHandler struct {
	service usecase.UseCase
	logger  *zap.SugaredLogger
}

func NewEndpointHandler(service usecase.UseCase, logger *zap.SugaredLogger) *EndpointHandler {
	return &EndpointHandler{
		service: service,
		logger:  logger,
	}
}

func validPassword(s string) error {
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
		"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range s {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("password must have at least one %s character", name)
	}
	return nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleError(w http.ResponseWriter, error string, statusCode int) {
	errResp := struct {
		Email string `json:"email"`
	}{
		Email: error,
	}
	js, err := json.Marshal(errResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}