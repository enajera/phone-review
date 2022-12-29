package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()

	//globals middlewares
	mux.Use(
         middleware.Logger, //log every http request
		 middleware.Recoverer, //recover if a panic occurs
	)
    
	mux.Post("/smartphones",nil)
	mux.Get("/hello",helloHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "elvin")

	res := map[string]interface{}{"message":"hello world"}

	_ = json.NewEncoder(w).Encode(res)
}