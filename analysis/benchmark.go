package analysis

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/wander4747/go-benchmark/httpclient"
	"github.com/wander4747/go-benchmark/utils"
)

func RunBenchmark(url string, method string, requests int, duration time.Duration, payload string) error {
	numWorkers := runtime.NumCPU()

	startTime := time.Now()

	var wg sync.WaitGroup
	var mu sync.Mutex

	completedRequests := 0
	successfulRequests := 0
	failedRequests := 0
	var totalResponseTime time.Duration

	infiniteRequests := requests <= 0

	requestCh := make(chan int)

	if infiniteRequests {
		go func() {
			for {
				requestCh <- 1
				if time.Since(startTime) > duration {
					close(requestCh)
					return
				}
			}
		}()
	} else {
		go func() {
			for i := 0; i < requests; i++ {
				requestCh <- i
			}
			close(requestCh)
		}()
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range requestCh {
				requestStartTime := time.Now()
				var err error

				if method == "POST" {
					_, err = httpclient.SendPostRequest(url, payload)
				} else {
					_, err = httpclient.SendGetRequest(url)
				}

				responseDuration := time.Since(requestStartTime)

				mu.Lock()
				completedRequests++
				totalResponseTime += responseDuration
				if err != nil {
					log.Printf("request error: %v", err)
					failedRequests++
				} else {
					successfulRequests++
				}
				mu.Unlock()

				if !infiniteRequests && time.Since(startTime) > duration {
					return
				}
			}
		}()
	}

	wg.Wait()

	var avgResponseTime time.Duration
	if successfulRequests > 0 {
		avgResponseTime = totalResponseTime / time.Duration(successfulRequests)
	} else {
		avgResponseTime = 0
	}

	rps := float64(successfulRequests) / time.Since(startTime).Seconds()

	if requests <= 0 {
		requests = successfulRequests + failedRequests
	}

	utils.PrintStats(requests, successfulRequests, failedRequests, avgResponseTime, rps, time.Since(startTime))
	return nil
}
