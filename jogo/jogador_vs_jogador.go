package jogo

import (
	"jogo_velha/ui"
	"math"
)

var tabuleiro *ui.Tabuleiro

func init() {
	tabuleiro = ui.GetTabuleiro()
}

func JogaJogadorVsJogador() {
	maxJogadas := int(math.Pow(ui.Q, 2))
	jogadorAtual := ui.GetJogador(0)
	proximoJogador := ui.GetJogador(1)

	for i := 0; i < maxJogadas; i++ {
		ui.LimpaTela()
		tabuleiro.Imprime()

		for {
			linha, coluna := jogadorAtual.EscolhaPosicao()

			if tabuleiro.MarcaPosicao(jogadorAtual, linha, coluna) {
				break
			} else {
				ui.PosicaoOcupada()
			}
		}

		jogadorAtual, proximoJogador = proximoJogador, jogadorAtual

		tabuleiro.Imprime()

		ui.Espere()
		ui.EnterParaContinuar()
	}

	ui.FinalizadoTabuleiro()

}
