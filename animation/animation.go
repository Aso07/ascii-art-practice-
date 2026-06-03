package animation

import (
	"strings"
	"fmt"
)

type Animation struct {
	text string
	frames int
	data []string
}
func NewAnimation(text string, frames int) *Animation {
	return &Animation{
		text: text,
		frames: frames,
		data: make([]string, frames),
	}
}
func (a *Animation) buildFrame(textLine int, prefix string, suffix string) string {
	textRow := prefix + a.text + suffix
	width := len(textRow)
	padding := strings.Repeat(" ", width)

	rows := make([]string, 10)
	for i := range rows {
		if i == textLine {
			rows[i] = textRow
		} else {
			rows[i] = padding
		}
	}
	return strings.Join(rows, "\n") + "\n"
}
func (a *Animation) GenerateSpinFrames() {
	spinChars := []string{"|", "/", "-", "\\"}
	for i := 0; i < a.frames; i++ {
		char := spinChars[i%4]
		a.data[i] = a.buildFrame(5, char+" ", " "+char)
	}
}
func (a *Animation) GenerateWaveFrames() {
	for i := 0; i < a.frames; i++ {
		line := i % 8
		a.data[i] = a.buildFrame(line, " ", " ")
	}
}
func (a *Animation) GenerateZoomFrames() {
	for i := 0; i < a.frames; i++ {
		spaces := strings.Repeat(" ", i)
		a.data[i] = a.buildFrame(5, spaces, spaces)
	}
}
func (a *Animation) GetFrame(index int) string {
	return a.data[index%a.frames]
}
func (a *Animation) Play() string {
	var result strings.Builder
	for i, frame := range a.data {
		result.WriteString(fmt.Sprintf("--- Frame %d ---\n", i))
		result.WriteString(frame)
	}
	return result.String()
}