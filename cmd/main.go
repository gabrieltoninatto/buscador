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
	var wg, showWg sync.WaitGroup
	wg.Add(3)
	showWg.Add(1)

	//Goroutines: funções ou métodos executados de forma concorrente.

	go func() {
		defer showWg.Done()
		processor.ShowPriceAVG(priceChannel)
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite3()
	}()

	wg.Wait()
	close(priceChannel)
	showWg.Wait()

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
