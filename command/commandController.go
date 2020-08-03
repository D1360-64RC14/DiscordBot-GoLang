package command

import (
	"fmt"
	"strings"

	"github.com/D1360-64RC14/utils"
	"github.com/bwmarrin/discordgo"
)

// OnMessagesEvent :
// Evento chamado toda vez que é
// recebido uma mensagem.
func OnMessagesEvent(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Mostra log de mensagens no terminal caso modo
	// verbose esteja habilitado (flag: -v | --verbose)
	utils.LogMessagesToConsole(session, message)

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
	for command := range Commands {
		if strings.Index(strings.ToLower(message.Content), strings.ToLower(command)) == 0 {
			Commands[command](session, message)
			return
		}
	}
}

// Retorna todas as keys do map
// `commands` em uma array.
func getCommandList() []string {
	var commandList []string
	for value := range Commands {
		commandList = append(commandList, value)
	}
	return commandList
}
