package main

import (
	"fmt"
	"net/http"
)

// ResponseWriter mocks http.ResponseWriter
type ResponseWriter struct {
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{}
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	fmt.Println("Writing Response: ", string(b))
	return 0, nil
}

func (w *ResponseWriter) WriteHeader(header int) {
	fmt.Println("Writing Response Header(): ", header)
}

func (w *ResponseWriter) Header() http.Header {
	fmt.Println("Getting Response Header: ", nil)
	return make(map[string][]string)
}
