OUT_FILE_NAME = discord-bot_byD1360
FILES = main.go commands.go commandController.go utils.go config.go

# [ Flags ] (go help build)
# -v    print the names of packages as they are compiled.
# -x    print the commands.
BUILD_FLAGS = -v -x

# Com flag ( -v | --verbose ) exibe
# log de mensagens no terminal
run : $(FILES)
	@ echo "=== Rodando progrma sem compilar ==="
	go run $(FILES) -v
	@ echo ""

build : $(FILES) config.yml
	@ echo "=== Criando diretório build/ ==="
	mkdir build
	@ echo "=== Compilando programa ==="
	go build -o ./build/$(OUT_FILE_NAME) $(BUILD_FLAGS) $(FILES)
	cp config.yml ./build/config.yml
	@ echo ""
	@ echo "=== Done ==="

buildrun : $(FILES) config.yml
	@ echo "=== Compilando e rodando o programa ==="
	go build -o ./build/$(OUT_FILE_NAME) $(BUILD_FLAGS) $(FILES)
	cp config.yml ./build/config.yml
	@ echo ""
	@ echo "=== Done ==="
	@ echo "=== Executando programa ==="
	./build/$(OUT_FILE_NAME)

clean : 
	@ echo "=== Limpando build/ ==="
	rm -rf ./build