package processor

import "fmt"

// chan<- = enviar <-chan = ler
func ShowPriceAVG(priceChannel <-chan float64, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("Preço recebido: R$ %.2f | Preço médio até agora: R$ %.2f\n", price, avgPrice) // Adicionado avgPrice
	}
	done <- true // Sinaliza que o cálculo da média está completo
}
