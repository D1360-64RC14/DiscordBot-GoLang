OUT_FILE_NAME = discord-bot_byD1360
FILESgo       = main.go command/commands.go command/commandController.go utils/utils.go config/config.go config/dataStructs.go youtube/youtube.go youtube/dataStructs.go
FILESetc      = config.yml go.mod

# [ Flags ] (go help build)
# -v    print the names of packages as they are compiled.
# -x    print the commands.
BUILD_FLAGS = -v -x

# Com flag ( -v | --verbose ) exibe
# log de mensagens no terminal
run : $(FILESgo)
	@ echo "=== Rodando progrma sem compilar ==="
	go run main.go -v
	@ echo ""

build : $(FILESgo) $(FILESetc)
	@ echo "=== Criando diretório build/ ==="
	mkdir build
	@ echo "=== Compilando programa ==="
	go build -o ./build/$(OUT_FILE_NAME) $(BUILD_FLAGS)
	cp config.yml ./build/config.yml
	@ echo ""
	@ echo "=== Done ==="

buildrun : $(FILESgo) $(FILESetc)
	@ echo "=== Compilando e rodando o programa ==="
	go build -o ./build/$(OUT_FILE_NAME) $(BUILD_FLAGS)
	cp config.yml ./build/config.yml
	@ echo ""
	@ echo "=== Done ==="
	@ echo "=== Executando programa ==="
	./build/$(OUT_FILE_NAME)

clean : 
	@ echo "=== Limpando build/ ==="
	rm -rf ./build