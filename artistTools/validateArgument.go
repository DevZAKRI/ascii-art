package ascii

func ValidateArgument(args []string) string {
	if len(args) < 2 {
		return "Error: You need to enter the STRING you want in a graphic representation using ASCII"
	} else if len(args) > 2 {
		return "Error: Only one STRING at a time is allowed!"
	}

	if !IsValidASCII(args[1]) {
		return "Error: Input contains invalid characters"
	}
	return ""
}
