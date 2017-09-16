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

// LowLevelHTTPClientGET ...
func LowLevelHTTPClientGET() error {
	err := errors.Unavailable(baseErrors.New("Global Insomnia, nobody can sleep")) // Imagine that we just failed parsing JSON

	// TODO: Move this to func in nskeleton/errors called ConvertHTTPClient(err error) error
	if errors.IsNError(err) {
		e := err.(errors.NError)

		switch e.Status() {
		case errors.StatusUnavailable:
			err = errors.UpstreamUnavailable(err)
		}
	}

	return err
}

// ReadDream ...
func ReadDream(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := LowLevelHTTPClientGET()

	return errors.Wrap(err, "Failed getting from dream cache")
}

func getPath(url *url.URL) string {
	return url.String()[len(urlBase)+2:]
}

// APIMiddleware ...
func APIMiddleware(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background() // TODO

	if err := ReadDream(ctx, w, r); err != nil {
		var opts []errors.Option

		opts = append(opts, errors.ToHTTPStatus())

		logFields := log.Fields{
			"Path":   getPath(r.URL), // TODO
			"Method": r.Method,
		}

		opts = append(opts, errors.LogError(logFields))

		opts = append(opts, errors.WriteHeader(w, http.StatusInternalServerError)) // Default InternalServerError

		errors.HandleError(err, opts...)
	}
}
