package main

import (
	"flag"
	"fmt"
	"lse/ansi"
	"lse/color"
	"lse/config"
	"lse/util"
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

type FileRow struct {
	Perm string
	Size string
	Date string
	Name string
}

func main() {
	cfg := config.LoadConfig(configFile)

	pattern := "."
	if flag.NArg() > 0 {
		pattern = flag.Arg(0)
	}

	paths := util.CollectPaths(pattern)

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

	type result struct {
		row FileRow
		ok  bool
	}
	results := make(chan result, len(paths))
	sem := make(chan struct{}, 16)

	for _, path := range paths {
		sem <- struct{}{}
		go func(p string) {
			defer func() { <-sem }()

			info, err := os.Lstat(p)
			if err != nil {
				results <- result{ok: false}
				return
			}

			var sizeBytes int64
			if info.IsDir() && realDirSize {
				sizeBytes = util.DirSize(p)
			} else {
				sizeBytes = info.Size()
			}

			row := FileRow{
				Perm: color.Permissions(info.Mode().String(), cfg.Permissions),
				Size: color.Size(sizeBytes, cfg.Size),
				Date: color.Date(info.ModTime(), cfg.Date),
				Name: color.Name(filepath.Base(p), info.Mode(), cfg.Icons, cfg.FileTypes),
			}

			results <- result{row: row, ok: true}
		}(path)
	}

	var files []FileRow
	for i := 0; i < len(paths); i++ {
		res := <-results
		if res.ok {
			files = append(files, res.row)
		}
	}

	if len(files) == 0 {
		return
	}

	colWidths := make([]int, 4)
	for _, f := range files {
		cols := []string{f.Perm, f.Size, f.Date, f.Name}
		for i, col := range cols {
			if w := ansi.VisibleLength(col); w > colWidths[i] {
				colWidths[i] = w
			}
		}
	}

	for _, f := range files {
		cols := []string{f.Perm, f.Size, f.Date, f.Name}
		for i, col := range cols {
			fmt.Print(ansi.PadString(col, colWidths[i]))
			if i < len(cols)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
