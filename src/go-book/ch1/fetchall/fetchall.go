// Fetchall busca URLs em paralelo e informa os tempos gastos e os tamanhos totais

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // inicia uma goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	f, err := os.OpenFile("fetchall.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		ch <- fmt.Sprint(err) // envia para o canal ch
		return
	}
	fmt.Fprintf(f, "%.2fs elapsed since the start.\n", time.Since(start).Seconds())
	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // envia para o canal ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // evitando vazamento de recursos
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
