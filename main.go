package main

import (
	"flag"
	"lse/config"
	"lse/util"
)

var (
	configFile  string
	dirsFirst   bool
	showAll     bool
	realDirSize bool
)

func init() {
	flag.StringVar(&configFile, "cfg", "~/.config/lse.json", "config file")
	flag.BoolVar(&dirsFirst, "d", false, "show directories first")
	flag.BoolVar(&showAll, "a", false, "show hidden files")
	flag.BoolVar(&realDirSize, "s", false, "show real dir size")
	flag.Parse()
}

func main() {
	cfg := config.LoadConfig(configFile)

	pattern := "."
	if flag.NArg() > 0 {
		pattern = flag.Arg(0)
	}

	entries := util.CollectEntries(pattern, showAll)
	if len(entries) == 0 {
		return
	}

	util.SortEntries(entries, dirsFirst)

	var rows [][]string
	for _, e := range entries {
		rows = append(rows, util.FormatEntry(e, realDirSize, cfg))
	}

	util.PrintTable(rows)
}
