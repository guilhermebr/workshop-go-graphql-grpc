package main

import "fmt"

var motivos = []string{
	"Compilada - desempenho aproximando C para tarefas de uso intensivo da CPU",
	"Concorrente - Projetado para hardware moderno (multicore, redes)",
	"Estaticamente Tipada (mas com sentimento de dinamico)",
	"Simples",
	"Produtiva",
	"Divertida",
	"UTF-8 ♥",
	"Sem ponto e virgula ツ",
}

func main() {
	for i, motivo := range motivos {
		fmt.Printf("%d - %s,\n", i, motivo)
	}
}
