package config

import (
	"encoding/json"
	"os"
)

type FileTypeColors struct {
	Directory string
	Regular   string
	Symlink   string
	BlockDev  string
	CharDev   string
	Socket    string
	Pipe      string
	Orphan    string
	Exec      string
}

type SizeColors struct {
	Small  string
	Medium string
	Large  string
	Huge   string
}

type UserGroupColors struct {
	User  string
	Group string
	Other string
}

type PermissionColors struct {
	Dir        string
	Read       string
	Write      string
	Exec       string
	ExecSticky string
	NoAccess   string
	Octal      string
	Acl        string
	Context    string
	SUID       string
	SGID       string
	Sticky     string
}

type DateColors struct {
	Seconds string
	Hours   string
	Days    string
	Weeks   string
}

type Icons struct {
	Directory string
	File      string
	Symlink   string
	Exec      string
	Socket    string
	Pipe      string
	BlockDev  string
	CharDev   string
	Orphan    string
	Image     string
	Video     string
	Audio     string
	Archive   string
	Code      string
	License   string

	// specific.
	Lock       string
	Golang     string
	Typescript string
	Javascript string
	Nix        string
	Rust       string
	Python     string
	Java       string
	CSharp     string
	Cpp        string
	C          string
	Haskell    string
	Lua        string
	Ruby       string
	PHP        string
	HTML       string
	CSS        string
	Markdown   string
	Json       string
	YAML       string
	TOML       string
	Shell      string
	Docker     string
	Kubernetes string
	SQL        string
}

type Config struct {
	Permissions PermissionColors
	Date        DateColors
	FileTypes   FileTypeColors
	Size        SizeColors
	UserGroup   UserGroupColors
	Icons       Icons
}

func DefaultConfig() Config {
	return Config{
		Permissions: PermissionColors{
			Dir:        "\033[34m", // blue
			Read:       "\033[32m", // green
			Write:      "\033[33m", // yellow
			Exec:       "\033[31m", // red
			ExecSticky: "\033[35m", // magenta
			NoAccess:   "\033[90m", // gray
			Octal:      "\033[36m", // cyan
			Acl:        "\033[34m", // blue
			Context:    "\033[37m", // white
			SUID:       "\033[41m", // red background
			SGID:       "\033[42m", // green background
			Sticky:     "\033[44m", // blue background
		},
		Date: DateColors{
			Seconds: "\033[1;31m",
			Hours:   "\033[31m",
			Days:    "\033[36m",
			Weeks:   "\033[32m",
		},
		FileTypes: FileTypeColors{
			Directory: "\033[34m",
			Regular:   "\033[0m",
			Symlink:   "\033[36m",
			BlockDev:  "\033[93m",
			CharDev:   "\033[95m",
			Socket:    "\033[35m",
			Pipe:      "\033[33m",
			Orphan:    "\033[31m",
			Exec:      "\033[1;31m",
		},
		Size: SizeColors{
			Small:  "\033[32m",
			Medium: "\033[33m",
			Large:  "\033[31m",
			Huge:   "\033[1;35m",
		},
		UserGroup: UserGroupColors{
			User:  "\033[36m",
			Group: "\033[35m",
			Other: "\033[90m",
		},
		Icons: Icons{
			Lock:       "",
			Directory:  "",
			File:       "",
			Symlink:    "",
			Exec:       "",
			Socket:     "",
			Pipe:       "󰈲",
			BlockDev:   "",
			CharDev:    "󱐋",
			Orphan:     "",
			Image:      "",
			Video:      "",
			Audio:      "",
			Archive:    "",
			Code:       "",
			License:    "",
			Golang:     "",
			Typescript: "",
			Javascript: "",
			Nix:        "󱄅",
			Rust:       "",
			Python:     "",
			Java:       "",
			CSharp:     "󰌛",
			Cpp:        "",
			C:          "",
			Haskell:    "",
			Lua:        "",
			Ruby:       "",
			PHP:        "",
			HTML:       "",
			CSS:        "",
			Markdown:   "",
			Json:       "",
			YAML:       "",
			TOML:       "",
			Shell:      "",
			Docker:     "󰡨",
			Kubernetes: "󱃾",
			SQL:        "",
		},
	}
}

func LoadConfig(path string) Config {
	cfg := DefaultConfig()

	f, err := os.Open(path)
	if err != nil {
		return cfg
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	if err := dec.Decode(&cfg); err != nil {
		return DefaultConfig()
	}

	return cfg
}
