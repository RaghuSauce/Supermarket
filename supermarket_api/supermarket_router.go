package supermarket_api

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Supermarket_router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = SupermarketLogger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func SupermarketLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		output := SuperMarketLog{
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		}
		//JsonFileLogger(output)
		StandardOutLogger(output)
	})
}
