package bot

type Context struct {
	Channel string
	User    string
	Args    []string
	bot     *Bot
}

func (c Context) Send(message string) {
	c.bot.Send(c.Channel, message)
}

func (c Context) Sendf(format string, args ...interface{}) {
	c.bot.Sendf(c.Channel, format, args...)
}

func (c Context) SendDirect(message string) {
	c.bot.Send(c.User, message)
}

func (c Context) SendDirectf(format string, args ...interface{}) {
	c.bot.Sendf(c.User, format, args...)
}

// This is a last resort method to get the bot instance.
// It is not recommended to use this method.
// Everything you should need is in the context.
func (c Context) LastResortGetBot() *Bot {
	return c.bot
}

type ContextBuilder struct {
	ctx Context
}

func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{}
}

func (b *ContextBuilder) SetChannel(channel string) *ContextBuilder {
	b.ctx.Channel = channel
	return b
}

func (b *ContextBuilder) SetUser(user string) *ContextBuilder {
	b.ctx.User = user
	return b
}

func (b *ContextBuilder) SetArgs(args []string) *ContextBuilder {
	b.ctx.Args = args
	return b
}

func (b *ContextBuilder) SetBot(bot *Bot) *ContextBuilder {
	b.ctx.bot = bot
	return b
}

func (b *ContextBuilder) Build() Context {
	return b.ctx
}
