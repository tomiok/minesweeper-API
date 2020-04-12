package api

import "net/http"

func (s *Services) healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("."))
}
