package recoverich

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	// Create an instance of the MockLogger
	mockLogger := &MockLogger{}

	// Replace the default logger with the mock logger
	log.SetOutput(mockLogger)

	// Panic
	func() {
		defer Recover()

		panic("test panic")
	}()

	// Assert the log messages
	assert.Equal(t, 2, len(mockLogger.logs))
	assert.Contains(t, mockLogger.String(), "ERROR: test panic")
	assert.Contains(t, mockLogger.String(), "ERROR: Stacktrace dump ***")

	// Restore the default logger
	log.SetOutput(os.Stdout)
}
