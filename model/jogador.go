package model

type Jogador struct {
	Nome      string
	Simbolo   int
	Pontuacao []bool
}

// Type Jogadores é um slice de ponteiros de Jogador
type Jogadores []*Jogador

var listaJogadores Jogadores

func AddJogador(jogador *Jogador) {
	listaJogadores = append(listaJogadores, jogador)
}

func GetJogador(numero int) *Jogador {
	jogador := listaJogadores[numero]
	// &*nomeVar retorna o endereço apontado pelo *ponteiro
	// jogador é um ponteiro de ponteiro, **ponteiro
	return &*jogador
}

// GetJogadorPorSimbolo retorna um *Jogador de acordo com simbolo informado
func GetJogadorPorSimbolo(simbolo int) *Jogador {
	for _, jogador := range listaJogadores {
		if jogador.Simbolo == simbolo {
			return &*jogador
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

func (jogador *Jogador) ResetaPontuacao() {
	jogador.Pontuacao = []bool{}
}
