package ui

import "fmt"

const (
	Vazio = 0
	X     = 1
	O     = 2
	Q     = 3 // Quantidade de elemento no tabuleiro(3x3)
)

var I = map[int]string{0: " ", 1: "X", 2: "O"} // Representação visual de Vazio, X e O

type Tabuleiro [Q][Q]int   // Tipo Tabuleiro, matrix de 3x3
type PosRelativa [Q][2]int // Posição relativa para linhas, colunas e diagonais

type Jogador struct {
	nome      string
	simbolo   int
	pontuacao []bool
}

// Variáveis não são acessíveis fora do seu pacote
// portanto, foi criado gets retornando ponteiros
var tab Tabuleiro
var jogador0 Jogador
var jogador1 Jogador

func init() {
	tab = [Q][Q]int{
		[Q]int{Vazio, Vazio, Vazio},
		[Q]int{Vazio, Vazio, Vazio},
		[Q]int{Vazio, Vazio, Vazio},
	}

	jogador0.nome, jogador0.simbolo = "Jogador nº 1", X
	jogador1.nome, jogador1.simbolo = "Computador", O
}

func GetTabuleiro() *Tabuleiro {
	return &tab
}

func GetJogador(numero int) *Jogador {
	if numero == 1 {
		return &jogador1
	}
	return &jogador0
}

func (jogador *Jogador) SetPontuacao(vitoria bool) {
	jogador.pontuacao = append(jogador.pontuacao, vitoria)
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

func (tabuleiro *Tabuleiro) MarcaPosicao(jogador *Jogador, linha int, coluna int) bool {
	if tabuleiro[linha][coluna] == Vazio {
		tabuleiro[linha][coluna] = jogador.simbolo
		return true
	}
	return false
}
