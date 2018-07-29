package supermarket_api

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

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
