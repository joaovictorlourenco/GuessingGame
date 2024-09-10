package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("*******************")
	fmt.Println("Um número será sorteado entre 0 e 100, tente acertar!")

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)

	guessNumbers := [10]int64{}

	for i := range guessNumbers {
 		fmt.Println("Digite um número:")
		scanner.Scan()

		guessNumber := scanner.Text()
		guessNumber = strings.TrimSpace(guessNumber)
		
		// Convertendo string para int64

		number, err := strconv.ParseInt(guessNumber, 10, 64)
		
		if err != nil {
			fmt.Println("Digite um número válido")
			i--
			return
		}

		switch {
		case number < x:
			fmt.Println("O número é maior")
		case number > x:
			fmt.Println("O número é menor")
		case number == x:
			fmt.Printf("Você venceu! O número era %d\n", x)
			fmt.Printf("Você acertou na %d tentativa \n", i+1)
			fmt.Printf("Essa foram suas tentativas: %v\n", guessNumbers[:i])
			return
		}

		guessNumbers[i] = number
	}

	fmt.Printf("Você perdeu! O número era %d\n", x)
	fmt.Printf("Você teve 10 tentaticas \n")
	fmt.Printf("Essa foram suas tentativas: %v\n", guessNumbers)
}