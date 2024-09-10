package ascii

// check if our input is valid ASCII and doesn't contain invalid chars
func IsValidASCII(word string) bool {
	for _, r := range word {
		if !(r >= 32 && r <= 126) && r != '\n' {
			return false
		}
	}
	return true
}
