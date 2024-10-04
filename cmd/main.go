package main

import (
	"github.com/captnCC/well-known/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	err, endpointConfig := config.ReadConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	for path, entry := range endpointConfig.Files {
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

	err = http.ListenAndServe(":3000", r)

	if err != nil {
		return
	}
}
