#!/bin/bash

# Função para perguntar qual script executar
choose_script() {
  echo "Qual script você deseja executar?"
  echo "1) Instalar Python e Node Exporter"
  echo "2) Instalar Grafana, Prometheus, Golang e Node Exporter"
  read -p "Escolha uma opção (1 ou 2): " choice

  case $choice in
    1)
      script_path="scriptPython/docker-compose.yml"
      ;;
    2)
      script_path="scriptGoNodePromethGrafa/docker-compose.yml"
      ;;
    *)
      echo "Opção inválida. Saindo."
      exit 1
      ;;
  esac
}

# Atualiza os pacotes e instala dependências
sudo apt-get update
sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# Adiciona a chave GPG oficial do Docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Adiciona o repositório do Docker
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Instala o Docker Engine
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# Instala o Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/$(curl -s https://api.github.com/repos/docker/compose/releases/latest | grep -oP '"tag_name": "\K(.*)(?=")')/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Verifica as instalações
docker --version
docker-compose --version

# Pergunta qual script executar
choose_script

# Executa o Docker Compose com o script escolhido
sudo docker-compose -f $script_path up -d