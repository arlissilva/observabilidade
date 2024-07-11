package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var customMetric = prometheus.NewGauge(prometheus.GaugeOpts{
    Name: "monitora_app_go",
    Help: "Métrica customizada para monitoramento da aplicação em Golang",
})

func main() {
    prometheus.MustRegister(customMetric)

    // Substitua pela URL da aplicação que você deseja testar
    url := "https://sua-aplicacao.com"

    // Execute a verificação em loop com intervalos de 1 minuto
    go func() {
        for {
            checkSiteStatus(url)
            time.Sleep(time.Minute)
        }
    }()

    // Crie um handler para expor as métricas
    http.Handle("/metrics", promhttp.Handler())

    // Inicie o servidor HTTP
    fmt.Println("Servidor iniciado na porta 8080")
    http.ListenAndServe(":8080", nil)
}

func checkSiteStatus(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Erro ao fazer a requisição:", err)
        customMetric.Set(0) // Site offline
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        fmt.Println("Aplicação está online!")
        customMetric.Set(1) // Site online
    } else {
        fmt.Println("Aplicação não está online. Código de status:", resp.StatusCode)
        customMetric.Set(0) // Site offline
    }
}
