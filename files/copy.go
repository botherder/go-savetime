package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyDir recursively copies a directory from source to destination.
func CopyDir(srcDir, dstDir string) error {
	err := os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		srcPath := filepath.Join(srcDir, file.Name())
		dstPath := filepath.Join(dstDir, file.Name())

		if file.IsDir() {
			err := CopyDir(srcPath, dstPath)
			if err != nil {
				fmt.Println("Error copying directory:", err)
			}
		} else {
			err := CopyFile(srcPath, dstPath)
			if err != nil {
				fmt.Println("Error copying file:", err)
			}
		}
	}

	return nil
}

// CopyFile copies a file from the source to the destination paths.
func CopyFile(srcPath, dstPath string) error {
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

// Copy is kept for retro-compatibility.
func Copy(srcPath, dstPath string) error {
	return CopyFile(srcPath, dstPath)
}
