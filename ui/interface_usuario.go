package ui

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var opcoes = []int{0, 1, 2}

func MostraMenuInicial() (opcao int) {
	for {
		LimpaTelaLinux()

		fmt.Println("Bem Vindo!")
		fmt.Println(opcoes[1], "- Jogador x Jogador")
		fmt.Println(opcoes[2], "- Jogador x Computador")
		fmt.Println(opcoes[0], "- Sair")
		fmt.Print("\t\t: ")

		_, err := fmt.Scan(&opcao)
		if err != nil {
			fmt.Println("Escolha uma das opções do menu")
			Espere()

			continue
		}

		for i := range opcoes {
			if opcao == opcoes[i] {
				return opcao
			}
		}

		fmt.Printf("Opção informada [%d] inválida\n", opcao)
		Espere()
	}
}

func DefineModoJogo() {
	for {
		op := MostraMenuInicial()

		switch op {
		case 1:
			fmt.Println("Jogador x Jogador")
			tab.Imprime()
			Espere()
		case 2:
			fmt.Println("Jogador x Computador")
			Espere()
		default:
			FimDeJogo()
			os.Exit(0)
		}
	}
}

func LimpaTelaLinux() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Espere() {
	time.Sleep(time.Second)
}

func FimDeJogo() {
	fmt.Println("END OF LINE.")
}
