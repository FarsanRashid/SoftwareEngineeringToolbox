package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

// Simulate a server with predefined canned responses
var responses = []int{
	http.StatusOK,                  // 200: success, closed -> closed
	http.StatusInternalServerError, // 500: failure, closed -> closed
	http.StatusInternalServerError, // 500: failure, closed -> open
	/*
		open state. request will not reach
	*/
	http.StatusOK, // 200:success, half-open -> closed

	http.StatusInternalServerError, // 500: failure, closed -> closed
	http.StatusInternalServerError, // 500: failure, closed -> open
	/*
		open state. request will not reach
	*/
	http.StatusInternalServerError, // 500: failure, half-open -> open
	/*
		open state. request will not reach
	*/
}

var responseIndex = 0

// Server function simulating various responses
func server(call_id int) (int, error) {
	fmt.Printf("Received cliend call %d\n", call_id)
	if responseIndex >= len(responses) {
		responseIndex = 0 // loop back to simulate repetitive state changes
	}
	resp := responses[responseIndex]
	responseIndex++

	if resp == http.StatusOK {
		return resp, nil
	}
	return resp, fmt.Errorf("server error: %d", resp)
}

// Client function using circuit breaker to call server
func client(cb *gobreaker.CircuitBreaker) {
	for i := 0; i < 15; i++ {
		fmt.Printf("Making client call %d (breaker state: %s)\n", i+1, cb.State().String())

		result, err := cb.Execute(func() (interface{}, error) {
			status, err := server(i + 1)
			return status, err
		})

		if err != nil {
			fmt.Printf("Failed: %v (breaker state: %s)\n", err, cb.State().String())
		} else {
			fmt.Printf("Success: %v (breaker state: %s)\n", result, cb.State().String())
		}

		time.Sleep(500 * time.Millisecond) // Delay between calls
	}
}

func main() {
	// Initialize the circuit breaker with a configuration
	cbConfig := gobreaker.Settings{
		Name:        "ServerBreaker",
		MaxRequests: 1,               // Allow 1 request in half-open state
		Interval:    0,               // No time-based reset
		Timeout:     2 * time.Second, // Time in open state
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 1 // Open on 2 consecutive failures
		},
	}

	cb := gobreaker.NewCircuitBreaker(cbConfig)

	// Call the client which uses the circuit breaker
	client(cb)
}
