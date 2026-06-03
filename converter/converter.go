package converter

import "strings"

func StringToArt(input string) string {
	digits := map[rune][]string{
		'0' : {
			" ___ ",
			" | | ",
			" | | ",
			" | | ",
			" |_| ",
		},
		'1' : {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2' : {
			" ___ ",
			" | | ",
			" _ | ",
			" |   ",
			" |_| ",
		},
		'3' : {
			" ___ ",
			"   | ",
			" __| ",
			"   | ",
			" __| ",		
		
		},
		'4' :  {
			" | | ",
			" |_|_",
			"   | ",
			"	| ",
			"	  ",
		},
		'5' : {
			" ___ ",
			" |   ",
			" |__ ",
			"   | ",
			" __| ",
		},
		'6' : {
			" ___ ",
			" |   ",
			" |__ ",
			" | | ",
			" |_| ",
		},
		'7' : {
			" ___ ",
			"   | ",
			"	| ",
			"	| ",
			"	| ",
		},
		'8' : {
			" ___ ",
			" | | ",
			" |_| ",
			" | | ",
			" |_| ",
		},
		'9' : {
			" ___ ",
			" | | ",
			" |_| ",
			"   | ",
			"   | ",
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