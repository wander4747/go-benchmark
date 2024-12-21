package main

import (
	"flag"
	"github.com/wander4747/go-benchmark/analysis"
	"log"
	"os"
	"time"
)

func main() {
	url := flag.String("url", "https://test-api.k6.io/public/crocodiles/?format=api", "API URL for the benchmark")
	method := flag.String("method", "GET", "HTTP method for the requests (GET or POST)")
	requests := flag.Int("requests", 0, "Total number of requests to be performed")
	duration := flag.String("duration", "30s", "Duration of the benchmark (e.g., 30s, 1m, 2h)")
	payload := flag.String("payload", `{"key":"value"}`, "Payload for POST requests")

	flag.Parse()

	if *url == "" {
		log.Fatal("the URL must be provided!")
		os.Exit(1)
	}

	durationParsed, err := time.ParseDuration(*duration)
	if err != nil {
		log.Fatalf("error parsing duration: %v", err)
	}

	err = analysis.RunBenchmark(*url, *method, *requests, durationParsed, *payload)
	if err != nil {
		log.Fatalf("error running the analysis: %v", err)
	}
}
