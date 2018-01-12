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
			resetaPontuacao(2)

			partida = op

			jogador1.Nome = "Jogador nº 2"

			jogo.JogaJogadorVsJogador()

			view.EnterParaContinuar()
		case 2:
			resetaPontuacao(1)

			partida = op

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

func resetaPontuacao(partida int) {
	if partida == 1 {
		jogador0.ResetaPontuacao()
		jogador1.ResetaPontuacao()
	}
}
