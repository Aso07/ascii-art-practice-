package converter

import "strings"

func StringToArt(input string) string {
	digits := map[rune][]string{
		'0' : []string{
			" ___ ",
			" |  | ",
			" |  | ",
			" |  | ",
			" |__| ",
		},
		'1' : []string{
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2' : []string{
			" ___ ",
			" |  | ",
			" __ | ",
			" |    ",
			" |__| ",
		},
		'3' : []string{
			" ___ ",
			"    | ",
			"  __| ",
			"    | ",
			"  __| ",		
		
		},
		'4' :  []string{
			" | | ",
			" |_|_",
			"   | ",
			"	| ",
			"	  ",
		},
		'5' : []string{
			" ___ ",
			" |    ",
			" |___ ",
			"    | ",
			" ___| ",
		},
		'6' : []string{
			" ___ ",
			" |    ",
			" |___ ",
			" |  | ",
			" |__| ",
		},
		'7' : []string{
			" ___ ",
			"   | ",
			"	| ",
			"	| ",
			"	| ",
		},
		'8' : []string{
			" ___ ",
			" |  | ",
			" |__| ",
			" |  | ",
			" |__| ",
		},
		'9' : []string{
			" ___ ",
			" |  | ",
			" |__| ",
			"    | ",
			"    | ",
		},

	}
	if input == "" {
		return ""
	}
	lines := strings.Split(input, "\n")
	var result strings.Builder

	for _, line := range lines {
		for _, c := range line {
			if c < '0' || c > '9' {
				return ""
			}
		}
	
	for row := 0; row < 5; row++ {
		for _, c := range line {
			result.WriteString(digits[c][row])
		}
		result.WriteString("\n")
	}
	}
	return result.String()
}