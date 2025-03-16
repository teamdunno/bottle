package cmds

import "github.com/teamdunno/bottle/bot"

var registry = bot.NewCommandRegistry()

func GetRegistry() bot.CommandRegistry {
	return registry
}

func init() {}
