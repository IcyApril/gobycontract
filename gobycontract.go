// Go By Contract provides a lightweight mechanism to validate Design-by-Contract pre-conditions and post-conditions in
// production environments by using Sentry error reporting instead of "fail hard" behaviour.
//
// Setting the GOBYCONTRACT_DONTPANIC environment variable disables panics for a production environment when set to
// "true", whilst the SENTRY_DSN environment will independently enable reporting to Sentry.
//
// The following is an example of using Require for preconditions and Ensure for postconditions:
// 		func SecondsToSecondsAndMinutes(seconds int) (minutes int, remainingSeconds int) {
// 			gobycontract.Require(seconds >= 0, "Input seconds must be positive")
//
// 			minutes = seconds/60
// 			remainingSeconds = seconds % 60
//
// 			gobycontract.Ensure(minutes > 0, "Output minutes most be positive")
// 			gobycontract.Ensure(remainingSeconds > 0, "Output remaining seconds most be positive")
// 			gobycontract.Ensure(remainingSeconds < 59, "There can be no more than 59 remaining seconds")
//
// 			return
// 		}
package gobycontract

import (
	"os"
	"strings"
	"github.com/getsentry/raven-go"
)


func Require(pass bool, description string) {
	if pass == true {
		return
	}

	message := "Pre-Condition not met: " + description

	if shouldPanic() == true {
		panic(message)
	}

	logToSentry(message, "pre-condition")
}

func Ensure(pass bool, description string) {
	if pass == true {
		return
	}

	message := "Post-Condition not met: " + description

	if shouldPanic() == true {
		panic(message)
	}

	logToSentry(message, "post-condition")
}

func shouldPanic() (shouldPanic bool) {
	gbcShouldPanic := os.Getenv("GOBYCONTRACT_DONTPANIC")
	gbcShouldPanic = strings.ToLower(gbcShouldPanic)

	shouldPanic = true
	if gbcShouldPanic == "true" {
		shouldPanic = false
	}

	return
}

func logToSentry(message string, category string) {
	shouldLog := len(os.Getenv("SENTRY_DSN")) > 0

	if shouldLog == false {
		return
	}

	raven.CaptureMessageAndWait(message, map[string]string{"category": "contract"+category})

}