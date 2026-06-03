package font

func GenerateFont() map[rune][]string{
	font := make(map[rune][]string)

	for c := rune(32); c <= rune(126); c++ {
		font[c] = generateChar(c)
	}
	return font
}
func generateChar(c rune) []string {
	if c == ' ' {
		return []string{
			"        ",
			"        ",
			"        ",
			"        ",
			"        ",
			"        ",
			"        ",
			"        ",
		}
	}
	if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return []string{
			"  ****  ",
			" *    * ",
			"*      *",
			"********",
			"*      *",
			" *    * ",
			"  ****  ",
			"        ",
		}
	}
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return []string{
			"********",
			"*      *",
			"*      *",
			"*      *",
			"*      *",
			"*      *",
			"********",
			"*      *",
		}
	}
	if c >= '0' && c <= '9' {
		return []string{
			"   **   ",
			"   **   ",
			"********",
			"********",
			"   **   ",
			"   **   ",
			"   **   ",
			"        ",
		}
	}
	return []string{
		"........",
		".      .",
		".      .",
		".      .",
		".      .",
		".      .",
		"........",
		".      .",

	}
}