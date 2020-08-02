package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessagesEvent(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Mostra log de mensagens no terminal caso modo
	// verbose esteja habilitado (flag: -v | --verbose)
	logMessagesToConsole(session, message)

	// Ignora mensagens do próprio bot
	if message.Author.ID == session.State.User.ID {
		return
	}

	// Comando que envia toda a lista de comandos.
	// Não foi colocado dentro do commands map por
	// erro de Initialization Loop.
	if strings.Index(strings.ToLower(message.Content), "!comandos") > -1 || strings.Index(strings.ToLower(message.Content), "!commands") > -1 {
		session.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf("Comandos disponíveis:\n`%s`.", strings.Join(getCommandList(), ", ")),
		)
	}

	// Passa pela lista de comandos verificando
	// se existe algum na mensagem. Ao encontrar,
	// o executa.
	// Match apenas com o comando
	// estando no início da mensagem.
	for command := range commands {
		if strings.Index(strings.ToLower(message.Content), strings.ToLower(command)) == 0 {
			commands[command](session, message)
			return
		}
	}
}
