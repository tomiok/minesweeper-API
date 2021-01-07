package api

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest   = GameError{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Cannot process current request"}
	ErrUserNotFound = GameError{StatusCode: http.StatusBadRequest, Type: "user_not_found", Message: "Cannot find the given username"}
	ErrInvalidJSON  = GameError{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
	// decide between conflict and bad request
	ErrAlreadyExists = GameError{StatusCode: http.StatusConflict, Type: "duplicate_entry", Message: "Another entity has the same value as this field"}
	ErrWrongClick    = GameError{StatusCode: http.StatusBadRequest, Type: "click_not_allowed", Message: "May be a cell clicked twice or an invalid click type"}
)

type GameError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (e GameError) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
