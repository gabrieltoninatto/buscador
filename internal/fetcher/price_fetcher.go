package fetcher

import (
	"buscador/internal/models"
	"math/rand/v2"
	"sync"
	"time"
)

// chan<- envia informações para o canal
func FetchPrices(priceChannel chan<- models.PriceDetail) {
	var wg sync.WaitGroup
	wg.Add(4)
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
		defer wg.Done()
		FetchSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

// Buscar preços de diferentes sites
func FetchPriceFromSite1() models.PriceDetail {
	time.Sleep(1 * time.Second)
	return models.PriceDetail{
		StoreName: "Jogos",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPriceFromSite2() models.PriceDetail {
	time.Sleep(3 * time.Second)
	return models.PriceDetail{
		StoreName: "Roupas",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPriceFromSite3() models.PriceDetail {
	time.Sleep(2 * time.Second)
	return models.PriceDetail{
		StoreName: "Bicicletas",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchSendMultiplePrices(priceChannel chan<- models.PriceDetail) {
	time.Sleep(6 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}
	for _, price := range prices {
		priceChannel <- models.PriceDetail{
			StoreName: "Moveis",
			Value:     price,
			Timestamp: time.Now(),
		}
	}

}
