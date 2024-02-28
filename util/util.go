package util

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func ReadBody(body io.ReadCloser) ([]byte, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		logrus.WithError(err).Error("Error while reading body")
		return nil, err
	}
	return b, nil
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
}
