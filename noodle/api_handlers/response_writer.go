package api_handlers

import "net/http"

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --case underscore --name ResponseWriter
type ResponseWriter interface {
	http.ResponseWriter
}
