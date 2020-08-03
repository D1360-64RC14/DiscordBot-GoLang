package main

import (
	"fmt"
	"os"

	"github.com/D1360-64RC14/command"
	"github.com/D1360-64RC14/config"
	"github.com/D1360-64RC14/utils"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Possíveis argumentos
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help", "-h":
			fmt.Println(" Modos de inicialização:")
			fmt.Println("    -v  --verbose\t\tExibe todas as mensagens enviadas ao server")
			os.Exit(0)

		case "--verbose", "-v":
			utils.VerboseMode = true
		}
	}

	// Parseia o arquivo YAML numa struct
	// e disponibiliza na variável `config`.
	config.Configure("./config.yml")

	// Configura o Bot
	var app, appERR = discordgo.New(fmt.Sprintf("Bot %s", config.Data.DiscordBot.Token))
	if appERR != nil {
		fmt.Println("Erro ao criar a seção:")
		fmt.Println(appERR)
		os.Exit(1)
	}

	// Event listener de mensagens.
	app.AddHandler(command.OnMessagesEvent)

	// Inicia o Bot
	var appOpenERR = app.Open()
	if appOpenERR != nil {
		fmt.Println("Erro ao abrir conexão:")
		fmt.Println(appOpenERR)
		os.Exit(1)
	}

	// Espera até que o usuário
	// pressione CTRL + C.
	utils.WaitAMinute()

	// Desliga o Bot e sai com exit code 0.
	app.Close()
	os.Exit(0)
}
