package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNoContent)

	w.Write([]byte("Hello"))
}
