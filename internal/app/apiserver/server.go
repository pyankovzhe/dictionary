package apiserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type server struct {
	*http.Server
	logger *logrus.Logger
	// store
}

func newServer(logger *logrus.Logger, serverAddr string) *server {
	s := &server{
		Server: &http.Server{
			Addr: serverAddr,
		},
		logger: logger,
	}

	s.logger.Info("Server initialized")
	return s
}
