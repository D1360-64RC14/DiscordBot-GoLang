package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var debugMode bool = false

func debugMessage(text string) {
	if debugMode {
		fmt.Println(text)
	}
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help", "-h":
			fmt.Println(" Modos de inicialização:")
			fmt.Println("    -v  --view\t\tExibe todas as mensagens enviadas ao server")
			return

		case "--view", "-v":
			debugMode = true
		}
	}

	var TokenAPI, TokenAPIERR = ioutil.ReadFile("./token")
	if TokenAPIERR != nil {
		fmt.Println("Erro ao ler arquivo 'token':")
		fmt.Println(TokenAPIERR)
		return
	}
	var token = strings.Split(string(TokenAPI), "\n")[0]
	var app, appERR = discordgo.New("Bot " + token)
	if appERR != nil {
		fmt.Println("Erro ao criar a seção:")
		fmt.Println(appERR)
		return
	}

	app.AddHandler(onMessages)

	var appOpenERR = app.Open()
	if appOpenERR != nil {
		fmt.Println("Erro ao abrir conexão: ")
		fmt.Println(appOpenERR)
		return
	}

	waitAMinute()

	app.Close()
}

func waitAMinute() {
	debugMessage("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

func onMessages(session *discordgo.Session, message *discordgo.MessageCreate) {
	logMessagesToConsole(session, message) // Mostra todas as mensagens no temrinal
	if message.Author.ID == session.State.User.ID {
		return // Ignora mensagens do próprio bot
	}

	if search(message.Content, "hello") {
		session.ChannelMessageSend(message.ChannelID, "world")
	}

	if search(message.Content, "ping") {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}

	if search(message.Content, "!user") {
		if len(message.Mentions) == 0 {
			return
		}
		var user *discordgo.User = message.Mentions[0]

		userStats(session, message, user)
		session.ChannelMessageDelete(message.ChannelID, message.ID)
	}
}

func search(content, comparator string) bool {
	var final bool
	if strings.Index(strings.ToLower(content), comparator) > -1 {
		final = true
	}

	return final
}

func logMessagesToConsole(session *discordgo.Session, message *discordgo.MessageCreate) {
	var channelOptions, _ = session.Channel(message.ChannelID)

	// UTF Date | Channel Name | "User#0000": "Message content"
	var log = fmt.Sprintf(`%s | #%s | "%s#%s": "%s"`,
		message.Timestamp,
		channelOptions.Name,
		message.Author.Username,
		message.Author.Discriminator,
		message.Content,
	)

	debugMessage(log)
}

func userStats(session *discordgo.Session, message *discordgo.MessageCreate, user *discordgo.User) {
	var messageEmbed = discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: fmt.Sprintf("%s#%s", user.Username, user.Discriminator),
			URL:  user.AvatarURL("1024"),
		},
		Description: fmt.Sprintf("ID: %s\nBot: %t\nVerified: %t",
			user.ID,
			user.Bot,
			user.Verified,
		),
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Requested by %s#%s", message.Author.Username, message.Author.Discriminator),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.AvatarURL("1024"),
		},
	}

	session.ChannelMessageSendEmbed(message.ChannelID, &messageEmbed)
}
