// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package files

import (
	"os"
	"regexp"
)

func Expand(path string) string {
	re := regexp.MustCompile(`\%(?i)(\w*)\%`)
	match := re.FindStringSubmatch(path)
	if match != nil {
		newValue := os.Getenv(match[1])
		if newValue != "" {
			path = re.ReplaceAllString(path, newValue)
		}

	}

	// This is sometimes used in registry keys.
	re = regexp.MustCompile(`^\\(?i)SystemRoot`)
	path = re.ReplaceAllString(path, os.Getenv("SystemRoot"))

	return path
}
