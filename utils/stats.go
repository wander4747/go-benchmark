package utils

import (
	"fmt"
	"time"
)

func PrintStats(totalRequests, successfulRequests, failedRequests int, avgResponseTime time.Duration, rps float64, totalDuration time.Duration) {
	fmt.Printf("\nBenchmark completed in: %v\n", totalDuration)
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Successful requests: %d\n", successfulRequests)
	fmt.Printf("Failed requests: %d\n", failedRequests)
	fmt.Printf("Average response time: %.2f ms\n", avgResponseTime.Seconds()*1000) // Convert to milliseconds
	fmt.Printf("Requests per second: %.2f rps\n", rps)
}
