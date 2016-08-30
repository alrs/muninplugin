package muninplugin

// toYN converts boolean types to "yes" or "no" strings.
func toYN(b bool) string {
	if b {
		return "yes"
	} else {
		return "no"
	}
}
