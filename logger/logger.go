package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const logFile = "log.txt"

var (
	initialized bool
	mu          sync.Mutex
	file        *os.File
)

func checkDebug() bool {
	_, exists := os.LookupEnv("DEBUG")
	return exists
}

func dateString() string {
	return time.Now().Format("[02-01-2006_15:04]")
}

func Init() error {
	mu.Lock()
	defer mu.Unlock()

	if !checkDebug() || initialized {
		return nil
	}

	var err error

	file, err = os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening log file: %w", err)
	}

	if _, err := file.WriteString("========================================\n==== DEBUG =============================\n"); err != nil {
		return fmt.Errorf("error writing to log file: %w", err)
	}

	initialized = true
	return nil
}

func Log(msg string) {
	mu.Lock()
	defer mu.Unlock()

	if !checkDebug() || !initialized || file == nil {
		return
	}

	logLine := fmt.Sprintf("%s %s\n", dateString(), msg)
	if _, err := file.WriteString(logLine); err != nil {
		fmt.Println("error writing to log file")
	}
}

func Close() error {
	mu.Lock()
	defer mu.Unlock()

	if file != nil {
		err := file.Close()
		file = nil
		return err
	}

	return nil
}
