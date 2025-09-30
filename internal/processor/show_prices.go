package processor

import (
	"buscador/internal/models"
	"fmt"
)

// chan<- = enviar <-chan = ler
func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("Preço recebido de %s | R$ %.2f | Preço média até agora: R$ %.2f \n", price.StoreName, price.Value, avgPrice)
	}
	done <- true // Sinaliza que o cálculo da média está completo
}
