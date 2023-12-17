package recoverich

import (
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestRecoverWithContextValues(t *testing.T) {
	// Create an instance of the MockLogger
	mockLogger := &MockLogger{}

	// Replace the default logger with the mock logger
	log.SetOutput(mockLogger)

	// Setup object instance
	type Car struct {
		Color string
	}
	truck := &Car{Color: "red"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = context.WithValue(ctx, "an string", "an value")
	ctx = context.WithValue(ctx, "an struct", *truck)

	// Panic
	func() {
		defer RecoverWithContextValues(ctx)

		panic("test panic")
	}()

	// Assert the log messages
	assert.Equal(t, 3, len(mockLogger.logs))
	assert.Contains(t, mockLogger.String(), "ERROR: test panic")
	assert.Contains(t, mockLogger.String(), "ERROR: Stacktrace dump ***")
	assert.Contains(t, mockLogger.String(), "ERROR: Context values dump ***")

	// First value
	assert.Contains(t, mockLogger.String(), "Key: an string")
	assert.Contains(t, mockLogger.String(), "Type: string")
	assert.Contains(t, mockLogger.String(), "Value: an value")

	// Second value
	assert.Contains(t, mockLogger.String(), "Key: an struct")
	assert.Contains(t, mockLogger.String(), "Type: recoverich.Car")
	assert.Contains(t, mockLogger.String(), "Value: {red}")

	// Restore the default logger
	log.SetOutput(os.Stdout)
}
