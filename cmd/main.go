package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	done := make(chan bool) // Canal para sinalizar quando o cálculo da média estiver completo

	//WaitGroup (grupo de espera)
	//Goroutines: funções ou métodos executados de forma concorrente.

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done // Espera até receber um sinal de que o cálculo da média está completo
	//done atua como um mecanismo de sincronização

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
