package supermarket_api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

//Struct used to hold route information that will be logged out
type SuperMarketLog struct {
	Method     string        `JSON:"method"`
	RequestURI string        `JSON:"request-uri"`
	Name       string        `JSON:"path"`
	Time       time.Duration `JSON:"time"`
}

//Function to handle logging for the Supermarket application
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
		JsonFileLogger(output)
		StandardOutLogger(output)
	})
}

//Logs all accessed routes to a file in json format
func JsonFileLogger(out SuperMarketLog) error {
	output, _ := json.Marshal(out)        // Create he output to log
	stringOutput := string(output) + "\n" // Append a newline to the output
	//If the file doesn't exist, create it or append to the file
	f, err := os.OpenFile("rest.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(stringOutput)); err != nil { //Write out to the log
		log.Fatal(err)
	}
	if err := f.Close(); err != nil { //Close the writer
		log.Fatal(err)
	}
	return err
}

//Logs all accessed routes to sout in a tab separated line
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
