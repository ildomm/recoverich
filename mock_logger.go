package recoverich

import "strings"

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
