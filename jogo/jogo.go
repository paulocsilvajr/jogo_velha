package jogo

import (
	"jogo_velha/model"
	"jogo_velha/view"
	"math"
)

var tabuleiro *model.Tabuleiro
var primeiro, segundo int = 0, 1

func Jogar(jogada_esp func(jogador *model.Jogador)) {
	maxJogadas := int(math.Pow(model.Q, 2))
	jogadorAtual := model.GetJogador(primeiro)
	proximoJogador := model.GetJogador(segundo)
	model.ZeraTabuleiro()

	var jogador *model.Jogador
	for i := 0; i < maxJogadas; i++ {
		view.LimpaTela()
		view.ImprimeTabuleiro(tabuleiro)

		// função específica para cada tipo de partida:
		// jogadorVsJogador ou jogadorVsComputador
		jogada_esp(jogadorAtual)

		// verificação de ocorreu vitória, caso afirmativo
		// marca pontuação para jogadores
		jogador = tabuleiro.Vitoria()
		if jogador != nil {
			jogadorAtual.SetPontuacao(true)
			proximoJogador.SetPontuacao(false)

			break
		}

		jogadorAtual, proximoJogador = proximoJogador, jogadorAtual

		view.ImprimeTabuleiro(tabuleiro)

		// view.Espere()
		view.EnterParaContinuar()
	}

	if jogador != nil {
		view.Vitoria(jogador)
	} else {
		view.FinalizadoTabuleiro()
	}

	primeiro, segundo = segundo, primeiro

	view.ExibePontuacao(jogadorAtual)
	view.ExibePontuacao(proximoJogador)
}
