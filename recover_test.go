package recoverich

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

// MockLogger is a simple implementation of the Logger interface for testing
type MockLogger struct {
	logs []string
}

func (ml *MockLogger) Write(p []byte) (n int, err error) {
	ml.logs = append(ml.logs, string(p))
	return len(p), nil
}

// String returns all the log entries as one big string
func (ml MockLogger) String() string {
	return strings.Join(ml.logs, "\n")
}

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
	assert.Contains(t, mockLogger.String(), "test panic")
	assert.Contains(t, mockLogger.String(), "Stacktrace ***")

	// Restore the default logger
	log.SetOutput(os.Stdout)
}
