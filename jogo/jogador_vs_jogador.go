package jogo

import (
	"jogo_velha/model"
	"jogo_velha/view"
)

var tabuleiro *model.Tabuleiro
var primeiro, segundo int = 0, 1

func init() {
	tabuleiro = model.GetTabuleiro()
}

func JogaJogadorVsJogador() {
	Jogar(marcaPosicaoJogador)
}

func marcaPosicaoJogador(jogadorAtual *model.Jogador) {
	for {
		linha, coluna := view.EscolhaPosicao(jogadorAtual)

		if tabuleiro.MarcaPosicao(jogadorAtual, linha, coluna) {
			break
		} else {
			view.PosicaoOcupada()
		}
	}
}

// func JogaJogadorVsJogador() {
// 	maxJogadas := int(math.Pow(model.Q, 2))
// 	jogadorAtual := model.GetJogador(primeiro)
// 	proximoJogador := model.GetJogador(segundo)
// 	model.ZeraTabuleiro()
//
// 	var jogador *model.Jogador
// 	for i := 0; i < maxJogadas; i++ {
// 		view.LimpaTela()
// 		view.ImprimeTabuleiro(tabuleiro)
//
// 		for {
// 			linha, coluna := view.EscolhaPosicao(jogadorAtual)
//
// 			if tabuleiro.MarcaPosicao(jogadorAtual, linha, coluna) {
// 				break
// 			} else {
// 				view.PosicaoOcupada()
// 			}
// 		}
//
// 		jogador = tabuleiro.Vitoria()
// 		if jogador != nil {
// 			jogadorAtual.SetPontuacao(true)
// 			proximoJogador.SetPontuacao(false)
//
// 			break
// 		}
//
// 		jogadorAtual, proximoJogador = proximoJogador, jogadorAtual
//
// 		view.ImprimeTabuleiro(tabuleiro)
//
// 		// view.Espere()
// 		view.EnterParaContinuar()
// 	}
//
// 	if jogador != nil {
// 		view.Vitoria(jogador)
// 	} else {
// 		view.FinalizadoTabuleiro()
// 	}
//
// 	primeiro, segundo = segundo, primeiro
//
// 	view.ExibePontuacao(jogadorAtual)
// 	view.ExibePontuacao(proximoJogador)
// }
