package config

import "errors"

var MissingConfigFileError error = errors.New("No config file found.")
