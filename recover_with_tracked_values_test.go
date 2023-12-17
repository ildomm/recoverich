package recoverich

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestRecoverWithTrackedValues(t *testing.T) {
	// Create an instance of the MockLogger
	mockLogger := &MockLogger{}

	// Replace the default logger with the mock logger
	log.SetOutput(mockLogger)

	// Setup object instance
	type Car struct {
		Color string
	}
	truck := &Car{Color: "red"}

	// Panic
	func() {
		defer RecoverWithTrackedValues(
			"trackedRecord1",
			"trackedRecord2",
			*truck)

		panic("test panic")
	}()

	// Assert the log messages
	assert.Equal(t, 3, len(mockLogger.logs))
	assert.Contains(t, mockLogger.String(), "ERROR: test panic")
	assert.Contains(t, mockLogger.String(), "ERROR: Stacktrace dump ***")
	assert.Contains(t, mockLogger.String(), "ERROR: Tracked values dump ***")
	assert.Contains(t, mockLogger.String(), "trackedRecord1")
	assert.Contains(t, mockLogger.String(), "trackedRecord2")
	assert.Contains(t, mockLogger.String(), "{red}")

	// Restore the default logger
	log.SetOutput(os.Stdout)
}
