package jogo

import (
	"jogo_velha/model"
	"jogo_velha/view"
)

func init() {
	// inicialização de tabuleiro declarado em jogo.go(package jogo)
	tabuleiro = model.GetTabuleiro()
}

func JogaJogadorVsJogador() {
	Jogar(marcaPosicaoJogadorVsJogador)
}

func marcaPosicaoJogadorVsJogador(jogadorAtual *model.Jogador) {
	for {
		linha, coluna := view.EscolhaPosicao(jogadorAtual)

		if tabuleiro.MarcaPosicao(jogadorAtual, linha, coluna) {
			break
		} else {
			view.PosicaoOcupada()
		}
	}
}
