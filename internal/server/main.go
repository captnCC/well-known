package server

import (
	"fmt"
	"github.com/captnCC/well-known/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func Serve(config config.Config, port int, logger bool) {
	r := chi.NewRouter()

	if logger {
		r.Use(middleware.Logger)
	}

	for path, entry := range config.Files {
		r.Get("/.well-known/"+path, func(w http.ResponseWriter, r *http.Request) {

			for header, value := range entry.Headers {
				w.Header().Add(header, value)
			}

			w.WriteHeader(200)
			_, err := w.Write([]byte(entry.Content))
			if err != nil {
				return
			}
		})
	}

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on %s", addr)
	err := http.ListenAndServe(addr, r)

	if err != nil {
		log.Fatal(err)
	}
}
