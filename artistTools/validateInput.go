package ascii

import "os"

func ValidateInput() string {
	if len(os.Args) < 2 {
		return "Error: You need to enter the STRING you want in a graphic representation using ASCII\n"
	} else if len(os.Args) > 2 {
		return "Error: Only one STRING at a time is allowed!\n"
	}

	if !IsValidASCII(os.Args[1]) {
		return "Error: Input contains invalid characters\n"
	}
	return ""
}
