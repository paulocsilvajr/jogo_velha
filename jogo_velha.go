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

func defineModoJogo() {
	for {
		op := view.MostraMenuInicial()

		switch op {
		case 1:
			jogada(op, "Jogador nº 2", true)
		case 2:
			jogada(op, "Computador", false)
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

func jogada(op int, nomeJogador1 string, ehHumano bool) {
	resetaPontuacao(op)

	partida = op

	jogador1.Nome = nomeJogador1
	jogador1.Humano = ehHumano

	jogo.JogaJogadorVsComputador()

	view.EnterParaContinuar()
}

func resetaPontuacao(partidaEscolhida int) {
	if partidaEscolhida != partida {
		jogador0.ResetaPontuacao()
		jogador1.ResetaPontuacao()
	}
}
