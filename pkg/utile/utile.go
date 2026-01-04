package utile

import (
	"os"
	"time"
)

func StrPtr(s string) *string {
	return &s
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func TimePtr(time time.Time) *time.Time {
	return &time
}
