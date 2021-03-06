// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package files

import (
	"fmt"
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
)

// WatchFolder uses inotify to monitor a folder and call a callback function
// when changes are detected.
func WatchFolder(folder string, callback func()) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return fmt.Errorf("The specified folder does not exist at path %s", folder)
	}

	folderStat, err := os.Stat(folder)
	if err != nil {
		return err
	}

	switch mode := folderStat.Mode(); {
	case mode.IsRegular():
		// We make sure the path is a folder, for example in the case a
		// Yara "index" file with includes is passed as argument, instead
		// of a full folder.
		folder = path.Dir(folder)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("Unable to create a filesystem watch for folder: %s", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Remove == fsnotify.Remove {
					callback()
				}
			}
		}
	}()

	err = watcher.Add(folder)
	if err != nil {
		if err.Error() == "no space left on device" {
			return fmt.Errorf("You might be out of inotify watches, increase the value of " +
				"fs.inotify.max_user_watches in /etc/sysctl.conf")
		}
		return fmt.Errorf("Unable to add %s to filesystem watcher: %s", folder, err)
	}
	<-done

	return nil
}
