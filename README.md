<style>
    .this {

    }

    .is {
        opacity: 0;
    }

    .extreme {
        opacity: 0;
        display: none;
    }

    .go {

    }

    .horse {
        opacity: 0;
        display: none;
    }

    .extremeGoHorse:hover > .horse, .extremeGoHorse:hover > .extreme {
        display: unset;
        opacity: 1;
    }

    .this:hover > .is {
        opacity: .3;
    }

    .gif, .gif img {
        height: 200px;
        justify-content: start;
        border-radius: 4px;
        background-color: unset;
        transition: all 500ms;
    }
    .gif:hover > img, .gif:hover {
        height: 303px;
        background-color: #35383f;
        height: 303px;
        border-radius: 10px;
    }
    .gif:hover > img {
        transform: translate(25%);
    }
</style>

Um chatbot pro Discord feito em <span class="extremeGoHorse"><span class="extreme">**Extreme** </span><span class="go">**GO**</span><span class="horse"> **Horse**</span></span>
## Instalação
0. Clone o repositório;
0. Instale o GO;
    - [Windows](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-windows-10-pt)
    - [Linux](https://www.youtube.com/watch?v=dQw4w9WgXcQ)
0. <span class="this">Vá para a pasta do repositório<span class="is">, que provavelmente tu já sabe como fazer :D</span></span>
0. Coloque o [token do Discrd Bot](https://discord.com/developers/applications) dentro do arquivo `token`;
0. Instale a lib [discordgo](https://github.com/bwmarrin/discordgo), executando `go get github.com/bwmarrin/discordgo` aí no terminal;
0. Compile o programa executando `go build main.go` no terminal;
0. E se alguma coisa der errado você não é o escolhido.

## Comandos
- `!user <usuário>`

<div class="gif" style="display: flex;">
    <img src=".github/README/commandexample-user.gif">
</div>
<br>
Outros comandos serão adicionados quando eu tiver criatividade pra tal.