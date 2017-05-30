// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma
// vez na entrada. Ele lê de stdin ou de euma lista de arquivos nomeados.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin, counts, names, "")
	} else {
		for _,arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, names, arg)
			f.Close()
		}
	}
	
	for line, n := range counts {
		if n > 1 {
			fmt.Println( n, line, names[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, names map[string]string, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		names[input.Text()] += arg
		names[input.Text()] += " "
	}
	// NOTA: ignorando erros em potencial de input.Err()
}