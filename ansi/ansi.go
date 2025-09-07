package ansi

import (
	"regexp"
)

const Reset = "\033[0m"

var Regexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func VisibleLength(s string) int {
	plain := Regexp.ReplaceAllString(s, "")

	count := 0
	for _, r := range plain {
		if r >= 0xf000 {
			count += 2
		} else {
			count += 1
		}
	}
	return count
}

func PadString(s string, width int) string {
	visible := VisibleLength(s)
	if visible >= width {
		return s
	}
	return s + repeat(" ", width-visible)
}

func repeat(s string, n int) string {
	out := ""
	for range n {
		out += s
	}
	return out
}
