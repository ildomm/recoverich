package recoverich

import (
	errors "github.com/go-errors/errors"
	"log"
)

// standard wrap the error with a stack trace using go-errors,
// and custom format of the stack trace
func print(err interface{}) {
	e := errors.Wrap(err, 0)
	s := e.ErrorStack()

	log.Printf("Recovered from panic: %v", e)
	log.Printf("Stacktrace ***\n %s \n*** \n", s)
}

// printPile prints the pre-formatted list of values
func printPile(values []string) {
	for _, value := range values {
		log.Printf("%v \n", value)
	}
}

// Recover recovers from a panic and logs the error and stack trace.
// This is the most basic version of Recover.
func Recover() {
	if err := recover(); err != nil {
		print(err)
	}
}
