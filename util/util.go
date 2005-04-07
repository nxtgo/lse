package util

import (
	"fmt"
	"lse/ansi"
	"lse/color"
	"lse/config"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileEntry struct {
	Path string
	Info os.FileInfo
}

func CollectEntries(pattern string, showAll bool) []FileEntry {
	paths := CollectPaths(pattern)
	var entries []FileEntry

	for _, path := range paths {
		info, err := os.Lstat(path)
		if err != nil {
			continue
		}
		name := filepath.Base(path)
		if !showAll && strings.HasPrefix(name, ".") {
			continue
		}
		entries = append(entries, FileEntry{Path: path, Info: info})
	}
	return entries
}

func SortEntries(entries []FileEntry, dirsFirst bool) {
	sort.SliceStable(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		if dirsFirst {
			if a.Info.IsDir() && !b.Info.IsDir() {
				return true
			}
			if !a.Info.IsDir() && b.Info.IsDir() {
				return false
			}
		}
		return strings.ToLower(a.Info.Name()) < strings.ToLower(b.Info.Name())
	})
}

func FormatEntry(e FileEntry, realDirSize bool, cfg config.Config) []string {
	perm := color.Permissions(e.Info.Mode().String(), cfg.Permissions)

	var sizeBytes int64
	if e.Info.IsDir() && realDirSize {
		sizeBytes = DirSize(e.Path)
	} else {
		sizeBytes = e.Info.Size()
	}
	size := color.Size(sizeBytes, cfg.Size)

	date := color.Date(e.Info.ModTime(), cfg.Date)
	fullName := color.Name(e.Info.Name(), e.Info.Mode(), cfg.Icons, cfg.FileTypes)

	return []string{perm, size, date, fullName}
}

func PrintTable(rows [][]string) {
	if len(rows) == 0 {
		return
	}

	colWidths := make([]int, len(rows[0]))
	for _, row := range rows {
		for i, col := range row {
			if w := ansi.VisibleLength(col); w > colWidths[i] {
				colWidths[i] = w
			}
		}
	}

	for _, row := range rows {
		for i, col := range row {
			fmt.Print(ansi.PadString(col, colWidths[i]))
			if i < len(row)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func CollectPaths(pattern string) []string {
	if strings.Contains(pattern, "**") {
		root := strings.Split(pattern, "**")[0]
		if root == "" {
			root = "."
		}
		var paths []string
		filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			matched, _ := filepath.Match(pattern, path)
			if matched {
				paths = append(paths, path)
			}
			return nil
		})
		return paths
	}

	info, err := os.Stat(pattern)
	if err == nil && info.IsDir() {
		entries, err := os.ReadDir(pattern)
		if err != nil {
			return nil
		}
		var paths []string
		for _, e := range entries {
			paths = append(paths, filepath.Join(pattern, e.Name()))
		}
		return paths
	}

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil
	}
	return matches
}

func DirSize(path string) int64 {
	var total int64
	err := filepath.WalkDir(path, func(_ string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return nil
			}
			total += info.Size()
		}
		return nil
	})
	if err != nil {
		return 0
	}
	return total
}
