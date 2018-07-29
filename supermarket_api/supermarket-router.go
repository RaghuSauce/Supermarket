package supermarket_api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
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

//TODO
func JsonFileLogger(out SuperMarketLog) {
	output, _ := json.Marshal(out)
	fmt.Println(string(output))
}
func StandardOutLogger(out SuperMarketLog) {
	outputFormat := "%s\t%s\t%s\t%s"
	log.Printf(
		outputFormat,
		out.Method,
		out.RequestURI,
		out.Name,
		out.Time,
	)

}

//output ,_:= json.Marshal(outputStruct)
//fmt.Printf("%s\n",outputStruct.Time)

type SuperMarketLog struct {
	Method     string        `JSON:"method"`
	RequestURI string        `JSON:"request-uri"`
	Name       string        `JSON:"path"`
	Time       time.Duration `JSON:"time"`
}
