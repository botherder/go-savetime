// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package files

import (
	"io"
	"os"
)

// Copy a file from the source to the destination paths.
func Copy(srcPath, dstPath string) error {
	srcHandle, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcHandle.Close()

	dstHandle, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstHandle.Close()

	if _, err = io.Copy(dstHandle, srcHandle); err != nil {
		return err
	}

	err = dstHandle.Sync()
	return err
}
