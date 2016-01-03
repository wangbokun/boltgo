package boltgo

type Command struct {
	Name        string
	Description string
}

var boltgoCommands = []Command{
	{"cmd", " command,boltgo run $ipfile $cmd"},
	{"rsync", " files, rsync $ipfile $srcfile $dscfile"},
}

var BoltgoCommands = make(map[string]Command)

func init() {
	for _, cmd := range boltgoCommands {
		BoltgoCommands[cmd.Name] = cmd
	}
}
