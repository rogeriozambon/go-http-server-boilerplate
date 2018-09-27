package services

import (
	"encoding/json"
	"net/http"
)

type homeHandlerPayload struct{}

// Home handler
func (s Services) Home(w http.ResponseWriter, r *http.Request) {
	var p homeHandlerPayload

	if r.Method != http.MethodPost {
		s.render(w, http.StatusMethodNotAllowed, map[string]string{"message": "invalid request method"})
		return
	}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		s.render(w, http.StatusInternalServerError, map[string]string{"message": "error reading request body"})
		return
	}

	s.db.ExecContext(r.Context(), "")
	s.render(w, http.StatusOK, p)
}
