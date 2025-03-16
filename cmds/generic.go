package cmds

import (
	"slices"

	"github.com/teamdunno/bottle/bot"
)

func init() {
	registry.AddCommand("help", help)
	registry.SetHelp("help", "show this help message (generic)")

	registry.AddCommand("reload", reload)
	registry.SetHelp("reload", "reload the bot's config (generic)")
}

func help(ctx bot.Context) {
	for name, command := range registry.GetCommands() {
		ctx.Sendf("%s: %s", name, command.Help)
	}
}

func reload(ctx bot.Context) {
	bot := ctx.LastResortGetBot()

	if !slices.Contains(bot.Config.Moderators, ctx.User) {
		ctx.Send("you need to be a channel mod to do that")
		return
	}

	bot.ReloadConfig()
	ctx.Send("bot config reloaded!")
}
