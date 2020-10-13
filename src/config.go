package main;

import "bufio";
import "os";
import "errors";
import "strings";

var MissingConfigFileError error = errors.New("No config file found.");
var doReadConfig bool = true;

func readConfig() {
	if !doReadConfig { return; }
	configPath, err := findConfigFile();
	if configPath == "" {
		if err == MissingConfigFileError {
			if len(config_locs) > 0 {
				createConfigFile(config_locs[0]);
			}
		} else {
			warnConfig(err);
		}
	} else {
		file, err := os.OpenFile(configPath, os.O_RDONLY, 0644);
		if err != nil {
			warnConfig(err);
			return;
		}
		var scanner *bufio.Scanner = bufio.NewScanner(file);
		var line string;
		for scanner.Scan() {
			// remove spaces and tabs
			line = strings.Trim(scanner.Text(), " \t");
			if line != "" && !strings.HasPrefix(line, "#") {
				parseOption(line);
			}
		}
		file.Close();
		if err := scanner.Err(); err != nil {
			warnConfig(err);
			return;
		}
	}
}

/**
 * Search a valid config file
 * If none is found "" is returned and one has to be created
 */
func findConfigFile() (string, error) {
	var err error;
	for i := 0; i < len(config_locs); i++ {
		_, err = os.Stat(config_locs[i]);
		if err == nil { return config_locs[i], nil; }
	}
	return "", MissingConfigFileError;
}

func createConfigFile(filePath string) {
	var err error;
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0644);
	if err != nil {
		warnCreatingConfig(err);
	}
	var writer *bufio.Writer = bufio.NewWriter(file);
	var line string;
	for _, line = range sample_config {
		_, err = writer.WriteString(line + "\n");
		if err != nil { warnCreatingConfig(err); }
	}
	writer.Flush();
	file.Close();
}
