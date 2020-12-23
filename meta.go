package main

import "path/filepath"
import "os"

const EASYCOPY_NAME string = "EasyCopy"
const EASYCOPY_VERSION string = "0.3.0"
const EASYCOPY_ISSUES string = "https://github.com/scrouthtv/easy-copy"

// these ones are also hardcoded in the manpage

const folder_size int = 4
const symlink_size int = 16

var config_locs []string

func init() {
	var config_dir, home_dir string
	var err error
	config_dir, err = os.UserConfigDir()
	if err != nil {
		return
	}
	home_dir, err = os.UserHomeDir()
	if err != nil {
		return
	}

	config_locs = []string{
		filepath.Join(config_dir, "ec.conf"),
		filepath.Join(config_dir, "ec/", "ec.conf"),
		filepath.Join(home_dir, "ec.conf")}
}

var sample_config []string = []string{
	"# Print more information:",
	"verbose = false",
	"",
	"# Behaviour if a target file already exists:",
	"#  skip - overwrite - ask",
	"overwrite = ask",
	"",
	"# Handling of symbolic links in source:",
	"#  ignore - link - dereference",
	"symlinks = link"}
