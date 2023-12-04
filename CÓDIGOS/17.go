package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()

	mapa := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		mapa[str[i]]++
	}

	fmt.Println("Letras únicas na frase:")
	for i := 0; i < len(str); i++ {
		if mapa[str[i]] == 1 {
			fmt.Printf("%c ", str[i])
			mapa[str[i]] = 0 // marca como já impresso
		}
	}
	fmt.Println()

}
