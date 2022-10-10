package files

import (
	"os"
	"regexp"
)

// Expand converts a file path string containing environment variables
// such as %SystemRoot% into a an absolute path.
func Expand(path string) string {
	re := regexp.MustCompile(`\%(?i)(\w*)\%`)
	match := re.FindStringSubmatch(path)
	if match != nil {
		newValue := os.Getenv(match[1])
		if newValue != "" {
			path = re.ReplaceAllString(path, newValue)
		}

	}

	// This is sometimes used in registry keys.
	re = regexp.MustCompile(`^\\(?i)SystemRoot`)
	path = re.ReplaceAllString(path, os.Getenv("SystemRoot"))

	return path
}
