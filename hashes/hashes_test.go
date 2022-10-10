package hashes

import (
	"testing"
)

var (
	clearString  = "test"
	stringMD5    = "098f6bcd4621d373cade4e832627b4f6"
	stringSHA1   = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
	stringSHA256 = "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	stringSHA512 = "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
)

func TestStringMD5(t *testing.T) {
	md5, err := StringMD5(clearString)
	if err != nil {
		t.Errorf("TestStringMD5 failed with error: %s", err.Error())
	}
	if md5 != stringMD5 {
		t.Errorf("TestStringMD5 failed with wrong hash: %s", md5)
	}
}

func TestStringSHA1(t *testing.T) {
	sha1, err := StringSHA1(clearString)
	if err != nil {
		t.Errorf("TestStringSHA1 failed with error: %s", err.Error())
	}
	if sha1 != stringSHA1 {
		t.Errorf("TestStringSHA1 failed with wrong hash: %s", sha1)
	}
}

func TestStringSHA256(t *testing.T) {
	sha256, err := StringSHA256(clearString)
	if err != nil {
		t.Errorf("TestStringSHA256 failed with error: %s", err.Error())
	}
	if sha256 != stringSHA256 {
		t.Errorf("TestStringSHA256 failed with wrong hash: %s", sha256)
	}
}

func TestStringSHA512(t *testing.T) {
	sha512, err := StringSHA512(clearString)
	if err != nil {
		t.Errorf("TestStringSHA512 failed with error: %s", err.Error())
	}
	if sha512 != stringSHA512 {
		t.Errorf("TestStringSHA512 failed with wrong hash: %s", sha512)
	}
}
