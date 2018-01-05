package view

import (
	"fmt"
	"jogo_velha/model"
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

func ExibePontuacao(jogador *model.Jogador) {
	var contV, contD int

	fmt.Println(jogador.Nome)
	for i, vitoria := range jogador.Pontuacao {
		fmt.Printf("Partida %d: ", i+1)
		if vitoria {
			fmt.Println("Vitória")
			contV++
		} else {
			fmt.Println("Derrota")
			contD++
		}
	}
	fmt.Println("Vitórias:", contV, "\nDerrotas:", contD, "\n")
}

func ExibePontuacaoGeral() {
	jogadores := model.GetJogadores()
	for _, jogador := range jogadores {
		ExibePontuacao(&jogador)
	}
}

func EscolhaPosicao(jogador *model.Jogador) (linha, coluna int) {
	for {
		fmt.Printf("%s, informe uma \nlinha: ", jogador.Nome)
		fmt.Scanf("%d", &linha)
		fmt.Print("coluna: ")
		fmt.Scanf("%d", &coluna)

		linha--
		coluna--
		if linha >= 0 && linha < model.Q && coluna >= 0 && coluna < model.Q {
			fmt.Println()
			return
		} else {
			fmt.Println("\nLinha ou coluna inválida\n")
		}
	}
}

func ImprimeTabuleiro(tabuleiro *model.Tabuleiro) {
	fmt.Println(tabuleiro.Imprime())
}

func PosicaoOcupada() {
	fmt.Println("Posição ocupada\n")
}

func FinalizadoTabuleiro() {
	fmt.Println("Deu velha\n")
}

func Vitoria(jogador *model.Jogador) {
	fmt.Println("Vitória de", jogador.Nome, "\n")
}

func LimpaTela() {
	sistema := runtime.GOOS // retorna linux, windows, darwin, ...

	var comando string
	if sistema == "linux" {
		comando = "clear"
	} else if sistema == "windows" {
		comando = "cls"
	} else {
		panic("Plataforma não suportada para limpeza de tela")
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
