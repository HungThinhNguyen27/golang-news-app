package utils

import "time"

// GetCurrentTimestamp returns the current time formatted as YYYY-MM-DD HH:MM:SS
func GetCurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
