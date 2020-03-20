package controllers

import (
	"net/http"
)

func DriversController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World api go controller drivers!"))
}
