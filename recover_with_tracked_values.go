package recoverich

import (
	"fmt"
	"log"
	"reflect"
)

// RecoverWithTrackedValues recovers from a panic and logs the error and stack trace.
// It also logs the tracked values.
//
// The tracked values are any values that you want to log when a panic occurs.
func RecoverWithTrackedValues(values ...any) {
	if err := recover(); err != nil {

		print(err)

		// Pile the values formatted
		log.Print("Tracked values ***\n")
		printPile(gatherTrackedValues(values))
		log.Print("***")
	}
}

// gatherTrackedValues iterates over the context and extracts the tracked values
func gatherTrackedValues(values ...interface{}) []string {
	v := make([]string, 0)

	for i, t := range values[0].([]interface{}) {
		f := fmt.Sprintf("Name: %v, Type: %v, Value: %v", i, reflect.TypeOf(t).String(), t)
		v = append(v, f)
	}
	return v
}
