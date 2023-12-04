package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma string: ")
	scanner.Scan()
	str := scanner.Text()

	str = strings.ReplaceAll(str, "a", "")
	str = strings.ReplaceAll(str, "e", "")
	str = strings.ReplaceAll(str, "i", "")
	str = strings.ReplaceAll(str, "o", "")
	str = strings.ReplaceAll(str, "u", "")

}
