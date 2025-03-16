package bot

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

type BotConfig struct {
	Nick       string   `json:"nick"`
	Server     string   `json:"server"`
	Prefix     string   `json:"prefix"`
	Channels   []string `json:"channels"`
	Moderators []string `json:"mods"`
}

type Bot struct {
	Registry CommandRegistry
	Config   BotConfig
	Conn     *net.Conn
}

// NOT TO BE USED BY COMMANDS
func NewBot(Registry CommandRegistry) *Bot {
	bot := &Bot{
		Registry: Registry,
		Conn:     nil,
	}

	bot.ReloadConfig()

	return bot
}

func (bot *Bot) SendRaw(message string) {
	if bot.Conn != nil {
		fmt.Println("> " + message)
		(*bot.Conn).Write([]byte(message + "\r\n"))
	} else {
		fmt.Println("Not connected to a server.")
	}
}

func (b *Bot) ConnectToServer(server string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println(err)
		return
	}

	b.Conn = &conn
}

// WaitForMessage reads and parses a message from the server.
func (bot *Bot) WaitForRawMessage() (string, error) {
	reader := bufio.NewReader(*bot.Conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSpace(line)

	fmt.Println("< " + line)

	return line, nil
}

func (b *Bot) ReloadConfig() {
	conffile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	config := BotConfig{}

	decoder := json.NewDecoder(conffile)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	b.Config = config
}

func (b *Bot) Run() {
	b.ConnectToServer(b.Config.Server)

	b.SendRaw("NICK " + b.Config.Nick)
	b.SendRaw("USER 0 0 0 0")

	for _, channel := range b.Config.Channels {
		b.SendRaw("JOIN " + channel)
	}

	for {
		message, err := b.WaitForRawMessage()
		if err != nil {
			// fmt.Println(err)
		}

		if strings.Contains(message, "PRIVMSG") {
			parts := strings.Split(message, " ")
			channel := parts[2]
			user := strings.Split(parts[0], "!")[0][1:]
			message = strings.Join(parts[3:], " ")[1:]

			if strings.HasPrefix(message, b.Config.Prefix) {
				args := strings.Split(message, " ")[1:]

				ctx := NewContextBuilder().
					SetBot(b).
					SetChannel(channel).
					SetUser(user).
					SetArgs(args).
					Build()

				command := strings.Split(message, " ")[0][1:]
				go b.Registry.ExecuteCommand(command, ctx)
			}
		}

		if strings.HasPrefix(message, "PING") {
			b.SendRaw("PONG " + message[5:])
		}
	}
}

// To be used by commands
func (b *Bot) GetConfig() BotConfig {
	return b.Config
}

func (b *Bot) GetRegistry() CommandRegistry {
	return b.Registry
}

func (b *Bot) Send(channel string, message string) {
	b.SendRaw("PRIVMSG " + channel + " :" + message)
}

func (b *Bot) Sendf(channel string, format string, args ...interface{}) {
	b.Send(channel, fmt.Sprintf(format, args...))
}

func (b *Bot) SendNotice(channel string, message string) {
	b.SendRaw("NOTICE " + channel + " :" + message)
}

func (b *Bot) SendAction(channel string, message string) {
	b.SendRaw("PRIVMSG " + channel + " :\001ACTION " + message + "\001")
}
