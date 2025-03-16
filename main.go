package main

import (
	"github.com/teamdunno/bottle/bot"
	"github.com/teamdunno/bottle/cmds"
)

func main() {
	reg := cmds.GetRegistry()
	bot := bot.NewBot(reg)
	bot.Run()
}
