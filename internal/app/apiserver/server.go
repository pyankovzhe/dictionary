package apiserver

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	r := chi.NewRouter()
	s.configureRouter(r)

	s.Handler = r

	s.logger.Info("Initializing server on ", serverAddr)

	return s
}

func (s *server) configureRouter(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: s.logger, NoColor: false}))
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dictionary service"))
	})

	router.Get("/sleep-debug", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		w.Write([]byte("Ok"))
	})
}
