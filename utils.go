package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Em minha máquina o GoLang está retornando
// UTC 3 horas atrasado.
func discordUTCtime() string {
	return time.Now().Add(3 * time.Hour).UTC().Format(time.RFC3339)
}

// Espera um sinal de CTRL + C
// para fechar o programa.
func waitAMinute() {
	consoleMessage("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

// Exibe no terminal todas as mensagens
// mensagens recebidas no servidor.
// Habilitado apenas com flag --verbose
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

	consoleMessage(log)
}

// debugMode habilitado com argumento -v
var verboseMode bool = false

// Envia mensagens para o terminal caso a
// variável `verboseMode` esteja como true.
func consoleMessage(text string) {
	if verboseMode {
		fmt.Println(text)
	}
}

// Retorna todas as keys do map
// `commands` em uma array.
func getCommandList() []string {
	var commandList []string
	for value := range commands {
		commandList = append(commandList, value)
	}
	return commandList
}
