package impl

import (
	"easy-copy/config"
	"strings"
)

// LoadConfig checks if the config file should be read
// (e.g. option no-config is not present),
// and reads it if we want to.
func (s *settingsImpl) LoadConfig(args []string) error {
	for _, arg := range args {
		if arg == "--no-config" {
			return nil
		} else if arg == "--" {
			break
		}
	}

	kvs, err := config.Load()
	if err != nil {
		return err
	}

	for _, line := range kvs {
		s.parseOption(line)
	}

	return nil
}

func (s *settingsImpl) parseOption(line string) error {
	kv := strings.Split(line, "=")
	if len(kv) != 2 {
		return &ErrBadConfigLine{line}
	}

	s.parseKeyValue(kv[0], kv[1])

	return nil
}
