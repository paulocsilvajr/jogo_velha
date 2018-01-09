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
	jogador0.Nome, jogador0.Simbolo = "Jogador nº 1", model.X
	jogador1.Nome, jogador1.Simbolo = "Jogador nº 2", model.O

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
			if partida == 2 {
				jogador0.ResetaPontuacao()
				jogador1.ResetaPontuacao()
			}

			partida = 1

			jogador1.Nome = "Jogador nº 2"

			jogo.JogaJogadorVsJogador()

			view.EnterParaContinuar()
		case 2:
			if partida == 1 {
				jogador0.ResetaPontuacao()
				jogador1.ResetaPontuacao()
			}

			partida = 2

			jogador1.Nome = "Computador"

			view.EnterParaContinuar()
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
