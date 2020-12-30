// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package hashes

import (
	"regexp"
)

func ValidateMD5(target string) bool {
	rxp := regexp.MustCompile(`[a-fA-F0-9]{32}`)
	return rxp.MatchString(target)
}

func ValidateSHA1(target string) bool {
	rxp := regexp.MustCompile(`[a-fA-F0-9]{40}`)
	return rxp.MatchString(target)
}

func ValidateSHA256(target string) bool {
	rxp := regexp.MustCompile(`[a-fA-F0-9]{64}`)
	return rxp.MatchString(target)
}

func ValidateSHA512(target string) bool {
	rxp := regexp.MustCompile(`[a-fA-F0-9]{128}`)
	return rxp.MatchString(target)
}
