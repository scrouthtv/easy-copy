package impl

func configInterpretAutoOrBoolean(v string) int {
	switch v {
	case "never", "false", "no", "none":
		return 0
	case "auto":
		return 1
	case "always", "true", "yes", "all":
		return 2
	default:
		return -1
	}
}

func configInterpretBoolean(v string) bool {
	switch v {
	case "true", "on", "yes", "always":
		return true
	default:
		return false
	}
}
