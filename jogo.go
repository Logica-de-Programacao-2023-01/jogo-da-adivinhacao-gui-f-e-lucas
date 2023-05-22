package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// criação de uma vairavel booleana para gerar condição no laço for
	fmt.Println("Bem vindo ao jogo da adivinhação!")
	pedirJoogar := true
	var todasTentativas [][]int

	for pedirJoogar {
		//primeiro loop para criação de um numero aleatorio e repetição das condições básicas do game
		respoosta := numAleatorio()
		tentativas := make([]int, 0)
		encontrou := false

		for !encontrou {
			//loop para encontrar a resposta, caso a variavel "encontrou" seja verdadeira, o loop termina
			chute := numUsuario()
			tentativas = append(tentativas, chute)

			if chute == respoosta {
				fmt.Println("Parabéns, você acertou!")
				encontrou = true
				fmt.Printf("Você usou %d tentativas\n", len(tentativas))
			} else if chute < respoosta {
				fmt.Printf("O número é maior que %d\n", chute)
			} else if chute <= 0 || chute > 100 {
				fmt.Println("Digite um número valido")
			} else {
				fmt.Printf("O número é menor que %d\n", chute)
			}
		}
		todasTentativas = append(todasTentativas, tentativas)

		pedirJoogar = pedirParaJoogar()
	}

	fmt.Println("\nHistórico de tentativas:")

	for x, tentativas := range todasTentativas {
		fmt.Printf("Jogada : %d  -  Tentativas: %d\n ", x+1, tentativas)
	}
}

func numAleatorio() int { //gera um número aleatório
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

func numUsuario() int { //garda oo número do usuario
	var chute int
	fmt.Print("Digite um número entre 1 e 100: ")
	fmt.Scanln(&chute)
	return chute
}

func pedirParaJoogar() bool { //pede para joogar novamente
	var pedir string
	fmt.Println("Você deseja jogar novamente? (s/n):")
	fmt.Scanln(&pedir)

	//retorna true caso escreva "s", caso o cntrario, retorna falso
	return pedir == "s" || pedir == "S"
}
