# Circuit Breaker Demonstration
## Project Explanation

This project demonstrates the [circuit breaker](https://learn.microsoft.com/en-us/previous-versions/msp-n-p/dn589784(v=pandp.10)) pattern using the
[gobreaker](https://github.com/sony/gobreaker) library. It simulates the interaction between a client and server mediated by a circuit breaker.

### Key Components:
- **server()**: Simulated with predefined `200 OK` (success) and `500 Internal Server Error` (failure) responses.
- **client()**: Calls the server using a circuit breaker, logging successes and failures.
- **state transitions**: The canned server responses make following state transations
  - close -> open
  - open -> half open
  - half open -> close
  - close -> open
  - open -> half open
  - half open -> open

### Circuit Breaker Configuration:
- **ReadyToTrip**: Circuit opens after 2 consecutive failures.
- **MaxRequests**: 1 request allowed in half-open state.
- **Timeout**: Circuit stays open for 2 seconds before testing again.
