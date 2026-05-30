package logger

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type Logger struct {
	Writer *slog.Logger
	file   *os.File
}

// Function accepts serviceName as the name of a .log file
func NewLogWriter(serviceName string) (*Logger, error) {
	logFile := strings.TrimSpace(serviceName) + ".log"
	dir := "logs"
	path := filepath.Join(dir, logFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	handler := slog.NewJSONHandler(file, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return &Logger{Writer: logger, file: file}, nil
}

func (l *Logger) Close() error {
	return l.file.Close()
}
