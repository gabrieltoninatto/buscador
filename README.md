# Buscador de Preços (Concorrência em Go)

Este projeto é uma aplicação em Go projetada para demonstrar o poder da concorrência utilizando **Goroutines** e **Channels**. O sistema simula a busca de preços em múltiplos sites simultaneamente e processa esses dados para calcular uma média.

## 🚀 Funcionalidades

- **Busca Paralela**: Utiliza goroutines para buscar preços em diferentes fontes (simuladas) ao mesmo tempo, reduzindo o tempo total de espera.
- **Comunicação via Channels**: Transfere dados de forma segura entre o módulo de busca (`fetcher`) e o de processamento (`processor`).
- **Sincronização**: Uso de `sync.WaitGroup` para gerenciar múltiplas tarefas de busca e canais para sinalizar a conclusão (`done`).

## 📂 Estrutura do Projeto

A organização do código segue padrões comuns de projetos Go:

- **`cmd/main.go`**: Ponto de entrada da aplicação.
  - Inicializa o canal de preços (`priceChannel`) com buffer.
  - Dispara as goroutines de busca e processamento.
  - Mede e exibe o tempo total de execução.
- **`internal/fetcher`**: Contém a lógica de busca.
  - `FetchPrices`: Orquestra as buscas simultâneas.
  - Simula latência de rede (ex: `time.Sleep`) e gera valores aleatórios.
- **`internal/processor`**: Responsável por consumir os preços do canal e calcular a média (função `ShowPriceAVG`).
- **`internal/models`**: Define as estruturas de dados, como `PriceDetail`.

## 🛠️ Como Executar

Certifique-se de ter o Go instalado em sua máquina.

1. Navegue até a raiz do projeto:
   ```bash
   cd c:\Users\gabri\Desktop\buscador
   ```

2. Execute o comando:
   ```bash
   go run cmd/main.go
   ```

## ⚙️ Detalhes de Implementação

O projeto utiliza um **canal com buffer** (`make(chan models.PriceDetail, 4)`) para evitar bloqueios imediatos caso o processador seja momentaneamente mais lento que os buscadores, e um canal `done` para garantir que o programa principal (`main`) aguarde o término de todo o processamento antes de encerrar.