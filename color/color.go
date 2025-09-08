package color

import (
	"fmt"
	"lse/ansi"
	"lse/config"
	"os"
	"time"
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
	case bytes < 10*1024: // < 10 KB
		color = cfg.Small
	case bytes < 10*1024*1024: // < 10 MB
		color = cfg.Medium
	case bytes < 1024*1024*1024: // < 1 GB
		color = cfg.Large
	default: // >= 1 GB
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

// Date returns ANSI-colored relative time based on thresholds in cfg
func Date(t time.Time, cfg config.DateColors) string {
	diff := time.Since(t)

	var rel string
	var color string

	switch {
	case diff < time.Minute:
		color = cfg.Seconds
		rel = fmt.Sprintf("%ds ago", int(diff.Seconds()))
	case diff < time.Hour:
		color = cfg.Hours
		rel = fmt.Sprintf("%dm ago", int(diff.Minutes()))
	case diff < 24*time.Hour:
		color = cfg.Hours
		rel = fmt.Sprintf("%dh ago", int(diff.Hours()))
	case diff < 7*24*time.Hour:
		color = cfg.Days
		rel = fmt.Sprintf("%dd ago", int(diff.Hours()/24))
	default:
		color = cfg.Weeks
		rel = fmt.Sprintf("%dw ago", int(diff.Hours()/(24*7)))
	}

	return color + rel + ansi.Reset
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
