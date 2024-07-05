# Ambiente de testes com Grafana e Prometheus

### <br> Resumo do ambiente de testes 
 

O ambiente de testes é composto de acordo com as especificações abaixo: 

  * INSTANCIA AWS EC2 - 1

    *  Softwares instalados: 
       * Docker Container (Grafana) - Imagem grafana/grafana:latest
         * Grafana 
          
       *  Docker Container (Prometheus) - Imagem prom/prometheus:latest
          * Prometheus

       * Docker Container (Node Exporter/Go) - Imagem redhat/ubi9:latest
         * App Golang 
         * Node Exporter 
<br><br>
*  INSTANCIA AWS EC2 - 2 
    * Softwares instalados:
      * Docker Container (Python/Node Exporter) - Imagem redhat/ubi9:latest
        * App Python
        * Node Exporter
 

### <br> Descrição dos APPs utilizados 

Dois apps simples foram desenvolvidos para os testes, o primeiro utilizando linguagem de programação Golang dispara um ping continuo para o IP do servidor da segunda aplicação e valida se o servidor está online. 

O segundo app foi desenvolvido em python e fica criando, editando e excluindo um par de arquivos para gerar uso de memória. 

 ### <br>Implantação do ambiente de testes 
O ambiente de testes foi executado em duas instancias EC2 da AWS, sendo que a primeira instancia foi instalado o Prometheus, Grafana, app Golang e Node exporter e na seguda instancia o Node Exporter e um app python.

<br> Para agilizar o processo de criação do ambiente os scripts foram criados e podem ser executados dentro das instancias EC2 ou em ambiente local linux ou WSL.

<br> Dentro do terminal bash do linux entre na pasta GrafanaWithPrometheus e execute o comando abaixo:

``` bash
sudo chmod +x install_docker-docker-compose.sh && ./install_docker-docker-compose.sh
```

Esse comando permissiona o script como executavel em seguida instala o docker e o docker-compose e pergunta qual script será executado (<span style="color: green;"> 1 - Instalar Python e Node Exporter</span> ou <span style="color: green;">2 - Instalar Grafana, Prometheus, Golang e Node Exporter</span>  ) para criar os containers.

No ambiente da AWS foi executado um script para cada instancia, mas pode ser executado os dois scripts no mesmo local.

<br><br>
--------------------
                    TRECHO DESTINADO SOMENTE AO AMBIENTE AWS

 ### <br> Grupos de Segurança Instancias AWS
Para uma comunicação segura entre as instancias AWS e seus aplicativos é necessario criar os grupos de segurança para cada instancia com as suas respectivas politicas de portas TCP.

Dois grupos foram criados e adicionados para as suas intancias.

O primeiro com as regras de entrada (InboudRules)

<span style="color: green;"> Promethus - Protocolo: IPv4  - Port: 9090 - Source: 0.0.0.0/0 </span>
<br><span style="color: green;"> Grafana - Protocolo: IPv4  - Port: 3000 - Source: 0.0.0.0/0 </span>
<br><span style="color: green;"> Node Exporter - Protocolo: IPv4  - Port: 9100 - Source: 0.0.0.0/0 </span>
<br><span style="color: green;"> App Golang - Protocolo: IPv4  - Port: 8080 - Source: 0.0.0.0/0 </span>

O segundo grupo com as regras de entrada (InboudRules)

<span style="color: green;"> Python - Protocolo: IPv4  - Port: 8000 - Source: 0.0.0.0/0 </span>
<br><span style="color: green;"> Node Exporter - Protocolo: IPv4  - Port: 9300 - Source: 0.0.0.0/0 </span>


--------------------
                    TRECHO DESTINADO SOMENTE AO AMBIENTE AWS

Após execução do script install_docker-docker-compose.sh as aplicações ficam disponiveis nos endereços abaixo:
 > Grafana = Ip_do_host:3000
 <br>
 > Promethus = Ip_do_host:9090
 <br>
 > Node Exporter = Ip_do_host:9100
 <br>
 > App Golang = Ip_do_host:8080
 <br><br>
 > Node Exporter = Ip_do_host:9300
 <br>
 App Python = Ip_do_host:8080
 
  