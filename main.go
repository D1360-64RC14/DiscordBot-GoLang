package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Possíveis argumentos
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

	// Lê a primeira linha do arquivo 'token'.
	// Futuramente será adicionado um arquivo
	// de configuração mais descente.
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

	// Inicia o Bot
	var appOpenERR = app.Open()
	if appOpenERR != nil {
		fmt.Println("Erro ao abrir conexão: ")
		fmt.Println(appOpenERR)
		return
	}

	// Espera até que o usuário
	// pressione CTRL + C.
	waitAMinute()

	// Desliga o Bot.
	app.Close()
}
