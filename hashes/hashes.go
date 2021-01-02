// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package hashes

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
)

const (
	HASH_MD5    = "md5"
	HASH_SHA1   = "sha1"
	HASH_SHA256 = "sha256"
	HASH_SHA512 = "sha512"
)

func hashTarget(target, algorithm, format string) (string, error) {
	var h hash.Hash

	switch algorithm {
	case HASH_MD5:
		h = md5.New()
	case HASH_SHA1:
		h = sha1.New()
	case HASH_SHA256:
		h = sha256.New()
	case HASH_SHA512:
		h = sha512.New()
	default:
		return "", errors.New("No valid hashing algorithm specified")
	}

	switch format {
	case "file":
		file, err := os.Open(target)
		if err != nil {
			return "", err
		}
		defer file.Close()

		_, err = io.Copy(h, file)
		if err != nil {
			return "", err
		}
	case "string":
		h.Write([]byte(target))
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// HashFile returns the hex hash of a file at specified path with the
// specified hashing algorithm.
func HashFile(filePath string, algorithm string) (string, error) {
	return hashTarget(filePath, algorithm, "file")
}

// FileMD5 returns the MD5 hash of a file at the specified path.
func FileMD5(filePath string) (string, error) {
	return HashFile(filePath, HASH_MD5)
}

// FileSHA1 returns the SHA1 hash of a file at the specified path.
func FileSHA1(filePath string) (string, error) {
	return HashFile(filePath, HASH_SHA1)
}

// FileSHA256 returns the SHA256 hash of a file at the specified path.
func FileSHA256(filePath string) (string, error) {
	return HashFile(filePath, HASH_SHA256)
}

// FileSHA512 returns the SHA512 hash of a file at the specified path.
func FileSHA512(filePath string) (string, error) {
	return HashFile(filePath, HASH_SHA512)
}

// HashString returns the hex hash of the specified string.
func HashString(target string, algorithm string) (string, error) {
	return hashTarget(target, algorithm, "string")
}

// StringMD5 returns the MD5 hash of the specified string.
func StringMD5(target string) (string, error) {
	return HashString(target, HASH_MD5)
}

// StringSHA1 returns the SHA1 hash of the specified string.
func StringSHA1(target string) (string, error) {
	return HashString(target, HASH_SHA1)
}

// StringSHA256 returns the SHA256 hash of the specified string.
func StringSHA256(target string) (string, error) {
	return HashString(target, HASH_SHA256)
}

// StringSHA512 returns the SHA512 hash of the specified string.
func StringSHA512(target string) (string, error) {
	return HashString(target, HASH_SHA512)
}
