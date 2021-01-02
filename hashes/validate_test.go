// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package hashes

import (
	"testing"
)

func TestValidateMD5(t *testing.T) {
	if ValidateMD5(clearString) {
		t.Errorf("ValidateMD5 failed test 1: expected false, got true")
	}

	if !ValidateMD5(stringMD5) {
		t.Errorf("ValidateMD5 failed test 2: expected true, got false")
	}
}

func TestValidateSHA1(t *testing.T) {
	if ValidateSHA1(clearString) {
		t.Errorf("ValidateSHA1 failed test 1: expected false, got true")
	}

	if !ValidateSHA1(stringSHA1) {
		t.Errorf("ValidateSHA1 failed test 2: expected true, got false")
	}
}

func TestValidateSHA256(t *testing.T) {
	if ValidateSHA256(clearString) {
		t.Errorf("ValidateSHA256 failed test 1: expected false, got true")
	}

	if !ValidateSHA256(stringSHA256) {
		t.Errorf("ValidateSHA256 failed test 2: expected true, got false")
	}
}

func TestValidateSHA512(t *testing.T) {
	if ValidateSHA512(clearString) {
		t.Errorf("ValidateSHA512 failed test 1: expected false, got true")
	}

	if !ValidateSHA512(stringSHA512) {
		t.Errorf("ValidateSHA512 failed test 2: expected true, got false")
	}
}
