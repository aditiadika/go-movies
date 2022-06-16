package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port listen on")
	flag.StringVar(&cfg.env, "env", "development", "application Environment (development|production)")
	flag.Parse()

	http.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}

		js, err := json.MarshalIndent(currentStatus, "", "\t")

		if err != nil {
			log.Println(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(js)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)

	if err != nil {
		log.Println(err)
	}
}
