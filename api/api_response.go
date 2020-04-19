package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResponseAPI struct {
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func Success(result interface{}, status int) *ResponseAPI {
	return &ResponseAPI{
		Success: true,
		Status:  status,
		Result:  result,
	}
}

func (r *ResponseAPI) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}

func LostGame(clicks int, username string) *ResponseAPI {
	return &ResponseAPI{
		Success: true,
		Status:  http.StatusOK,
		Result: &LostResponse{
			Message: "You lost, " + username,
			Clicks:  clicks,
		},
	}
}

type LostResponse struct {
	Message      string        `json:"message"`
	Clicks       int           `json:"clicks"`
	GameDuration time.Duration `json:"game_duration"`
}
