package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/processor"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)

	//WaitGroup (grupo de espera)
	var showWg sync.WaitGroup
	showWg.Add(1)

	//Goroutines: funções ou métodos executados de forma concorrente.

	go func() {
		defer showWg.Done()
		processor.ShowPriceAVG(priceChannel)
	}()

	go fetcher.FetchPrices(priceChannel)

	showWg.Wait()

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
