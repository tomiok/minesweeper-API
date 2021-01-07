package api

import (
	"encoding/json"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type ResponseAPI struct {
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type LostResponse struct {
	Message      string        `json:"message"`
	Clicks       int           `json:"clicks"`
	GameDuration time.Duration `json:"game_duration"`
}

func (r *ResponseAPI) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	err := json.NewEncoder(w).Encode(r)

	if err != nil {
		logs.Log().Error("cannot process the response", zap.Error(err))
	}
}

func Success(result interface{}, status int) *ResponseAPI {
	return &ResponseAPI{
		Success: true,
		Status:  status,
		Result:  result,
	}
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
