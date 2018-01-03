package main

import (
	"jogo_velha/jogo"
	"jogo_velha/ui"
	"os"
)

func main() {
	defineModoJogo()
}

func defineModoJogo() {
	for {
		op := ui.MostraMenuInicial()

		switch op {
		case 1:
			jogo.JogaJogadorVsJogador()

			ui.EnterParaContinuar()
		case 2:

			ui.EnterParaContinuar()
		default:
			ui.FimDeJogo()
			os.Exit(0)
		}
	}
}
