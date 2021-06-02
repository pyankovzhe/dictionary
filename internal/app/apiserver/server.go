package apiserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pyankovzhe/dictionary/internal/app/store"
	"github.com/sirupsen/logrus"
)

type server struct {
	*http.Server
	logger *logrus.Logger
	store  store.Store
}

func newServer(logger *logrus.Logger, store store.Store, serverAddr string) *server {
	s := &server{
		Server: &http.Server{
			Addr: serverAddr,
		},
		logger: logger,
		store:  store,
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

	router.Mount("/api", s.apiRouter())
}

func (s *server) apiRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/card", s.CreateCard)

	return r
}
