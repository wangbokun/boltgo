package boltgo

type Command struct {
	Name        string
	Description string
}

var boltgoCommands = []Command{
	{"cmd", "run command,boltgo run $ipfile $cmd"},
	{"rsync", "rsync files"},
}

var BoltgoCommands = make(map[string]Command)

func init() {
	for _, cmd := range boltgoCommands {
		BoltgoCommands[cmd.Name] = cmd
	}
}
