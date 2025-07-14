# Circuit Breaker

This is a Golang project that demonstrates the [circuit breaker](https://learn.microsoft.com/en-us/previous-versions/msp-n-p/dn589784(v=pandp.10)) pattern. This project uses the [gobreaker](https://github.com/sony/gobreaker) library. `main.go` is the only file simulating client-server communication mediated by circuit breaker. The `main()` function configures the circuit breaker and calls the `client()` function. The `client()` function makes requests to the `server()` function mediated by a circuit breaker. The `server()` function returns either a success or a failure response from a predefined list of responses.

State transations of circuit breaker
- close -> open
- open -> half open
- half open -> close / open (depending on the server response)
