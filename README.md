# Consumo da API de Extrato do Banco do Brasil em Go

Este projeto em Go é um exemplo de como consumir a API de extrato do Banco do Brasil. O código está configurado para usar a API de testes, que é adequada apenas para desenvolvimento e não deve ser usada em ambientes de produção.

## Pré-requisitos

Antes de começar, certifique-se de que você tenha o Go instalado em sua máquina. Você pode baixá-lo e instalá-lo a partir do [site oficial do Go](https://golang.org/dl/).

## Configuração

1. **Clone o repositório**

   ```sh
   git clone https://github.com/lukaslimalkl/api-bb.git
   cd api-bb

2. **Crie um aquivo .env e use o arquivo .env.exemple como exemplo**

     ```sh
       cp .envexemple .env

3. **Instale as depedencias**

     ```sh
         go mod tidy

4. **Compile e use **

     ```sh
       go build -o /app
       cd ./app
