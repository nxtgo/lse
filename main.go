package main

import (
	"flag"
	"fmt"
	"lse/ansi"
	"lse/color"
	"lse/config"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var configFile string
var dirsFirst bool
var showAll bool
var realDirSize bool

func init() {
	flag.StringVar(&configFile, "cfg", "~/.config/lse.json", "config file")
	flag.BoolVar(&dirsFirst, "d", false, "show directories first")
	flag.BoolVar(&showAll, "a", false, "show hidden files")
	flag.BoolVar(&realDirSize, "s", false, "show real dir size")
	flag.Parse()
}

// todo: user & groups
func main() {
	cfg := config.LoadConfig(configFile)

	pattern := "."
	if flag.NArg() > 0 {
		pattern = flag.Arg(0)
	}

	paths := collectPaths(pattern)

	if !showAll {
		var filtered []string
		for _, path := range paths {
			name := filepath.Base(path)
			if !strings.HasPrefix(name, ".") {
				filtered = append(filtered, path)
			}
		}
		paths = filtered
	}

	if len(paths) == 0 {
		return
	}

	if dirsFirst {
		sort.SliceStable(paths, func(i, j int) bool {
			infoI, errI := os.Lstat(paths[i])
			infoJ, errJ := os.Lstat(paths[j])
			if errI != nil || errJ != nil {
				return false
			}
			if infoI.IsDir() && !infoJ.IsDir() {
				return true
			}
			if !infoI.IsDir() && infoJ.IsDir() {
				return false
			}
			return strings.ToLower(filepath.Base(paths[i])) < strings.ToLower(filepath.Base(paths[j]))
		})
	}

	var files [][]string
	for _, path := range paths {
		info, err := os.Lstat(path)
		if err != nil {
			continue
		}

		perm := color.Permissions(info.Mode().String(), cfg.Permissions)
		var sizeBytes int64
		if info.IsDir() && realDirSize {
			sizeBytes = dirSize(path)
		} else {
			sizeBytes = info.Size()
		}
		size := color.Size(sizeBytes, cfg.Size)
		date := color.Date(info.ModTime().Format("Mon Jan 02 15:04:05 2006"), cfg.Date)
		name := filepath.Base(path)
		fullName := color.Name(name, info.Mode(), cfg.Icons, cfg.FileTypes)

		files = append(files, []string{perm, size, date, fullName})
	}

	if len(files) == 0 {
		return
	}

	colWidths := make([]int, len(files[0]))
	for _, f := range files {
		for i, col := range f {
			if w := ansi.VisibleLength(col); w > colWidths[i] {
				colWidths[i] = w
			}
		}
	}

	for _, f := range files {
		for i, col := range f {
			fmt.Print(ansi.PadString(col, colWidths[i]))
			if i < len(f)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func collectPaths(pattern string) []string {
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

func dirSize(path string) int64 {
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
