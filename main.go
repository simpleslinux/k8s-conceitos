package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var conteudoPadrao = "<h1>Teste APP!</h1>"
var portaPadrao = "6377"

func handler(w http.ResponseWriter, r *http.Request) {
	content := os.Getenv("CONTEUDO")
	if content == "" {
		content = conteudoPadrao
	}
	fmt.Fprint(w, content)
}

func main() {
	porta := os.Getenv("PORTA")
	if porta == "" {
		porta = portaPadrao
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":" + porta, nil))
}
