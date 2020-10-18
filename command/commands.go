package command

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/D1360-64RC14/config"
	"github.com/D1360-64RC14/utils"
	"github.com/D1360-64RC14/youtube"
	"github.com/bwmarrin/discordgo"
)

// Commands :
// Map de todos os comandos disponíveis.
var Commands = map[string]func(session *discordgo.Session, message *discordgo.MessageCreate){

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
			Timestamp: utils.GetNowUTCTime(),
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("Requested by %s#%s", message.Author.Username, message.Author.Discriminator),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: user.AvatarURL("1024"),
			},
		}
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

	// Retorna um Embed com informações de um video do youtube.
	// Possível utilização de pesquisa ou video ID.
	// Utlização:
	//   - !youtubesearch <pesquisa> | <video id>
	"!youtubesearch": func(session *discordgo.Session, message *discordgo.MessageCreate) {
		// Argumentos de pesquisa após o comando.
		var search = message.Content[strings.Index(message.Content, " ")+1:]

		var youtubeSearchStruct = youtube.GetSearch(search)
		var youtubeChannelsStruct = youtube.GetChannels(youtubeSearchStruct.Items[0].Snippet.ChannelID)

		var videoUploadTime, _ = time.Parse(time.RFC3339, youtubeSearchStruct.Items[0].Snippet.PublishTime)
		// Formato de data:
		// DD/MM/AAAA HH:MM:SS
		var footerString = fmt.Sprintf(
			"Vídeo lançado em %02d/%02d/%d, ás %02d:%02d:%02d",
			videoUploadTime.Hour(),
			videoUploadTime.Minute(),
			videoUploadTime.Second(),

			videoUploadTime.Day(),
			videoUploadTime.Month(),
			videoUploadTime.Year(),
		)

		if youtubeSearchStruct.Error != nil {
			session.ChannelMessageSend(message.ChannelID, youtubeSearchStruct.Error.String())
		} else {
			var messageEmbed = &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{
					Name: html.UnescapeString(youtubeChannelsStruct.Items[0].Snippet.Title),
					URL:  fmt.Sprintf("%s/%s", config.Data.YoutubeAPI.RequestURLs.Channels, youtubeSearchStruct.Items[0].Snippet.ChannelID),
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: youtubeChannelsStruct.Items[0].Snippet.Thumbnails.High.URL,
				},
				Image: &discordgo.MessageEmbedImage{
					URL: youtube.GetMaxResThumb(youtubeSearchStruct.Items[0].ID.VideoID),
				},
				Title: html.UnescapeString(youtubeSearchStruct.Items[0].Snippet.Title),
				URL:   fmt.Sprintf("%s%s", config.Data.YoutubeAPI.RequestURLs.Video, youtubeSearchStruct.Items[0].ID.VideoID),
				Color: int(utils.RGB2Int(255, 20, 20)),
				Footer: &discordgo.MessageEmbedFooter{
					Text: footerString,
				},
			}

			session.ChannelMessageSendEmbed(message.ChannelID, messageEmbed)
		}
	},
}
