// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package runtime

import (
	"os"
	"path"
)

// GetExecutableDirectory returns the path to the folder containing the program
// we are executing.
func GetExecutableDirectory() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}

	return path.Dir(exe)
}
