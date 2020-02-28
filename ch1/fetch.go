// Fetch exibe o conteúdo encontrado em cada URL especificada
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		} else {

			nbytes, err := io.Copy(ioutil.Discard, resp.Body)
			if err != nil {
				log.Fatal(os.Stderr, "fetch: %v\n", err)
			}
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(os.Stderr, "fetch: %v\n", err)
			}
			log.Printf("Código HTTP %s: %d6 bytes foram transferidos", resp.Status, nbytes)
			resp.Body.Close()
		}
	}
}
