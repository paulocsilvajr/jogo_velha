package model

import "fmt"

// Elemento de uma posição do tabuleiro
type Elemento int

// Posicao de um elemento [linha, coluna]
type Posicao [2]int

// Tabuleiro matrix de 3x3
type Tabuleiro [Q][Q]Elemento

// PosRelativa  são as posições relativas para linhas, colunas e diagonais
type PosRelativa [Q]Posicao

// Conjunto representando linha, coluna ou diagonal
type Conjunto [Q]Elemento

// Conjuntos é um grupo de conjuntos para agrupar todas as linhas, colunas e diagonal
type Conjuntos []Conjunto

// Posicoes é um grupo de Posições relativas ao grupo de conjuntos
type Posicoes []PosRelativa

// ElementosSimples são conjuntos de todos os elementos do tabuleiro
type ElementosSimples [Q * Q]Elemento

// PosicoesSimples são posicões de todos os elementos do tabuleiro
type PosicoesSimples [Q * Q]Posicao

// Vazio representa a casa vazia no tabuleiro.
// X representa a casa marcada com X(xis).
// O representa a casa marcada com O(bola).
// Q representa a quantidade de elementos no tabuleiro(3x3)
const (
	Vazio = 0
	X     = 1
	O     = 2
	Q     = 3
)

// I representa a forma visual de Vazio, X e O
var I = map[int]string{Vazio: " ", X: "x", O: "o"}

// Variáveis iniciadas com minúsculo não são acessíveis fora do seu pacote
// portanto, foi criado gets retornando ponteiros.
var tab Tabuleiro

func init() {
	ZeraTabuleiro()
}

// ZeraTabuleiro atribui em cada posição do tabuleiro VAZIO.
func ZeraTabuleiro() {
	tab = [Q][Q]Elemento{
		[Q]Elemento{Vazio, Vazio, Vazio},
		[Q]Elemento{Vazio, Vazio, Vazio},
		[Q]Elemento{Vazio, Vazio, Vazio},
	}
}

// GetTabuleiro retorna um ponteiro de tabuleiro referente a variável local tab.
func GetTabuleiro() *Tabuleiro {
	return &tab
}

// Imprime retorna uma string contendo a forma de exibição do tabuleiro.
func (tabuleiro *Tabuleiro) Imprime() (impressao string) {
	numHorizontais := "    1   2   3"
	linhaSeparadora := "  +---+---+---+"
	impressao = fmt.Sprintf("%s\n", numHorizontais)
	impressao += fmt.Sprintf("%s\n", linhaSeparadora)

	for numVerticais, linha := range tabuleiro {
		impressao += fmt.Sprintf("%d ", numVerticais+1)
		for _, coluna := range linha {
			impressao += fmt.Sprintf("| %s ", GetSimbolo(coluna))
		}
		impressao += fmt.Sprintf("|\n%s\n", linhaSeparadora)
	}

	return
}

// GetSimbolo retorna uma string de acordo com o elemento da casa informado.
func GetSimbolo(valor Elemento) string {
	switch valor {
	case O:
		return I[int(valor)]
	case X:
		return I[int(valor)]
	default:
		return I[int(0)]
	}
}

// GetLinha retorna a linha e as posiçoes relativas referente ao parâmetro posicao informado.
func (tabuleiro *Tabuleiro) GetLinha(posicao int) (Conjunto, PosRelativa) {
	var posicoes PosRelativa
	for i := 0; i < Q; i++ {
		posicoes[i] = Posicao{posicao, i}
	}

	return tabuleiro[posicao], posicoes
}

// GetColuna retorna a coluna e as posiçoes relativas referente ao parâmetro posicao informado.
func (tabuleiro *Tabuleiro) GetColuna(posicao int) (Conjunto, PosRelativa) {
	var colunaTemp [Q]Elemento

	var posicoes PosRelativa
	for i := 0; i < Q; i++ {
		posicoes[i] = Posicao{i, posicao}
	}

	for i, linha := range tabuleiro {
		for j, coluna := range linha {
			if j == posicao {
				colunaTemp[i] = coluna
			}
		}
	}

	return colunaTemp, posicoes
}

