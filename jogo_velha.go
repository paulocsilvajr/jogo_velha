package main

import (
	"os"

	"github.com/paulocsilvajr/jogo_velha/jogo"
	"github.com/paulocsilvajr/jogo_velha/model"
	"github.com/paulocsilvajr/jogo_velha/view"
)

var jogador0 model.Jogador
var jogador1 model.Jogador

var partida int

func init() {
	jogador0.Nome, jogador0.Simbolo, jogador0.Humano = "Jogador nº 1", model.X, true
	jogador1.Nome, jogador1.Simbolo = "Douglas Adams", model.O

	model.AddJogador(&jogador0)
	model.AddJogador(&jogador1)
}

// defineModoJogo é a função principal do jogo da velha.
// A partir dela é possivel capturar a escolha de modo de jogo e
// inicializar o devido tipo de partida ou a exibição do pontuação
// acumulada do ultimo tipo selecionado.
func defineModoJogo() {
	for {
		op := view.MostraMenuInicial()

		switch op {
		case 1:
			jogada(op, "Jogador nº 2", true, jogo.JogaJogadorVsJogador)
		case 2:
			jogada(op, "Computador", false, jogo.JogaJogadorVsComputador)
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

func jogada(op int, nomeJogador1 string, ehHumano bool, joga func()) {
	resetaPontuacao(op)

	partida = op

	jogador1.Nome = nomeJogador1
	jogador1.Humano = ehHumano

	joga()

	view.EnterParaContinuar()
}

func resetaPontuacao(partidaEscolhida int) {
	if partidaEscolhida != partida {
		jogador0.ResetaPontuacao()
		jogador1.ResetaPontuacao()
	}
}
