package api_handlers

import "net/http"

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name ResponseWriterTest
type ResponseWriterTest interface {
	http.ResponseWriter
}
