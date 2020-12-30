// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package slices

import (
	"strings"
)

func stringInSlice(item string, slice []string, caseSensitive bool) bool {
	if !caseSensitive {
		item = strings.ToLower(item)
	}

	for _, entry := range slice {
		if !caseSensitive {
			entry = strings.ToLower(entry)
		}
		if entry == item {
			return true
		}
	}

	return false
}

func IsStringInSlice(item string, slice []string) bool {
	return stringInSlice(item, slice, true)
}

func IsStringInSliceNoCase(item string, slice []string) bool {
	return stringInSlice(item, slice, false)
}
