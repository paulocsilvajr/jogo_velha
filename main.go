package main

import (
	"jogo_velha/jogo"
	"jogo_velha/view"
	"os"
)

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
		default:
			view.FimDeJogo()
			os.Exit(0)
		}
	}
}
