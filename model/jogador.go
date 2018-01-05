package model

type Jogador struct {
	Nome      string
	Simbolo   int
	Pontuacao []bool
}

type Jogadores []Jogador

var listaJogadores Jogadores

func AddJogador(jogador Jogador) {
	listaJogadores = append(listaJogadores, jogador)
}

func GetJogador(numero int) *Jogador {
	return &listaJogadores[numero]
}

// GetJogadorPorSimbolo retorna um *Jogador de acordo com simbolo informado
func GetJogadorPorSimbolo(simbolo int) *Jogador {
	for _, jogador := range listaJogadores {
		if jogador.Simbolo == simbolo {
			return &jogador
		}
	}
	return nil
}

func GetJogadores() Jogadores {
	return listaJogadores
}

func (jogador *Jogador) SetPontuacao(vitoria bool) {
	jogador.Pontuacao = append(jogador.Pontuacao, vitoria)
}
