package main

import (
	"jogo_velha/jogo"
	"jogo_velha/model"
	"jogo_velha/view"
	"os"
)

var jogador0 model.Jogador
var jogador1 model.Jogador

func init() {
	jogador0.Nome, jogador0.Simbolo = "Jogador nº 1", model.X
	jogador1.Nome, jogador1.Simbolo = "Jogador nº 2", model.O

	model.AddJogador(jogador0)
	model.AddJogador(jogador1)
}

func main() {
	defineModoJogo()
}

func defineModoJogo() {
	for {
		op := view.MostraMenuInicial()

		switch op {
		case 1:
			jogo.JogaJogadorVsJogador()

			view.EnterParaContinuar()
		case 2:

			view.EnterParaContinuar()
		case 3:
			view.ExibePontuacaoGeral()

			view.EnterParaContinuar()
		default:
			view.FimDeJogo()
			os.Exit(0)
		}
	}
}
