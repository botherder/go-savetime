// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package hashes

import (
	"os"
	"io"
	"errors"
	"hash"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

const (
	HASH_MD5 = "md5"
	HASH_SHA1 = "sha1"
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

func HashFile(target string, algorithm string) (string, error) {
	return hashTarget(target, algorithm, "file")
}

func FileMD5(filePath string) (string, error) {
	return HashFile(filePath, HASH_MD5)
}

func FileSHA1(filePath string) (string, error) {
	return HashFile(filePath, HASH_SHA1)
}

func FileSHA256(filePath string) (string, error) {
	return HashFile(filePath, HASH_SHA256)
}

func FileSHA512(filePath string) (string, error) {
	return HashFile(filePath, HASH_SHA512)
}

func HashString(target string, algorithm string) (string, error) {
	return hashTarget(target, algorithm, "string")
}

func StringMD5(target string) (string, error) {
	return HashString(target, HASH_MD5)
}

func StringSHA1(target string) (string, error) {
	return HashString(target, HASH_SHA1)
}

func StringSHA256(target string) (string, error) {
	return HashString(target, HASH_SHA256)
}

func StringSHA512(target string) (string, error) {
	return HashString(target, HASH_SHA512)
}
