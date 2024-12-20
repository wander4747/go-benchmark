package main

import (
	"flag"
	"github.com/wander4747/go-benchmark/analysis"
	"log"
	"os"
	"time"
)

func main() {
	// Definindo os parâmetros via flags
	url := flag.String("url", "", "URL da API para o analysis")
	method := flag.String("method", "GET", "Método HTTP para as requisições (GET ou POST)")
	requests := flag.Int("requests", 100, "Número total de requisições a serem realizadas")
	duration := flag.String("duration", "30s", "Duração do analysis (ex: 30s, 1m, 2h)")
	payload := flag.String("payload", `{"key":"value"}`, "Payload para requisição POST")

	// Parse dos parâmetros
	flag.Parse()

	// Validando se a URL foi informada
	if *url == "" {
		log.Fatal("A URL deve ser informada!")
		os.Exit(1)
	}

	// Convertendo a duração para time.Duration
	durationParsed, err := time.ParseDuration(*duration)
	if err != nil {
		log.Fatalf("Erro ao parsear duração: %v", err)
	}

	// Executando o analysis
	err = analysis.RunBenchmark(*url, *method, *requests, durationParsed, *payload)
	if err != nil {
		log.Fatalf("Erro ao executar o analysis: %v", err)
	}
}
