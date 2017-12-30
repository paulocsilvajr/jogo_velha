package main

import "fmt"

const (
	Vazio        = 0
	O            = 1
	X            = 2
	LimiteLinha  = 3
	LimiteColuna = LimiteLinha
)

var I = map[int]string{0: " ", 1: "O", 2: "X"}

type ticTacToe struct {
	tabuleiro [LimiteLinha][LimiteColuna]int
}

func main() {
	// velha := makeTicTacToe(1)

	for i := 0; i < 24; i++ {
		velha := makeTeste(i)
		velha.imprime()

		s, pos := velha.getSimboloPosicaoMarcacoesEmSequencia()
		fmt.Printf("[%s] linha: %d, coluna: %d\n\n", getSimbolo(s), pos[0]+1, pos[1]+1)
	}
}

// FUNÇÕES RELACIONADAS A STRUCT tictacToe

func (self *ticTacToe) zera() {
	for i := 0; i < LimiteLinha; i++ {
		for j := 0; j < LimiteColuna; j++ {
			self.tabuleiro[i][j] = Vazio
		}
	}
}

func (self *ticTacToe) exemplo() {
	self.tabuleiro = [3][3]int{
		[3]int{X, Vazio, Vazio},
		[3]int{X, Vazio, Vazio},
		[3]int{Vazio, Vazio, Vazio},
	}
}

func makeTeste(numero int) (jogoVelha ticTacToe) {
	jogoVelha = ticTacToe{}
	switch numero {
	case 0:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, X, X},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 1:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{X, Vazio, X},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 2:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{X, X, Vazio},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 3:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, O, O},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 4:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{O, Vazio, O},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 5:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{O, O, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 6:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, X, X},
		}
	case 7:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{X, Vazio, X},
		}
	case 8:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{X, X, Vazio},
		}
	case 9:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{O, Vazio, Vazio},
			[3]int{O, Vazio, Vazio},
		}
	case 10:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{O, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{O, Vazio, Vazio},
		}
	case 11:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{O, Vazio, Vazio},
			[3]int{O, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 12:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, X, Vazio},
			[3]int{Vazio, X, Vazio},
		}
	case 13:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, X, Vazio},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, X, Vazio},
		}
	case 14:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, X, Vazio},
			[3]int{Vazio, X, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 15:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, Vazio, O},
		}
	case 16:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, O},
		}
	case 17:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 18:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, X, Vazio},
			[3]int{Vazio, Vazio, X},
		}
	case 19:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{X, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, X},
		}
	case 20:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{X, Vazio, Vazio},
			[3]int{Vazio, X, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 21:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, O, Vazio},
			[3]int{O, Vazio, Vazio},
		}
	case 22:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, Vazio, Vazio},
			[3]int{O, Vazio, Vazio},
		}
	case 23:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, O, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	}

	return

}

func makeTicTacToe(op int) ticTacToe {
	jogoVelha := ticTacToe{}
	if op == 1 {
		jogoVelha.exemplo()
	} else {
		jogoVelha.zera()
	}

	return jogoVelha
}

func (self *ticTacToe) imprime() {
	numHorizontais := "    1   2   3"
	linhaSeparadora := "  +---+---+---+"
	fmt.Println(numHorizontais)
	fmt.Println(linhaSeparadora)

	for numVerticais, linha := range self.tabuleiro {
		fmt.Printf("%d ", numVerticais+1)
		for _, coluna := range linha {
			fmt.Printf("| %s ", getSimbolo(coluna))
		}
		fmt.Println("|\n" + linhaSeparadora)
	}

}

func getSimbolo(valor int) string {
	switch valor {
	case O:
		return I[valor]
	case X:
		return I[valor]
	default:
		return I[0]
	}
}

func (self *ticTacToe) linha(posicao int) [LimiteLinha]int {
	posicaoValida(&posicao)
	return self.tabuleiro[posicao]
}

func (self *ticTacToe) coluna(posicao int) [LimiteColuna]int {
	posicaoValida(&posicao)
	var colunaTemp [LimiteColuna]int

	for i, linha := range self.tabuleiro {
		for j, coluna := range linha {
			if j == posicao {
				colunaTemp[i] = coluna
			}
		}
	}

	return colunaTemp
}

func posicaoValida(posicao *int) {
	if *posicao >= LimiteLinha {
		*posicao = 0
	} else if *posicao < 0 {
		*posicao = 0
	}
}

func (self *ticTacToe) diagonal(posicao int) [LimiteLinha]int {
	var diagonalTemp [LimiteLinha]int

	if posicao == 1 {
		// diagonal secundária
		posicao = LimiteColuna
		for i, linha := range self.tabuleiro {
			posicao--
			diagonalTemp[i] = linha[posicao]
		}
	} else {
		// diagonal principal
		posicao = 0
		for i, linha := range self.tabuleiro {
			diagonalTemp[i] = linha[posicao]
			posicao++
		}
	}

	return diagonalTemp
}

func (self *ticTacToe) getSimboloPosicaoMarcacoesEmSequencia() (simbolo int, posicao [2]int) {
	/*
		Função que retorna o símbolo(int) e a posição[linha,coluna]([2]int)
		que falta para ser marcado para ganhar o jogo da velha. Se retornar Vazio(0),
		a posição retornada é a última vazia livre.
	*/
	quantSequencia := LimiteLinha - 1
	quantLacos := LimiteLinha*2 + 2
	for i := 0; i < quantLacos; i++ {
		x, o, v := 0, 0, 0

		var sequencia [LimiteLinha]int
		p := i
		if i < LimiteLinha {
			sequencia = self.linha(p)
		} else if i < LimiteLinha*2 {
			p -= LimiteLinha
			sequencia = self.coluna(p)
		} else {
			p -= LimiteLinha * 2
			sequencia = self.diagonal(p)
		}

		// k é usado na posição da diagonal secundária
		k := len(sequencia) - 1
		for j, item := range sequencia {
			switch item {
			case X:
				x++
			case O:
				o++
			default:
				v++
				if i < LimiteLinha {
					posicao[0] = p
					posicao[1] = j
				} else if i < LimiteLinha*2 {
					posicao[1] = p
					posicao[0] = j
				} else {
					posicao[0] = j
					if p == 0 {
						posicao[1] = j
					} else {
						k -= j
						posicao[1] = k
					}
				}
			}
		}

		if x == quantSequencia && v == 1 && o == 0 {
			simbolo = X
			return
		} else if o == quantSequencia && v == 1 && x == 0 {
			simbolo = O
			return
		}
	}

	simbolo = Vazio
	return

}

// FUNÇÕES RELACIONADAS A STRUCT tictacToe
