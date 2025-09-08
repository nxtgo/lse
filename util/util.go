package util

import (
	"os"
	"path/filepath"
	"strings"
)

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
