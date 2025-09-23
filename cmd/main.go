package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	price1 := fetcher.FetchPriceFromSite1()
	price2 := fetcher.FetchPriceFromSite2()
	price3 := fetcher.FetchPriceFromSite3()

	fmt.Printf("R$%.2f \n", price1)
	fmt.Printf("R$%.2f \n", price2)
	fmt.Printf("R$%.2f \n", price3)

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
