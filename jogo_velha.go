package main

import (
	"fmt"
)

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

type posicoesRelativas [LimiteLinha][2]int

func main() {
	// velha := makeTicTacToe(1)

	for i := 0; i < 32; i++ {
		velha := makeTeste(i)
		velha.imprime()

		s, pos := velha.getSimboloPosicaoMarcacoesEmSequencia()
		fmt.Printf("Marcaçoes em sequência: [%s] linha: %d, coluna: %d\n\n", getSimbolo(s), pos[0]+1, pos[1]+1)

		s, pos = velha.getSimboloPosicaoDuasMarcacoesEmSequencia()
		fmt.Printf("2 Marcaçoes em sequência: [%s] linha: %d, coluna: %d\n\n", getSimbolo(s), pos[0]+1, pos[1]+1)
	}

	// os.Exit(0)
}

// FUNÇÕES RELACIONADAS A STRUCT ticTacToe

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
	case 24:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, X, Vazio},
			[3]int{X, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 25:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, O, Vazio},
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 26:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{X, Vazio, Vazio},
			[3]int{Vazio, X, Vazio},
		}
	case 27:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, O, Vazio},
		}
	case 28:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{O, X, Vazio},
			[3]int{X, Vazio, Vazio},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 29:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, O, X},
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, Vazio, Vazio},
		}
	case 30:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{X, Vazio, Vazio},
			[3]int{O, X, Vazio},
		}
	case 31:
		jogoVelha.tabuleiro = [3][3]int{
			[3]int{Vazio, Vazio, Vazio},
			[3]int{Vazio, Vazio, O},
			[3]int{Vazio, O, X},
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

func (self *ticTacToe) getLinha(posicao int) ([LimiteLinha]int, posicoesRelativas) {
	ehPosicaoValida(&posicao)

	var posicoes posicoesRelativas
	for i := 0; i < LimiteLinha; i++ {
		posicoes[i] = [2]int{posicao, i}
	}

	return self.tabuleiro[posicao], posicoes
}

func (self *ticTacToe) getColuna(posicao int) ([LimiteColuna]int, posicoesRelativas) {
	ehPosicaoValida(&posicao)
	var colunaTemp [LimiteColuna]int

	var posicoes posicoesRelativas
	for i := 0; i < LimiteLinha; i++ {
		posicoes[i] = [2]int{i, posicao}
	}

	for i, linha := range self.tabuleiro {
		for j, coluna := range linha {
			if j == posicao {
				colunaTemp[i] = coluna
			}
		}
	}

	return colunaTemp, posicoes
}

func (self *ticTacToe) getDiagonal(posicao int) ([LimiteLinha]int, posicoesRelativas) {
	var diagonalTemp [LimiteLinha]int

	var posicoes posicoesRelativas

	if posicao == 1 {
		// diagonal secundária
		posicao = LimiteColuna
		for i, linha := range self.tabuleiro {
			posicao--
			diagonalTemp[i] = linha[posicao]
		}

		col := LimiteLinha
		for i := 0; i <= 2; i++ {
			col--
			posicoes[i] = [2]int{i, col}
		}

	} else {
		// diagonal principal
		posicao = 0
		for i, linha := range self.tabuleiro {
			diagonalTemp[i] = linha[posicao]
			posicao++
		}

		for i := 0; i < LimiteLinha; i++ {
			posicoes[i] = [2]int{i, i}
		}
	}

	return diagonalTemp, posicoes
}

func ehPosicaoValida(posicao *int) {
	if *posicao >= LimiteLinha {
		*posicao = 0
	} else if *posicao < 0 {
		*posicao = 0
	}
}

func (self *ticTacToe) getSimboloPosicaoMarcacoesEmSequencia() (simbolo int, posicao [2]int) {
	/*
		Função que retorna o símbolo(int) e a posição[linha,coluna]([2]int)
		que falta para ser marcado para fazer 3 marcações em sequência.
		Se retornar Vazio(0), a posição retornada é a última vazia livre.
	*/
	quantSequencia := LimiteLinha - 1
	quantLacos := LimiteLinha*2 + 2
	for i := 0; i < quantLacos; i++ {
		x, o, v := 0, 0, 0 // quantidade de x, o, v(azio)

		var sequencia [LimiteLinha]int
		var posicoes [LimiteLinha][2]int

		p := i // p é a posição relativa para as funções de linha, coluna e diagonal
		if i < LimiteLinha {
			sequencia, posicoes = self.getLinha(p)
		} else if i < LimiteLinha*2 {
			p -= LimiteLinha
			sequencia, posicoes = self.getColuna(p)
		} else {
			p -= LimiteLinha * 2
			sequencia, posicoes = self.getDiagonal(p)
		}

		for j, item := range sequencia {
			switch item {
			case X:
				x++
			case O:
				o++
			default:
				v++

				posicao = posicoes[j]
			}

		}

		// se falta uma célula para marcar X
		if x == quantSequencia && v == 1 && o == 0 {
			simbolo = X
			return
			// se falta uma célula para marcar O
		} else if o == quantSequencia && v == 1 && x == 0 {
			simbolo = O
			return
		}
	}

	simbolo = Vazio
	return

}

func (self *ticTacToe) getSimboloPosicaoDuasMarcacoesEmSequencia() (simbolo int, posicao [2]int) {
	/*
		Função que retorna o símbolo(int) e a posição[linha,coluna]([2]int)
		que falta para formar duas marcações em sequência em dois eixos do tabuleiro.
		Se retornar Vazio(0), a posição retornada
	*/
	quantSequencia := LimiteLinha - 1
	quantLacos := 5

	for i := 0; i <= quantLacos; i++ {
		x, o, v := 0, 0, 0 // quantidade de x, o, v(vazio)

		var sequencia []int

		switch i {
		case 0:
			sequencia, posicao = self.getSequenciaPosicao(0, 0)
		case 1:
			sequencia, posicao = self.getSequenciaPosicao(0, 2)
		case 2:
			sequencia, posicao = self.getSequenciaPosicao(2, 2)
		case 3:
			sequencia, posicao = self.getSequenciaPosicao(2, 0)
		default:
			sequencia, posicao = self.getSequenciaPosicao(1, 1)
		}

		for j := 0; j < len(sequencia); j++ {
			item := sequencia[j]
			switch item {
			case X:
				x++
			case O:
				o++
			default:
				v++
			}
		}

		if x == quantSequencia && v == 3 && o == 0 && sequencia[0] == Vazio {
			simbolo = X
			return
		} else if o == quantSequencia && v == 3 && x == 0 && sequencia[0] == Vazio {
			simbolo = O
			return
		} else if x == quantSequencia && v == 3 && o == 0 && sequencia[1] == Vazio {
			simbolo = X
			return
		} else if o == quantSequencia && v == 3 && x == 0 && sequencia[1] == Vazio {
			simbolo = O
			return
		}

	}

	simbolo = Vazio
	return
}

func (self *ticTacToe) getSequenciaPosicao(linha int, coluna int) (sequencia []int, posicao [2]int) {
	var posicoes [][2]int

	s, p := self.getLinha(linha)
	for k := 0; k < len(s); k++ {
		sequencia = append(sequencia, s[k])
		posicoes = append(posicoes, p[k])
	}

	if linha == 1 && coluna == 1 {
		s, p = self.getColuna(coluna)
		for k := 0; k < len(s); k++ {
			if k != coluna {
				sequencia = append(sequencia, s[k])
				posicoes = append(posicoes, p[k])
			}
		}
	} else {
		s, p = self.getColuna(coluna)
		for k := 1; k < len(s); k++ {
			sequencia = append(sequencia, s[k])
			posicoes = append(posicoes, p[k])
		}
	}

	posicao = posicoes[coluna]
	return
}

// FUNÇÕES RELACIONADAS A STRUCT tictacToe
