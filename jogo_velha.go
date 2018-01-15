package main

import (
	"jogo_velha/jogo"
	"jogo_velha/model"
	"jogo_velha/view"
	"os"
)

var jogador0 model.Jogador
var jogador1 model.Jogador

var partida int

func init() {
	jogador0.Nome, jogador0.Simbolo, jogador0.Humano = "Jogador nº 1", model.X, true
	jogador1.Simbolo = model.O

	model.AddJogador(&jogador0)
	model.AddJogador(&jogador1)
}

func main() {
	defineModoJogo()
}

// defineModoJogo é a função principal do jogo da velha.
// A partir dela é possivel capturar a escolha de modo de jogo e
// inicializar o devido processo.
func defineModoJogo() {
	for {
		op := view.MostraMenuInicial()

		switch op {
		case 1:
			jogada(op, "Jogador nº 2", true, jogo.JogaJogadorVsJogador)
		case 2:
			jogada(op, "Computador", false, jogo.JogaJogadorVsComputador)
		case 3:
			view.LimpaTela()

			view.ExibePontuacaoGeral()

			view.EnterParaContinuar()
		default:
			view.FimDeJogo()
			os.Exit(0)
		}
	}
}

func jogada(op int, nomeJogador1 string, ehHumano bool, joga func()) {
	resetaPontuacao(op)

	partida = op

	jogador1.Nome = nomeJogador1
	jogador1.Humano = ehHumano

	// if op == 1 {
	// 	jogo.JogaJogadorVsJogador()
	// } else {
	// 	jogo.JogaJogadorVsComputador()
	// }
	joga()

	view.EnterParaContinuar()
}

func resetaPontuacao(partidaEscolhida int) {
	if partidaEscolhida != partida {
		jogador0.ResetaPontuacao()
		jogador1.ResetaPontuacao()
	}
}
