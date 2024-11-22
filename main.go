package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	url         string
	totalReqs   int
	concurrency int
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "loadtest",
		Short: "CLI de testes de stress para serviços web",
		Run:   runLoadTest,
	}

	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado (required)")
	rootCmd.Flags().IntVar(&totalReqs, "requests", 0, "Número total de requests (required)")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 0, "Número de chamadas simultâneas (required)")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func runLoadTest(cmd *cobra.Command, args []string) {
	fmt.Printf("Iniciando teste na URL: %s\n", url)
	fmt.Printf("Total de requests: %d\n", totalReqs)
	fmt.Printf("Chamadas Simultâneas: %d\n\n", concurrency)

	startTime := time.Now()
	results := make(chan int, totalReqs)

	var wg sync.WaitGroup

	worker := func(requests <-chan int, results chan<- int) {
		for range requests {
			resp, err := http.Get(url)
			if err != nil {
				results <- -1
				continue
			}
			results <- resp.StatusCode
			resp.Body.Close()
		}
		wg.Done()
	}

	requests := make(chan int, totalReqs)
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(requests, results)
	}

	go func() {
		for i := 0; i < totalReqs; i++ {
			requests <- i
		}
		close(requests)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	statusCounts := make(map[int]int)
	for status := range results {
		statusCounts[status]++
	}

	duration := time.Since(startTime)

	fmt.Printf("\nTempo total gasto na execução: %v\n", duration)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalReqs)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", statusCounts[200])
	fmt.Println("Outros status:")
	for status, count := range statusCounts {
		if status != 200 && status != -1 {
			fmt.Printf("  HTTP %d: %d\n", status, count)
		}
	}
	if statusCounts[-1] > 0 {
		fmt.Printf("  Errors: %d\n", statusCounts[-1])
	} else {
		fmt.Println("  Não houveram outros status")
	}
}
