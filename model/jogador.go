package model

type Jogador struct {
	Nome      string
	Simbolo   int
	Pontuacao []bool
}

// Variáveis não são acessíveis fora do seu pacote
// portanto, foi criado gets retornando ponteiros
var jogador0 Jogador
var jogador1 Jogador

func init() {
	jogador0.Nome, jogador0.Simbolo = "Jogador nº 1", X
	jogador1.Nome, jogador1.Simbolo = "Computador", O
}

func GetJogador(numero int) *Jogador {
	if numero == 1 {
		return &jogador1
	}
	return &jogador0
}

func (jogador *Jogador) SetPontuacao(vitoria bool) {
	jogador.Pontuacao = append(jogador.Pontuacao, vitoria)
}
