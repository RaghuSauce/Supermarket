package supermarket_api

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

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
