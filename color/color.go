package color

import (
	"fmt"
	"lse/ansi"
	"lse/config"
	"os"
	"strings"
)

// Permissions returns ANSI-colored permission string
func Permissions(perm string, cfg config.PermissionColors) string {
	colored := ""
	for _, c := range perm {
		switch c {
		case 'd':
			colored += cfg.Dir + string(c)
		case 'r':
			colored += cfg.Read + string(c)
		case 'w':
			colored += cfg.Write + string(c)
		case 'x':
			colored += cfg.Exec + string(c)
		case '-':
			colored += cfg.NoAccess + string(c)
		case 's':
			colored += cfg.SUID + string(c)
		case 't':
			colored += cfg.Sticky + string(c)
		default:
			colored += string(c)
		}
	}
	colored += ansi.Reset
	return colored
}

// Size returns a colored, human-readable size string from bytes
func Size(bytes int64, cfg config.SizeColors) string {
	if bytes == 0 {
		return "--"
	}

	var value float64
	var unit string

	switch {
	case bytes >= 1024*1024*1024:
		value = float64(bytes) / (1024 * 1024 * 1024)
		unit = "G"
	case bytes >= 1024*1024:
		value = float64(bytes) / (1024 * 1024)
		unit = "M"
	case bytes >= 1024:
		value = float64(bytes) / 1024
		unit = "K"
	default:
		value = float64(bytes)
		unit = "B"
	}

	var color string
	switch {
	case bytes < 1024*1024: // <10 KB
		color = cfg.Small
	case bytes < 1024*1024*1024: // <1 MB
		color = cfg.Medium
	case bytes < 1024*1024*1024*1024: // <1 GB
		color = cfg.Large
	default:
		color = cfg.Huge
	}

	var formatted string
	if unit == "B" {
		formatted = fmt.Sprintf("%d%s", int(value), unit)
	} else {
		formatted = fmt.Sprintf("%.1f%s", value, unit)
	}

	return color + formatted + ansi.Reset
}

// Date returns ANSI-colored date string (you can customize thresholds)
func Date(date string, cfg config.DateColors) string {
	if strings.Contains(date, ":") {
		return cfg.Recent + date + ansi.Reset
	} else if len(date) <= 6 {
		return cfg.Week + date + ansi.Reset
	}
	return cfg.Old + date + ansi.Reset
}

// Name returns the colored file name with icon based on type
func Name(name string, mode os.FileMode, icons config.Icons, colors config.FileTypeColors) string {
	icon := config.GetFileIcon(name, mode, icons)

	var color string
	switch {
	case mode.IsDir():
		color = colors.Directory
	case mode&os.ModeSymlink != 0:
		color = colors.Symlink
	case mode&os.ModeSocket != 0:
		color = colors.Socket
	case mode&os.ModeNamedPipe != 0:
		color = colors.Pipe
	case mode&os.ModeDevice != 0 && mode&os.ModeCharDevice == 0:
		color = colors.BlockDev
	case mode&os.ModeCharDevice != 0:
		color = colors.CharDev
	case mode&0111 != 0:
		color = colors.Exec
	default:
		color = colors.Regular
	}

	return fmt.Sprintf("%s%s %s%s", color, icon, name, ansi.Reset)
}

func Text(s, color string) string {
	return color + s + ansi.Reset
}
