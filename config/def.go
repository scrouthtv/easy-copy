package config

import "errors"

// ErrMissingConfigFile is returned if no ec.conf could be found.
var ErrMissingConfigFile error = errors.New("no config file found")

// ErrReadingConfigFile is returned if an ec.conf could be found,
// but an error occurred reading the file.
type ErrReadingConfigFile struct {
	err  error
	path string
}

func (e *ErrReadingConfigFile) Error() string {
	return "error reading the config file at " + e.path + ": " + e.err.Error()
}

func (e *ErrReadingConfigFile) Unwrap() error {
	return e.err
}
