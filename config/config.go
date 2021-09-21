//go:build !noconfig

package config

import (
	"bufio"
	"os"
	"strings"
)

// Load reads and parses the first config that can be found.
// If an error occurred (no config file / error reading the file),
// the error is returned.
// If not, an array of lines for reading is returned.
func Load() ([]string, error) {
	var kvs []string

	configPath, err := findConfigFile()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(configPath, os.O_RDONLY, 0o644)
	if err != nil {
		return nil, &ErrReadingConfigFile{err, configPath}
	}

	scanner := bufio.NewScanner(file)
	var line string

	for scanner.Scan() {
		// remove spaces and tabs
		line = strings.Trim(scanner.Text(), " \t")
		if line != "" && !strings.HasPrefix(line, "#") {
			kvs = append(kvs, line)
		}
	}
	file.Close()

	err = scanner.Err()
	if err != nil {
		return nil, &ErrReadingConfigFile{err, configPath}
	}

	return kvs, nil
}

// findConfigFile searches a valid config file.
// If none is found, an empty string is returned
// and one has to be created.
func findConfigFile() (string, error) {
	var err error
	for i := 0; i < len(configLocs); i++ {
		_, err = os.Stat(configLocs[i])
		if err == nil {
			return configLocs[i], nil
		}
	}

	return "", ErrMissingConfigFile
}
