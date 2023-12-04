package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite a primeira string: ")
	scanner.Scan()
	str1 := scanner.Text()

	fmt.Print("Digite a segunda string: ")
	scanner.Scan()
	str2 := scanner.Text()

	if isAnagram(str1, str2) {
		fmt.Printf("%s e %s são anagramas.\n", str1, str2)
	} else {
		fmt.Printf("%s e %s não são anagramas.\n", str1, str2)
	}

}
func isAnagram(str1, str2 string) bool {
	// remove espaços em branco e converter para letras minusculas
	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	// converter strings para slice em bytes
	str1Byte := []byte(str1)
	str2Byte := []byte(str2)

	// ordena slices em byte
	sort.Slice(str1Byte, func(i, j int) bool {
		return str1Byte[i] < str1Byte[j]
	})
	sort.Slice(str2Byte, func(i, j int) bool {
		return str2Byte[i] < str2Byte[j]
	})

	// comparando slices
	return string(str1Byte) == string(str2Byte)
}