// GetDiagonal retorna a diagonal e as posiçoes relativas referente ao parâmetro posicao informado.
func (tabuleiro *Tabuleiro) GetDiagonal(posicao int) (Conjunto, PosRelativa) {
	var diagonalTemp [Q]Elemento

	var posicoes PosRelativa

	if posicao == 1 {
		// diagonal secundária
		posicao = Q
		for i, linha := range tabuleiro {
			posicao--
			diagonalTemp[i] = linha[posicao]
		}

		col := Q
		for i := 0; i <= 2; i++ {
			col--
			posicoes[i] = Posicao{i, col}
		}

	} else {
		// diagonal principal
		posicao = 0
		for i, linha := range tabuleiro {
			diagonalTemp[i] = linha[posicao]
			posicao++
		}

		for i := 0; i < Q; i++ {
			posicoes[i] = Posicao{i, i}
		}
	}
	return diagonalTemp, posicoes
}

// MarcaPosicao marca a linha e coluna informada com o símbolo referente ao jogador informado.
// Retorna true se a casa estiver vazia e seja efetuado a atribuição.
func (tabuleiro *Tabuleiro) MarcaPosicao(jogador *Jogador, linha int, coluna int) bool {
	if tabuleiro[linha][coluna] == Vazio {
		tabuleiro[linha][coluna] = Elemento(jogador.Simbolo)
		return true
	}
	return false
}

// GetPosicoes Retorna as linhas, colunas e diagonais juntamente com slice contendo as posições relativas(linha, coluna).
// Retorno da função são dois slices, portanto é necessário atribuir a duas variáveis para utilizar seu retorno.
func (tabuleiro *Tabuleiro) GetPosicoes() (conjuntos Conjuntos, posicoes Posicoes) {
	for i := 0; i < Q; i++ {
		linha, pos := tabuleiro.GetLinha(i)
		conjuntos = append(conjuntos, linha)
		posicoes = append(posicoes, pos)
	}

	for i := 0; i < Q; i++ {
		coluna, pos := tabuleiro.GetColuna(i)
		conjuntos = append(conjuntos, coluna)
		posicoes = append(posicoes, pos)
	}

	for i := 0; i < 2; i++ {
		diagonal, pos := tabuleiro.GetDiagonal(i)
		conjuntos = append(conjuntos, diagonal)
		posicoes = append(posicoes, pos)
	}

	return
}

// GetElementos retorna todos os elementos do tabuleiro em um array(elementos) de 9 posições, juntamente com
// outro array(posicoes) com tamanho semelhante, mas contendo em cada posição as posições relativas(linha, coluna) de cada elemento
func (tabuleiro *Tabuleiro) GetElementos() (elementos ElementosSimples, posicoes PosicoesSimples) {
	cont := 0
	for i := range tabuleiro {
		for j := range tabuleiro[i] {
			elementos[cont] = tabuleiro[i][j]
			posicoes[cont] = Posicao{i, j}
			cont++
		}
	}
	return
}

// GetElementosVazios retorna os elementos e as posições atribuidos com vazio.
func (tabuleiro *Tabuleiro) GetElementosVazios() (elementos []Elemento, posicoes []Posicao) {
	e, p := tabuleiro.GetElementos()
	length := len(e)
	for i := 0; i < length; i++ {
		if e[i] == Vazio {
			elementos = append(elementos, e[i])
			posicoes = append(posicoes, p[i])
		}
	}
	return
}

// Vitoria retorna o jogador que completo uma linha, coluna ou diagonal.
func (tabuleiro *Tabuleiro) Vitoria() (jogador *Jogador) {
	conjuntos, _ := tabuleiro.GetPosicoes()

	for i := 0; i < len(conjuntos); i++ {
		x, o := 0, 0

		conjunto := conjuntos[i]
		for j := 0; j < len(conjunto); j++ {
			switch conjunto[j] {
			case X:
				x++
			case O:
				o++
			}

			if x == 3 {
				jogador = GetJogadorPorSimbolo(X)
				return
			} else if o == 3 {
				jogador = GetJogadorPorSimbolo(O)
				return
			}
		}

	}

	return nil
}
