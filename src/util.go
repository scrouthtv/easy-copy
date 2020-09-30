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
