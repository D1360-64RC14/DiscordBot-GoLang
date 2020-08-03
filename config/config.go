package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Data :
// Variável pública que armazenará
// a struct de configurações.
var Data configStruct

// Configure :
// Exporta o arquivo `config.yml` como uma struct.
func Configure(configFilePath string) {
	// Lê o arquivo passado pelo usuário.
	var configFile, configFileERR = ioutil.ReadFile(configFilePath)
	if configFileERR != nil {
		fmt.Printf("Erro ao abrir '%s':", configFilePath)
		fmt.Println(configFileERR)
		os.Exit(1)
	}

	// Parseia o YAML para a struct `configStruct`
	// e retorna o valor em `config`.
	var configYAMLERR = yaml.Unmarshal(configFile, &Data)
	if configYAMLERR != nil {
		fmt.Println("Erro ao ler arquivo de configuração:")
		fmt.Println(configYAMLERR)
		os.Exit(1)
	}
}
