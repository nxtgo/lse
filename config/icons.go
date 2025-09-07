package config

import (
	"os"
	"path/filepath"
)

func GetFileIcon(name string, mode os.FileMode, cfg Icons) string {
	switch {
	case mode.IsDir():
		return cfg.Directory
	case mode&os.ModeSymlink != 0:
		return cfg.Symlink
	case mode&os.ModeSocket != 0:
		return cfg.Socket
	case mode&os.ModeNamedPipe != 0:
		return cfg.Pipe
	case mode&os.ModeDevice != 0:
		if mode&os.ModeCharDevice != 0 {
			return cfg.CharDev
		}
		return cfg.BlockDev
	}

	ext := filepath.Ext(name)
	switch {
	case name == "LICENSE" || name == "COPYING" || name == "LICENSE.md":
		return cfg.License
	case name == "Dockerfile" || name == ".dockerignore":
		return cfg.Docker

	case ext == ".go" || ext == ".mod" || ext == ".sum":
		return cfg.Golang
	case ext == ".lock":
		return cfg.Lock
	case ext == ".ts":
		return cfg.Typescript
	case ext == ".js" || ext == ".mjs" || ext == ".cjs":
		return cfg.Javascript
	case ext == ".rs":
		return cfg.Rust
	case ext == ".py":
		return cfg.Python
	case ext == ".java" || ext == ".jar":
		return cfg.Java
	case ext == ".cs":
		return cfg.CSharp
	case ext == ".cpp" || ext == ".cxx" || ext == ".cc" || ext == ".hpp" || ext == ".hxx":
		return cfg.Cpp
	case ext == ".c" || ext == ".h":
		return cfg.C
	case ext == ".hs":
		return cfg.Haskell
	case ext == ".lua":
		return cfg.Lua
	case ext == ".rb":
		return cfg.Ruby
	case ext == ".php":
		return cfg.PHP
	case ext == ".html" || ext == ".htm":
		return cfg.HTML
	case ext == ".css":
		return cfg.CSS
	case ext == ".md" || ext == ".markdown":
		return cfg.Markdown
	case ext == ".json":
		return cfg.Json
	case ext == ".yml" || ext == ".yaml":
		return cfg.YAML
	case ext == ".toml":
		return cfg.TOML
	case ext == ".sh" || ext == ".bash" || ext == ".zsh":
		return cfg.Shell
	case ext == ".sql":
		return cfg.SQL
	case ext == ".nix":
		return cfg.Nix

	default:
		return cfg.File
	}
}
