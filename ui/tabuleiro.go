package ui

import "fmt"

const (
	Vazio = 0
	X     = 1
	O     = 2
	Q     = 3 // Quantidade de elemento no tabuleiro(3x3)
)

var I = map[int]string{0: " ", 1: "O", 2: "X"} // Representação visual de Vazio, O e X

type Tabuleiro [Q][Q]int   // Tipo Tabuleiro, matrix de 3x3
type PosRelativa [Q][2]int // Posição relativa para linhas, colunas e diagonais

var tab Tabuleiro

func init() {
	tab = [Q][Q]int{
		[Q]int{Vazio, Vazio, Vazio},
		[Q]int{Vazio, Vazio, Vazio},
		[Q]int{Vazio, Vazio, Vazio},
	}
}

func (tabuleiro *Tabuleiro) Imprime() {
	numHorizontais := "    1   2   3"
	linhaSeparadora := "  +---+---+---+"
	fmt.Println(numHorizontais)
	fmt.Println(linhaSeparadora)

	for numVerticais, linha := range tabuleiro {
		fmt.Printf("%d ", numVerticais+1)
		for _, coluna := range linha {
			fmt.Printf("| %s ", GetSimbolo(coluna))
		}
		fmt.Println("|\n" + linhaSeparadora)
	}
}

func GetSimbolo(valor int) string {
	switch valor {
	case O:
		return I[valor]
	case X:
		return I[valor]
	default:
		return I[0]
	}
}

func (tabuleiro *Tabuleiro) GetLinha(posicao int) ([Q]int, PosRelativa) {
	var posicoes PosRelativa
	for i := 0; i < Q; i++ {
		posicoes[i] = [2]int{posicao, i}
	}

	return tabuleiro[posicao], posicoes
}

func (tabuleiro *Tabuleiro) GetColuna(posicao int) ([Q]int, PosRelativa) {
	var colunaTemp [Q]int

	var posicoes PosRelativa
	for i := 0; i < Q; i++ {
		posicoes[i] = [2]int{i, posicao}
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

func (tabuleiro *Tabuleiro) GetDiagonal(posicao int) ([Q]int, PosRelativa) {
	var diagonalTemp [Q]int

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
			posicoes[i] = [2]int{i, col}
		}

	} else {
		// diagonal principal
		posicao = 0
		for i, linha := range tabuleiro {
			diagonalTemp[i] = linha[posicao]
			posicao++
		}

		for i := 0; i < Q; i++ {
			posicoes[i] = [2]int{i, i}
		}
	}

	return diagonalTemp, posicoes
}
