package cmds

import (
	"strings"

	"github.com/teamdunno/bottle/bot"
)

func init() {
	registry.AddCommand("patpat", patpat)
	registry.SetHelp("patpat", "pat someone (fun)")

	registry.AddCommand("gopher", gopher)
	registry.SetHelp("gopher", "show a picture of a gopher (fun)")
}

func patpat(ctx bot.Context) {
	if len(ctx.Args) < 1 {
		ctx.SendDirect("you need to give someone to patpat")
		return
	}

	target := strings.Join(ctx.Args[0:], " ")

	ctx.Sendf("%s pats %s", ctx.User, target)
}

func gopher(ctx bot.Context) {
	ctx.Send("Hello! I was written in Go.")
	ctx.Send("Here's a gopher: https://golang.org/doc/gopher/frontpage.png")
}
