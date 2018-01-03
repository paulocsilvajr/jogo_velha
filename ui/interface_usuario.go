package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var opcoes = []int{0, 1, 2, 3}

func MostraMenuInicial() int {
	for {
		LimpaTela()

		fmt.Println("Bem Vindo!")
		fmt.Println(opcoes[1], "- Jogador x Jogador")
		fmt.Println(opcoes[2], "- Jogador x Computador")
		fmt.Println(opcoes[3], "- Pontuação")
		fmt.Println(opcoes[0], "- Sair")
		fmt.Print("\t\t: ")

		var entrada string
		fmt.Scan(&entrada)
		opcao, err := strconv.Atoi(entrada) // converte entrada(string) para int
		if err != nil {
			fmt.Println("Escolha uma das opções do menu")
			Espere()
			EnterParaContinuar()

			continue
		}

		for i := range opcoes {
			if opcao == opcoes[i] {
				return opcao
			}
		}

		fmt.Printf("Opção informada [%d] inválida\n", opcao)
		Espere()
		EnterParaContinuar()
	}
}

func DefineModoJogo() {
	for {
		op := MostraMenuInicial()

		switch op {
		case 1:
			fmt.Println("Jogador x Jogador")
			tab.Imprime()

			EnterParaContinuar()
		case 2:
			fmt.Println("Jogador x Computador")

			EnterParaContinuar()
		default:
			FimDeJogo()
			os.Exit(0)
		}
	}
}

func ExibePontuacao(jogador *Jogador) {
	fmt.Println(jogador.nome)
	for i, vitoria := range jogador.pontuacao {
		fmt.Printf("Partida %d: ", i+1)
		if vitoria {
			fmt.Println("Vitória")
		} else {
			fmt.Println("Derrota")
		}
	}
}

func LimpaTela() {
	sistema := runtime.GOOS // retorna linux, windows, darwin, ...

	var comando string
	if sistema == "linux" {
		comando = "clear"
	} else if sistema == "windows" {
		comando = "cls"
	} else {
		panic("Plataforma não suportada para limpar a tela")
	}

	cmd := exec.Command(comando)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Espere() {
	time.Sleep(time.Second)
}

func EnterParaContinuar() {
	fmt.Print("\nENTER para continuar ")
	var temp string
	fmt.Scanf("%s", &temp)
}

func FimDeJogo() {
	fmt.Println("END OF LINE.")
}
