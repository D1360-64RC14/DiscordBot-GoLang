package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// GetNowUTCTime :
// Retorna tempo em UTC
// compatível com Discord.
func GetNowUTCTime() string {
	return time.Now().Add(3 * time.Hour).UTC().Format(time.RFC3339)
}

// WaitAMinute :
// Espera um sinal de CTRL + C
// para fechar o programa.
func WaitAMinute() {
	ConsoleMessage("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

// LogMessagesToConsole :
// Exibe no terminal todas as mensagens
// recebidas no servidor.
// Habilitado apenas com flag --verbose
func LogMessagesToConsole(session *discordgo.Session, message *discordgo.MessageCreate) {
	var channelOptions, _ = session.Channel(message.ChannelID)

	// UTF Date | Channel Name | "User#0000": "Message content"
	var log = fmt.Sprintf(`%s | #%s | "%s#%s": "%s"`,
		message.Timestamp,
		channelOptions.Name,
		message.Author.Username,
		message.Author.Discriminator,
		message.Content,
	)

	ConsoleMessage(log)
}

// VerboseMode :
// Variável pública utilizada para
// habilitar modo verboso.
var VerboseMode bool = false

// ConsoleMessage :
// Envia mensagens para o terminal caso a
// variável `VerboseMode` esteja como true.
func ConsoleMessage(text string) {
	if VerboseMode {
		fmt.Println(text)
	}
}

// RGB2Int :
// Converte valor RGB
// para Integer
func RGB2Int(red, green, blue uint8) uint32 {
	if red > 255 {
		red = 255
	}
	if green > 255 {
		green = 255
	}
	if blue > 255 {
		blue = 255
	}

	if red < 0 {
		red = 0
	}
	if green < 0 {
		green = 0
	}
	if blue < 0 {
		blue = 0
	}
	return 1<<16*uint32(red) + 256*uint32(green) + uint32(blue)
}
