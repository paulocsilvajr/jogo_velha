package view

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/paulocsilvajr/jogo_velha/model"
)

var opcoes = []int{0, 1, 2, 3}

// MostraMenuInicial retorna a opção do usuário de acordo com sua escolha.
// A variável global opcoes possui a numeração para as opções da listagem.
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

// ExibePontuacao exibe visualmente a pontuação do jogador informado.
func ExibePontuacao(jogador *model.Jogador) {
	var contV, contD int
	nome := strings.ToUpper(jogador.Nome)

	if len(jogador.Pontuacao) > 0 {
		fmt.Println(nome)

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
	} else {
		fmt.Printf("%s: Nenhum registro de pontuação\n\n", nome)
	}
}

// ExibePontuacaoGeral invoca a função ExibePontuacao para cada jogador cadastrado.
func ExibePontuacaoGeral() {
	jogadores := model.GetJogadores()
	for _, jogador := range jogadores {
		ExibePontuacao(jogador)
	}
}

// EscolhaPosicao define a interface para captura da linha e coluna para jogadores humanos.
// Faz validação de acordo com o tamanho do tabuleiro.
// Retorna a linha e coluna informada.
func EscolhaPosicao(jogador *model.Jogador) (linha, coluna int) {
	for {
		fmt.Printf("%s [ %s ], informe uma \nlinha:  ",
			strings.ToUpper(jogador.Nome),
			model.GetSimbolo(model.Elemento(jogador.Simbolo)))
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

// EscolhaComputador define a interface visual para exibir a escolha do computador.
// O parâmetro tipo é utilizado para imprimir se a jogada foi aleatória(R) ou de acordo com algoritmo de IA.
func EscolhaComputador(jogador *model.Jogador, linha int, coluna int, tipo string) {
	linha++
	coluna++
	fmt.Printf("::%s:: %s [ %s ], informou uma \nlinha:  %d\ncoluna: %d\n\n",
		tipo,
		strings.ToUpper(jogador.Nome),
		model.GetSimbolo(model.Elemento(jogador.Simbolo)),
		linha,
		coluna)
}

// ImprimeTabuleiro imprime o retorno da função tabuleiro.Imprime().
func ImprimeTabuleiro(tabuleiro *model.Tabuleiro) {
	fmt.Println(tabuleiro.Imprime())
}

// PosicaoOcupada imprime a mensagem para posição ocupada.
func PosicaoOcupada() {
	fmt.Println("Posição ocupada\n")
}

// FinalizadoTabuleiro imprime mensagem quando o jogo dá velha.
func FinalizadoTabuleiro() {
	LimpaTela()
	fmt.Println("Deu velha\n")
}

// Vitoria imprime mensagem de vitória.
func Vitoria(jogador *model.Jogador) {
	LimpaTela()
	fmt.Println("Vitória de", jogador.Nome, "\n")
}

// LimpaTela executa a limpeza da tela de acordo com sistema operacional.
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

// Espere é uma função para generalizar a espera entre tarefas.
func Espere() {
	time.Sleep(time.Second)
}

// EnterParaContinuar define mensagem e captura do ENTER para pausar o programa quando necessário.
func EnterParaContinuar() {
	fmt.Print("\nENTER para continuar ")
	var temp string
	fmt.Scanf("%s", &temp)
}

// FimDeJogo define a mensagem quando finaliza o jogo.
func FimDeJogo() {
	fmt.Println("END OF LINE.")
}
