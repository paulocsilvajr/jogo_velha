package jogo

import (
	"fmt"
	"jogo_velha/model"
	"jogo_velha/view"
)

// var tabuleiro *model.Tabuleiro
// var primeiro, segundo int = 0, 1

func init() {
	tabuleiro = model.GetTabuleiro()
}

func JogaJogadorVsComputador() {
	Jogar(marcaPosicaoJogadorVsComputador)
}

func marcaPosicaoJogadorVsComputador(jogadorAtual *model.Jogador) {
	for {
		if jogadorAtual.EhComputador() {
			fmt.Println("Jogada computador")
			break
		} else {
			linha, coluna := view.EscolhaPosicao(jogadorAtual)

			if tabuleiro.MarcaPosicao(jogadorAtual, linha, coluna) {
				break
			} else {
				view.PosicaoOcupada()
			}
		}
	}
}
