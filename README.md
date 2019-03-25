# Go By Contract

This is a simple implementation of production-safe Design by Contract in Go, using pre-conditions and post-conditions.

Whilst this implementation allows for errors to "fail hard" in development and staging environments, it also can be
configured in a report-only mode for production environments.

This is an example of the following pattern: [A Pattern for Validating Design by Contract Assertions in Production (with Go and Sentry)](https://icyapril.com/go/programming/2019/03/25/a-pattern-for-validating-design-by-contract-in-Production-with-go-and-sentry.html)

## Environment Variables

* `GOBYCONTRACT_DONTPANIC` when set to "true" will stop panicking when the ````Require```` or `````Ensure```` checks fail
* `SENTRY_DSN` is required to be set for the Sentry reporting of contract violations to be enabled

Per the [Sentry Raven-Go](https://docs.sentry.io/clients/go/) documentation, the ````SENTRY_RELEASE```` and
````SENTRY_ENVIRONMENT```` environment variables may also be set.

## Example Implementation

* Require is used to determine if a precondition is met
* Ensure is used to determine if a postcondition is met

Example of a function using such methods:

```go
func SecondsToSecondsAndMinutes(seconds int) (minutes int, remainingSeconds int) {
	gobycontract.Require(seconds >= 0, "Input seconds must be positive")

	minutes = seconds/60
	remainingSeconds = seconds % 60

	gobycontract.Ensure(minutes >= 0, "Output minutes must be positive")
	gobycontract.Ensure(remainingSeconds >= 0, "Output remaining seconds must be positive")
	gobycontract.Ensure(remainingSeconds < 60, "There can be no more than 59 remaining seconds")

	return
}
```