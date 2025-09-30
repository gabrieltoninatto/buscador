package fetcher

import (
	"math/rand/v2"
	"sync"
	"time"
)

// chan<- envia informações para o canal
func FetchPrices(priceChannel chan<- float64) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()
	}()

	go func() {
		FetchSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

// Buscar preços de diferentes sites
func FetchPriceFromSite1() float64 {
	time.Sleep(1 * time.Second)
	return rand.Float64() * 100
}

func FetchPriceFromSite2() float64 {
	time.Sleep(3 * time.Second)
	return rand.Float64() * 100
}

func FetchPriceFromSite3() float64 {
	time.Sleep(2 * time.Second)
	return rand.Float64() * 100
}

func FetchSendMultiplePrices(priceChannel chan<- float64) {
	time.Sleep(6 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}
	for _, price := range prices {
		priceChannel <- price
	}
}
