package jogo

import (
	"jogo_velha/model"
	"jogo_velha/view"
	"math/rand"
)

func init() {
	// inicialização de tabuleiro declarado em jogo.go(package jogo)
	tabuleiro = model.GetTabuleiro()
}

func JogaJogadorVsComputador() {
	Jogar(marcaPosicaoJogadorVsComputador)
}

func marcaPosicaoJogadorVsComputador(jogadorAtual *model.Jogador) {
	for {
		if jogadorAtual.EhComputador() {
			marcaJogadaComputador(jogadorAtual)
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

func marcaJogadaComputador(jogador *model.Jogador) {
	posicao := retornaPosicaoVazia()
	linha, coluna := posicao[0], posicao[1]

	view.EscolhaComputador(jogador, linha, coluna)
	//fmt.Printf("linha: %d\ncoluna:%d\n\n", linha, coluna)
	tabuleiro.MarcaPosicao(jogador, linha, coluna)
	view.Espere()
}

func retornaPosicaoVazia() model.Posicao {
	_, posicoes := tabuleiro.GetElementosVazios()
	posicaoAleatoria := rand.Intn(len(posicoes))
	return posicoes[posicaoAleatoria]
}
