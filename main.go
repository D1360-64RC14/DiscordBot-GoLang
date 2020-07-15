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

func main() {
	var TokenAPI, _ = ioutil.ReadFile("./token")
	var app, err = discordgo.New("Bot " + string(TokenAPI))
	if err != nil {
		fmt.Print("Erro ao criar a seção: ")
		fmt.Println(err)
		return
	}

	app.AddHandler(onMessages)

	err = app.Open()
	if err != nil {
		fmt.Print("Erro ao abrir conexão: ")
		fmt.Println(err)
		return
	}

	// <<< DAQUI ATÉ A LINHA 37 EU NÃO SEI O QUE ESTÁ ACONTECENDO >>>
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	app.Close()
}

func onMessages(session *discordgo.Session, message *discordgo.MessageCreate) {

	logMessagesToConsole(message) // Mostra todas as mensagens no temrinal
	if message.Author.ID == session.State.User.ID {
		return // Ignora mensagens do próprio bot
	}

	if strings.Index(strings.ToLower(message.Content), "hello") > -1 {
		session.ChannelMessageSend(message.ChannelID, "world")
	}

	if strings.Index(strings.ToLower(message.Content), "ping") > -1 {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}

func logMessagesToConsole(message *discordgo.MessageCreate) {
	var Time = fmt.Sprintf(`%d/%d/%d %d:%d:%d`,
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Day(),
	)

	var text = fmt.Sprintf(`%s | "%s#%s": %s`, // Eu não sei como pegar o nome do canal de texto
		Time,
		message.Author.Username,
		message.Author.Discriminator,
		message.Content,
	)

	fmt.Println(text)
}
