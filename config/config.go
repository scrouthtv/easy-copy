// +build !noconfig

package config

import (
	"bufio"
	"os"
	"strings"
)

// Load reads and parses the first config that can be found.
func Load() ([]string, error) {
	var kvs []string
	configPath, err := findConfigFile()
	if configPath == "" {
		if err == MissingConfigFileError {
			if len(config_locs) > 0 {
				err = createConfigFile(config_locs[0])
				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	} else {
		file, err := os.OpenFile(configPath, os.O_RDONLY, 0o644)
		if err != nil {
			return nil, err
		}
		var scanner *bufio.Scanner = bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			// remove spaces and tabs
			line = strings.Trim(scanner.Text(), " \t")
			if line != "" && !strings.HasPrefix(line, "#") {
				kvs = append(kvs, line)
			}
		}
		file.Close()
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return kvs, nil
}

/**
 * Search a valid config file
 * If none is found "" is returned and one has to be created
 */
func findConfigFile() (string, error) {
	var err error
	for i := 0; i < len(config_locs); i++ {
		_, err = os.Stat(config_locs[i])
		if err == nil {
			return config_locs[i], nil
		}
	}
	return "", MissingConfigFileError
}

func createConfigFile(filePath string) error {
	var err error
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0o644)
	defer file.Close()
	if err != nil {
		return err
	}
	var writer *bufio.Writer = bufio.NewWriter(file)
	var line string
	for _, line = range sample_config {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}
