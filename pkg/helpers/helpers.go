package helpers

func Contains(history []string, recipient string) bool {
	for _, r := range history {
		if r == recipient {
			return true
		}
	}
	return false
}
