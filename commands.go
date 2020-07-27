package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(session *discordgo.Session, message *discordgo.MessageCreate){

	// Retorna um Embed com as informações:
	// Username, id, is bot?, is verified?, profile image.
	// Utilização:
	//   - !user <username>
	//   - Execução do comando sem nome de usuário
	//     retorna as informações do owner da mensagem.
	"!user": func(session *discordgo.Session, message *discordgo.MessageCreate) {
		var user *discordgo.User
		if len(message.Mentions) == 0 {
			user = message.Author
		} else {
			user = message.Mentions[0]
		}
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
			Timestamp: discordUTCtime(),
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("Requested by %s#%s", message.Author.Username, message.Author.Discriminator),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: user.AvatarURL("1024"),
			},
		}
		session.ChannelMessageDelete(message.ChannelID, message.ID)
		session.ChannelMessageSendEmbed(message.ChannelID, &messageEmbed)
	},

	// Verificação de atividade do servidor.
	// Utilização:
	//   - !ping
	"!ping": func(session *discordgo.Session, message *discordgo.MessageCreate) {
		session.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf("Pong %s!", message.Author.Mention()),
		)
	},
}
