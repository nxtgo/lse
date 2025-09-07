# lse (ls-enhanced)

enhanced `ls` command.

## usage

```bash
lse
```

## flags

* `-a` - *show hidden files*
* `-d` - *display dirs first*

## default config

```json
{
  "Permissions": {
    "Dir": "\u001b[34m",
    "Read": "\u001b[32m",
    "Write": "\u001b[33m",
    "Exec": "\u001b[31m",
    "ExecSticky": "\u001b[35m",
    "NoAccess": "\u001b[90m",
    "Octal": "\u001b[36m",
    "Acl": "\u001b[34m",
    "Context": "\u001b[37m",
    "SUID": "\u001b[41m",
    "SGID": "\u001b[42m",
    "Sticky": "\u001b[44m"
  },
  "Date": {
    "Recent": "\u001b[32m",
    "Week": "\u001b[33m",
    "Old": "\u001b[90m"
  },
  "FileTypes": {
    "Directory": "\u001b[34m",
    "Regular": "\u001b[0m",
    "Symlink": "\u001b[36m",
    "BlockDev": "\u001b[93m",
    "CharDev": "\u001b[95m",
    "Socket": "\u001b[35m",
    "Pipe": "\u001b[33m",
    "Orphan": "\u001b[31m",
    "Exec": "\u001b[32m"
  },
  "Size": {
    "Small": "\u001b[37m",
    "Medium": "\u001b[33m",
    "Large": "\u001b[31m",
    "Huge": "\u001b[35m"
  },
  "UserGroup": {
    "User": "\u001b[36m",
    "Group": "\u001b[35m",
    "Other": "\u001b[90m"
  },
  "Icons": {
    "Lock": "",
    "Directory": "",
    "File": "",
    "Symlink": "",
    "Exec": "",
    "Socket": "",
    "Pipe": "󰈲",
    "BlockDev": "",
    "CharDev": "󱐋",
    "Orphan": "",
    "Image": "",
    "Video": "",
    "Audio": "",
    "Archive": "",
    "Code": "",
    "License": "",
    "Golang": "",
    "Typescript": "",
    "Javascript": "",
    "Nix": "󱄅",
    "Rust": "",
    "Python": "",
    "Java": "",
    "CSharp": "󰌛",
    "Cpp": "",
    "C": "",
    "Haskell": "",
    "Lua": "",
    "Ruby": "",
    "PHP": "",
    "HTML": "",
    "CSS": "",
    "Markdown": "",
    "Json": "",
    "YAML": "",
    "TOML": "",
    "Shell": "",
    "Docker": "󰡨",
    "Kubernetes": "󱃾",
    "SQL": ""
  }
}
```

# license

CC0 1.0 (public domain) + ip waiver.
