package libraries

func ContainsString(slice []string, item string) bool {
	for _, str := range slice {
		if str == item {
			return true
		}
	}

	return false
}
