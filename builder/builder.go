package builder
 
import "strings"

type ArtBuilder struct {
	text string
	style string
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{}
}

func (box *ArtBuilder) AddText(text string) *ArtBuilder {
	box.text += text
	return box
}

func (box *ArtBuilder) SetStyle(style string) *ArtBuilder {
	validStyles := map[string]bool{
		"normal": true, "bold": true, "italic": true, "outline": true,
	}
	if !validStyles[style] {
		panic("invalid style:" + style)
	}
	box.style = style
	return box
}

func (box *ArtBuilder) Build() string {
	font := map[rune][]string{
		'A' : {
			"   ##   ",
			"  #  #  ",
			"  #  #  ",
			"  ####  ",
			"  #  #  ",
			"  #  #  ",
			"  #  #  ",
			"        ",
		},
		'B' : {
			"####### ",
			"#      #",
			"#      #",
			"####### ",
			"#      #",
			"#      #",
			"####### ",
			"        ",
		},
		'H' : {
			"#      #",
			"#      #",
			"#      #",
			"########",
			"#      #",
			"#      #",
			"#      #",
			"        ",
		},
		'I' : {
			" ###### ",
			"   ##   ",
			"   ##   ",
			"   ##   ",
			"   ##   ",
			"   ##   ",
			" ###### ",
			"        ",
		},
		'T' : {
			"########",
			"   ##   ",
			"   ##   ",
			"   ##   ",
			"   ##   ",
			"   ##   ",
			"   ##   ",
			"        ",
		},
	}
	defaultPattern := []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
	}
	rows := make ([]string, 8)

	for _, c := range box.text {
		pattern, exists := font[c]
		if !exists {
			pattern = defaultPattern
		}
		for i := 0; i < 8; i++ {
			rows[i] += applyStyle(pattern[i], box.style)
		}
	}
	return strings.Join(rows, "\n") + "\n"
}
func applyStyle(line string, style string) string {
	switch style {
	case "bold" :
		var b strings.Builder
		for _, c := range line {
			b.WriteRune(c)
			b.WriteRune(c)
		}
		return b.String()
	case "italic":
		return " " + line
	case "outline":
		return "|" + line + "|"
	default:
		return line
	}
}
