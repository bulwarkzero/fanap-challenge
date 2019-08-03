package utils

import (
	"os"
	"path/filepath"
)

// GetCurrentPath returns absolute path for current running binary
func GetCurrentPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}
