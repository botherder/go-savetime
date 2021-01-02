// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package text

import (
	"testing"
)

var text string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi
ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit
in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur
sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt
mollit anim id est laborum.`

func TestContains(t *testing.T) {
	if Contains(text, "This string should not be contained") {
		t.Errorf("Contains failed case-sensitive test 1: expected false, got true")
	}

	if !Contains(text, "Lorem ipsum") {
		t.Errorf("Contains failed case-sensitive test 2: expected true, got false")
	}

	if Contains(text, "LOREM IPSUM") {
		t.Errorf("Contains failed case-sensitive test 3: expected false, got true")
	}
}

func TestContainsNoCase(t *testing.T) {
	if ContainsNoCase(text, "This string should not be contained") {
		t.Errorf("ContainsNoCase failed case-insensitive test 1: expected false, got true")
	}

	if !ContainsNoCase(text, "LOREM IPSUM") {
		t.Errorf("ContainsNoCase failed case-insensitive test 2: expected true, got false")
	}

	if !ContainsNoCase(text, "Lorem ipsum") {
		t.Errorf("ContainsNoCase failed case-insensitive test 3: expected true, got false")
	}
}
