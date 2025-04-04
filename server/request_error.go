package server

import "net/http"

func HttpInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
}

func HttpBadRequestError(w http.ResponseWriter, err error) {
	http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
}

func HttpMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
}

func HttpCreated(w http.ResponseWriter, response any) {
	http.Error(w, "CREATED", http.StatusCreated)
	// TODO - agregar body
}

func HttpOk(w http.ResponseWriter, response any) {
	http.Error(w, "OK", http.StatusOK)
	// TODO - agregar body
}
