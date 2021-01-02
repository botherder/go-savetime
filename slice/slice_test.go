// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package slices

import (
	"testing"
)

var slice []string = []string{
	"qwe",
	"rty",
	"uiop",
}

func TestContains(t *testing.T) {
	if Contains(slice, "asd") {
		t.Errorf("Contains failed case-sensitive test 1: expected false, got true")
	}

	if !Contains(slice, "qwe") {
		t.Errorf("Contains failed case-sensitive test 2: expected true, got false")
	}

	if Contains(slice, "QWE") {
		t.Errorf("Contains failed case-sensitive test 3: expected false, got true")
	}
}

func TestContainsNoCase(t *testing.T) {
	if ContainsNoCase(slice, "asd") {
		t.Errorf("ContainsNoCase failed case-insensitive test 1: expected false, got true")
	}

	if !ContainsNoCase(slice, "qwe") {
		t.Errorf("ContainsNoCase failed case-insensitive test 2: expected true, got false")
	}

	if !ContainsNoCase(slice, "QWE") {
		t.Errorf("ContainsNoCase failed case-insensitive test 3: expected true, got false")
	}
}
