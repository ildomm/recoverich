package recoverich

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

const stackTraceMaxSize = 64 << 10

// Recover recovers from a panic and logs the error and stack trace.
// This is the most basic version of Recover.
func Recover() {
	if err := recover(); err != nil {

		stackTrace := make([]byte, stackTraceMaxSize)
		stackTrace = stackTrace[:runtime.Stack(stackTrace, false)]

		switch x := err.(type) {
		case string:
			err = errors.New(x)
		default:
			err = fmt.Errorf("unknown panic: %w", x.(error))
		}

		log.Printf("ERROR: %v", err)
		log.Printf("ERROR: Stacktrace dump ***\n%s\n*** end\n", stackTrace)
	}
}
