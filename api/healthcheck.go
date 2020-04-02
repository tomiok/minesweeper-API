package api

import "net/http"

func (s *Services) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("."))
}
