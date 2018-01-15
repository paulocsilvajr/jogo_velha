package jogo

import (
	"fmt"
	"jogo_velha/model"
	"jogo_velha/view"
	"math/rand"
	"time"
)

func init() {
	// inicialização de tabuleiro declarado em jogo.go(package jogo)
	tabuleiro = model.GetTabuleiro()
}

func JogaJogadorVsComputador() {
	// torna rand.Intn realmente aleatório
	rand.Seed(time.Now().UTC().UnixNano())

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
	var posicao model.Posicao
	var tipo string

	if tipoJogada() {
		tipo = "IA" // Inteligência Artificial
		// Baseado no algoritmo apresentado no livro(versão kindle):
		// Domingos, Pedro. O algoritmo Mestre: Como a busca pelo algoritmo de machine learning definitivo recriará nosso mundo.
		// Capítulo: A revolução do machine learning, posição 349.

		// define posição para completar conjunto(linha, coluna, diagonal)
		if p, ok := getPosicaoParaVencer(jogador.Simbolo); ok {
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //
			//fmt.Println("Marcações em sequência")
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //

			posicao = p

			// define posição para evitar que o adversário complete conjunto
		} else if p, ok := getPosicaoParaVencer(jogador.GetSimboloAdversario()); ok {
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //
			//fmt.Println("Marcações em sequência adversário")
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //

			posicao = p

			// define posição central, se vazia
		} else if p, ok := getPosicaoCentral(); ok {
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //
			//fmt.Println("Marca o centro, se livre")
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //

			posicao = p

			// define posição do canto oposto do adversário, se vazia
		} else if p, ok := getPosicaoCantoOposto(jogador); ok {
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //
			//fmt.Println("Marca canto oposto ao adversário, se livre")
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //

			posicao = p

			// define posição como o canto vazio
		} else if p, ok := getPosicaoCantoVazio(); ok {
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //
			//fmt.Println("Marca um canto vazio")
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //

			posicao = p

			// define posição vazia aleatória
		} else {
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //
			//fmt.Println("Posição vazia aleatória")
			// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //

			posicao = getPosicaoVazia()
		}

	} else {
		tipo = "R" // Randômico

		// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //
		fmt.Println("Posição vazia aleatória")
		// APAGAR // APAGAR // APAGAR // APAGAR // APAGAR //

		posicao = getPosicaoVazia()
	}

	linha, coluna := posicao[0], posicao[1]

	view.EscolhaComputador(jogador, linha, coluna, tipo)
	tabuleiro.MarcaPosicao(jogador, linha, coluna)
	view.Espere()
}

func tipoJogada() bool {
	// DESCOMENTAR // CÓDIGO ORIGINAL //
	return rand.Intn(2) != 0

	// TESTE // APAGAR // TESTE // APAGAR //
	//return true
	// TESTE // APAGAR // TESTE // APAGAR //
}

func getPosicaoVazia() model.Posicao {
	_, posicoes := tabuleiro.GetElementosVazios()
	posicaoAleatoria := rand.Intn(len(posicoes))
	return posicoes[posicaoAleatoria]
}

func getPosicaoParaVencer(simbolo model.Elemento) (posicao model.Posicao, temPosicao bool) {
	conjuntos, posicoes := tabuleiro.GetPosicoes()
	temPosicao = false

	for i := range conjuntos {
		contSimbolo := 0
		contVazio := 0
		for j, elemento := range conjuntos[i] {
			switch elemento {
			case simbolo:
				contSimbolo++
			case model.Vazio:
				contVazio++

				posicao = posicoes[i][j]
			default:
				// se achou o simbolo oposto
				// para a verificação
				break
			}
		}

		if contSimbolo == 2 && contVazio == 1 {
			temPosicao = true
			return
		}
	}

	return
}

func getPosicaoCentral() (posicao model.Posicao, temPosicao bool) {
	elementos, posicoes := tabuleiro.GetElementos()

	// posição 4 é a linha 2 coluna 2, portanto o centro
	centro := 4
	if elementos[centro] == model.Vazio {
		posicao = posicoes[centro]
		temPosicao = true
		return
	}

	temPosicao = false
	return
}

func getPosicaoCantoOposto(jogador *model.Jogador) (posicao model.Posicao, temPosicao bool) {
	diagonalPrincipal, posicoesPrincipal := tabuleiro.GetDiagonal(0)
	diagonalSecundaria, posicoesSecundaria := tabuleiro.GetDiagonal(1)

	simboloAdversario := jogador.GetSimboloAdversario()

	diagonais := []model.Conjunto{diagonalPrincipal, diagonalSecundaria}
	posicoes := []model.PosRelativa{posicoesPrincipal, posicoesSecundaria}
	for i := range diagonais {
		if diagonais[i][0] == model.Vazio && diagonais[i][2] == simboloAdversario {
			posicao = posicoes[i][0]
			temPosicao = true
			return
		} else if diagonais[i][2] == model.Vazio && diagonais[i][0] == simboloAdversario {
			posicao = posicoes[i][2]
			temPosicao = true
			return
		}

	}

	temPosicao = false
	return
}

func getPosicaoCantoVazio() (posicao model.Posicao, temPosicao bool) {
	elementos, posicoes := tabuleiro.GetElementos()

	cantos := []int{0, 2, 6, 8}
	for _, canto := range cantos {
		if elementos[canto] == model.Vazio {
			posicao = posicoes[canto]
			temPosicao = true
			return
		}
	}

	temPosicao = false
	return
}
