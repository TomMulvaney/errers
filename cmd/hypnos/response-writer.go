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
	fmt.Println("ResponseWriter.Write(): ", string(b))
	return 0, nil
}

func (w *ResponseWriter) WriteHeader(header int) {
	fmt.Println("ResponseWriter.WriteHeader(): ", header)
}

func (w *ResponseWriter) Header() http.Header {
	fmt.Println("ResponseWriter.Header()")
	return make(map[string][]string)
}
