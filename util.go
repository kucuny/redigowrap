package redis

func getBool(result string) bool {
	if result == "OK" {
		return true
	} else {
		return false
	}
}
