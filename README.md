# go-savetime

[![Go Report Card](https://goreportcard.com/badge/github.com/botherder/go-savetime)](https://goreportcard.com/report/github.com/botherder/go-savetime)
[![Go Reference](https://pkg.go.dev/badge/github.com/botherder/go-savetime.svg)](https://pkg.go.dev/github.com/botherder/go-savetime)

A collection of Go libraries to save time from re-writing common functions for operations on files, strings, slices, and more. Refer to the [pkg.go.dev](https://pkg.go.dev/github.com/botherder/go-savetime) page for a complete reference of all available functions.

## Files

```go
package main

import (
	"github.com/botherder/go-savetime/files"
)

func main() {
	var err error

	// Copy a file from a source to a destination path.
	err = files.Copy("/path/to/src", "/path/to/dst")

	// Read a JSON file at path into data.
	var data interface{}
	err = files.ReadJSON("/path/file.json", data)

	// Zip a path into archive at destination.
	err = files.Zip("/path/to/folder", "/path/to/archive.zip")
}
```

## Hashes

```go
package main

import (
	"github.com/botherder/go-savetime/hashes"
)

func main() {
	// Hash a file at specified path.
	md5, err := hashes.FileMD5("/path/to/file")
	sha1, err := hashes.FileSHA1("/path/to/file")
	sha256, err := hashes.FileSHA256("/path/to/file")
	sha512, err := hashes.FileSHA512("/path/to/file")

	// Hash a string.
	md5, err = hashes.StringMD5("target")
	sha1, err = hashes.StringSHA1("target")
	sha256, err = hashes.StringSHA256("target")
	sha512, err = hashes.StringSHA512("target")

	// Validate a hash format.
	var valid bool
	valid = hashes.ValidateMD5(md5)
	valid = hashes.ValidateSHA1(sha1)
	valid = hashes.ValidateSHA256(sha256)
	valid = hashes.ValidateSHA512(sha512)
}
```

## Slices

```go
package main

import (
	"github.com/botherder/go-savetime/slice"
)

func main() {
	var contained bool

	// Check if a slice contains a string (case-sensitive).
	contained = slice.Contains([]string{"qwe", "rty"}, "qwe")
	// Check if a slice contains a string (case-insensitive).
	contained = slice.ContainsNoCase([]string{"qwe", "rty"}, "QWE")
}
```

## Text

```go
package main

import (
	"github.com/botherder/go-savetime/text"
)

func main() {
	var contained bool

	// Check if a text contains a string (case-sensitive).
	contained = text.Contains("This is a TEXT", "TEXT")
	// Check if a text contains a string (case-insensitive).
	contained = text.ContainsNoCase("This is a TEXT", "text")
}
```
