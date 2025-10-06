package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/models"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// buffer
	priceChannel := make(chan models.PriceDetail, 4) // Canal para transmitir preços
	//len: tamanho atual do buffer e cap: tamanho máximo do buffer
	done := make(chan bool) // Canal para sinalizar quando o cálculo da média estiver completo

	//WaitGroup (grupo de espera)
	//Goroutines: funções ou métodos executados de forma concorrente.

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done // Espera até receber um sinal de que o cálculo da média está completo
	//done atua como um mecanismo de sincronização

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
