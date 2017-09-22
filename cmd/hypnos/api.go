package main

import (
	"context"
	"net/http"
	"net/url"

	"google.golang.org/grpc/codes"

	baseErrors "errors"

	"github.com/nskeleton/errors"
	"github.com/nskeleton/errors/grpc"

	log "github.com/sirupsen/logrus"
)

const (
	urlBase = "http://hypnos.com/api/v1"
)

// GRPCClient ...
func GRPCClient() error {
	err := errors.WrapStatus(baseErrors.New("Global Insomnia"), int(codes.Unavailable))
	return grpc.FromGRPCStatus(err)
}

// ReadDream ...
func ReadDream(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := GRPCClient()
	return errors.BadReq(err, "Failed getting from dream database")
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

		doers = append(doers, errors.StatusMessage)

		logFields := log.Fields{
			"Path":   getPath(r.URL),
			"Method": r.Method,
		}

		doers = append(doers, errors.LogError(logFields))

		doers = append(doers, errors.WriteHTTPHeader(w))

		errors.HandleError(err, doers...)
	}
}
