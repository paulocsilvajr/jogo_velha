package jogo

import (
	"jogo_velha/model"
	"jogo_velha/view"
	"math"
)

var tabuleiro *model.Tabuleiro

func init() {
	tabuleiro = model.GetTabuleiro()
}

func JogaJogadorVsJogador() {
	// TESTE
	// c, p := tabuleiro.GetPosicoes()
	// for i := 0; i < len(p); i++ {
	// 	fmt.Println(c[i], p[i])
	// }
	// TESTE

	maxJogadas := int(math.Pow(model.Q, 2))
	jogadorAtual := model.GetJogador(0)
	proximoJogador := model.GetJogador(1)

	for i := 0; i < maxJogadas; i++ {
		view.LimpaTela()
		view.ImprimeTabuleiro(tabuleiro)

		for {
			linha, coluna := view.EscolhaPosicao(jogadorAtual)

			if tabuleiro.MarcaPosicao(jogadorAtual, linha, coluna) {
				break
			} else {
				view.PosicaoOcupada()
			}
		}

		jogadorAtual, proximoJogador = proximoJogador, jogadorAtual

		tabuleiro.Imprime()

		// view.Espere()
		view.EnterParaContinuar()
	}

	view.FinalizadoTabuleiro()

}
