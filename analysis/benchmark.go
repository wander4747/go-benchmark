package analysis

import (
	"github.com/wander4747/go-benchmark/httpclient"
	"github.com/wander4747/go-benchmark/utils"
	"log"
	"time"
)

func RunBenchmark(url string, method string, requests int, duration time.Duration, payload string) error {
	startTime := time.Now()
	completedRequests := 0
	successfulRequests := 0
	failedRequests := 0
	var totalResponseTime time.Duration

	for completedRequests < requests {
		var err error
		var responseDuration time.Duration

		requestStartTime := time.Now()

		if method == "POST" {
			_, err = httpclient.SendPostRequest(url, payload)
		} else {
			_, err = httpclient.SendGetRequest(url)
		}

		responseDuration = time.Since(requestStartTime)
		totalResponseTime += responseDuration

		if err != nil {
			log.Printf("Erro na requisição: %v", err)
			failedRequests++
		} else {
			successfulRequests++
		}

		completedRequests++

		if time.Since(startTime) > duration {
			break
		}
	}

	avgResponseTime := totalResponseTime / time.Duration(successfulRequests)
	rps := float64(successfulRequests) / time.Since(startTime).Seconds()

	utils.PrintStats(completedRequests, successfulRequests, failedRequests, avgResponseTime, rps, time.Since(startTime))

	return nil
}
