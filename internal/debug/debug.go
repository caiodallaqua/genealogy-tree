// Helper package to handle prints accross the application
// to reduce boilerplate while keeping contextual information
// such as function names for errors.

// KISS: Keep it simple, stupid (a reminder for myself).
// These functions can be called many times during the application runtime
// so avoid fancy features that may cause performance overhead.
package debug

import (
	"log"
	"os"
)

var (
	codeEnv      string
	codeEnvIsSet bool
)

func init() {
	codeEnv, codeEnvIsSet = os.LookupEnv("CODE_ENV")

	// Default behavior
	if !codeEnvIsSet {
		codeEnv = "dev"
	}
}

// Print context with message and error
// "{datetime} | {function name} | {message}: {error}"
func ShowErr(caller, msg string, err error) {
	if caller == "" || err == nil {
		return
	}

	if codeEnv == "dev" {
		log.Printf("| %s | %s: %v", caller, msg, err)
	}
}
