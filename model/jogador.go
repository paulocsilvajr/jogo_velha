package model

// Jogador é uma struct que representa as características de um jogador.
// Características: Nome, Simbolo, Humano, Pontuação.
type Jogador struct {
	Nome      string
	Simbolo   Elemento
	Humano    bool
	Pontuacao []bool
}

// Jogadores é um slice de ponteiros de Jogador
type Jogadores []*Jogador

var listaJogadores Jogadores

// AddJogador adiciona o jogador na lista local de jogadores.
func AddJogador(jogador *Jogador) {
	listaJogadores = append(listaJogadores, jogador)
}

// GetJogador retorna um ponteiro de jogador de acordo com o número informado.
// Numero é a posição na lista de jogadores, geralmente de 0 a 1.
func GetJogador(numero int) *Jogador {
	jogador := listaJogadores[numero]
	// &*nomeVar retorna o endereço apontado pelo *ponteiro
	// jogador é um ponteiro de ponteiro, **ponteiro
	return &*jogador
}

// GetJogadorPorSimbolo retorna um ponteiro de Jogador de acordo com simbolo informado.
func GetJogadorPorSimbolo(simbolo int) *Jogador {
	for _, jogador := range listaJogadores {
		if jogador.Simbolo == Elemento(simbolo) {
			return &*jogador
		}
	}
	return nil
}

// GetJogadores retorna um slice de ponteiro de todos os jogadores cadastrados.
func GetJogadores() Jogadores {
	return listaJogadores
}

// SetPontuacao define a pontuação(true = vitória, false = derrota) do jogador.
func (jogador *Jogador) SetPontuacao(vitoria bool) {
	jogador.Pontuacao = append(jogador.Pontuacao, vitoria)
}

// ResetaPontuacao reseta a pontuação do jogador.
func (jogador *Jogador) ResetaPontuacao() {
	jogador.Pontuacao = []bool{}
}

// EhComputador retorna true se o jogador é o computador.
func (jogador *Jogador) EhComputador() bool {
	return !jogador.Humano
}

// GetSimboloAdversario retorna o simbolo do adversário do jogador.
func (jogador *Jogador) GetSimboloAdversario() (simbolo Elemento) {
	switch jogador.Simbolo {
	case X:
		simbolo = Elemento(O)
	case O:
		simbolo = Elemento(X)
	}
	return
}
