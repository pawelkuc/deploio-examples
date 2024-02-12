package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	responseEnv   = "RESPONSE_TEXT"
	delayResponse = "DELAY_RESPONSE"
	defaultDelay  = 2000 * time.Microsecond
)

func main() {
	delay := defaultDelay

	if val, ok := os.LookupEnv(delayResponse); ok {
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Fatalf("error parsing delay value %q to int: %s", val, err)
		}
		delay = d
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("got request\n")
		io.WriteString(w, os.Getenv(responseEnv)+"\n")
	})

	http.HandleFunc("/delayedresponse", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("got delayed response request (%v)\n", delay)
		var result = `{"data":[]}`
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(result))
	})

	if err := http.ListenAndServe(":1337", nil); err != nil {
		log.Fatal(err)
	}
}
