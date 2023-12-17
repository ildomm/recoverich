package recoverich

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
)

// RecoverWithTrackedValues recovers from a panic and logs the error and stack trace.
// It also logs the tracked values.
//
// The tracked values are any values that you want to log when a panic occurs.
func RecoverWithTrackedValues(values ...any) {
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

		valuesMap := make(map[string]interface{})
		for index, trackedRecord := range values {
			key := fmt.Sprintf("%s_%d", reflect.TypeOf(trackedRecord).String(), index)
			valuesMap[key] = fmt.Sprintf("%v", trackedRecord)
		}
		log.Printf("ERROR: Tracked values dump ***\n%v\n*** end\n", valuesMap)
	}
}
