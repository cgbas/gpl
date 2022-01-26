// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma
// vez na entrada. Ele lê de stdin ou de uma lista de arquivos nomeados

// Ex 1.4: modifique dup3 para que exiba os nomes de todos os arquivos em que cada linha duplicada ocorre
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "os.Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, filenames := range counts {
		if len(filenames) <= 1 {
			continue
		}

		fmt.Printf("Encontrado em %d arquivo(s) o item [%s]\n", len(filenames), line)
		total := 0
		for name, count := range filenames {
			fmt.Printf("\t%d ocorrencia(s) em %s\n", count, name)
			total += count
		}
		fmt.Printf("Total de: %d ocorrencia(s)\n\n", total)
	}
}
func countLines(f *os.File, counts map[string]map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
	// NOTA: ignorando erros em potencial de input.Err()
}
