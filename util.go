package main

/**
* Actually returns a random pair.
*/
func nextMapPair(m map[string]string) string {
	for k := range m {
		return k;
	}
	return "";
}

func parseFlag(prefix string, flag string) {
	switch(flag) {
	case "h": case "help":
		printHelp();
		break;
	case "v": case "version":
		printVersion();
		break;
	case "V": case "verbose":
		verbose = true;
		break;
	case "f": case "force":
		force = true;
		break;
	case "i": case "interactive":
		interactive = true;
		break;
	default:
		unknownOption(prefix + flag);
	}
}
