package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    url := "http://example.com/health" // Substitua pela URL da rota de health check da sua aplicação

    client := http.Client{
        Timeout: 5 * time.Second,
    }

    for {
        resp, err := client.Get(url)
        if err != nil {
            fmt.Printf("Erro ao acessar a URL: %v\n", err)
        } else {
            defer resp.Body.Close()
            if resp.StatusCode == http.StatusOK {
                fmt.Println("A aplicação está online!")
            } else {
                fmt.Printf("A aplicação retornou um status diferente de OK: %d\n", resp.StatusCode)
            }
        }

        // Aguarda 10 segundos antes de fazer o próximo health check
        time.Sleep(10 * time.Second)
    }
}