package main

import (
	"context"
	"net/http"
	"net/url"

	baseErrors "errors"

	"github.com/nskeleton/errors"

	log "github.com/sirupsen/logrus"
)

const (
	urlBase = "http://hypnos.com/api/v1"
)

// DreamDBClientGET ...
func DreamDBClientGET() error {
	return errors.Unavailable(baseErrors.New("Global Insomnia")) // Imagine that we just failed parsing JSON
}

// ReadDream ...
func ReadDream(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := DreamDBClientGET()

	return errors.Wrap(err, "Failed getting from dream database")
}

func getPath(url *url.URL) string {
	return url.String()[len(urlBase)+2:]
}

// APIMiddleware ...
func APIMiddleware(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background() // TODO

	if err := ReadDream(ctx, w, r); err != nil {

		var doers []errors.Doer

		doers = append(doers, errors.Upstream)

		doers = append(doers, errors.ToHTTPStatus)

		logFields := log.Fields{
			"Path":   getPath(r.URL),
			"Method": r.Method,
		}

		doers = append(doers, errors.LogError(logFields))

		doers = append(doers, errors.WriteHTTPHeader(w))

		errors.HandleError(err, doers...)
	}
}
