package main;

import "bufio";
import "fmt";
import "os";
import "errors";
import "strings";

var MissingConfigFileError error = errors.New("No config file found.");

func readConfig() {
	configPath, err := findConfigFile();
	if configPath == "" {
		if err == MissingConfigFileError {
			if len(config_locs) > 0 {
				createConfigFile(config_locs[0]);
			}
		} else {
			fmt.Println(err);
		}
	} else {
		file, err := os.OpenFile(configPath, os.O_RDONLY, 0644);
		if err != nil { fmt.Println(err); }
		var scanner *bufio.Scanner = bufio.NewScanner(file);
		var line string;
		for scanner.Scan() {
			// remove spaces and tabs
			line = strings.Trim(scanner.Text(), " \t");
			if line != "" && !strings.HasPrefix(line, "#") {
				var kv []string = strings.Split(line, "=");
				if len(kv) != 2 { fmt.Println("ERROR: bad pair"); }
				var k string = strings.Trim(kv[0], " \t");
				var v string = strings.Trim(kv[1], " \t");
				parseOption(k, v);
			}
		}
		file.Close();
		if err := scanner.Err(); err != nil { fmt.Println(err); }
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
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644);
	if err != nil {
		fmt.Println(err);
	}
	var writer *bufio.Writer = bufio.NewWriter(file);
	var line string;
	for _, line = range sample_config {
		_, err = writer.WriteString(line + "\n");
		if err != nil { fmt.Println(err); }
	}
	writer.Flush();
	file.Close();
}
