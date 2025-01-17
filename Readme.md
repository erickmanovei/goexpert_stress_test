# Desafio Stress Test

Este repositório se trata do desafio **Stress Test**, da Pós Graduação **Go Expert / FullCycle**.

## Escopo do Desafio

Objetivo: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.


Execução do Teste:

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.

Geração de Relatório:

Apresentar um relatório ao final dos testes contendo:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

1- Execução da aplicação:
Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
docker run <sua imagem docker> —url=http://google.com —requests=1000 —concurrency=10

## Passos para executar o projeto

1. Crie o build executando este comando na raiz do projeto:
`docker build --no-cache -t loadtest .`

2. Execute este comando para realizar o teste de stress na URL desejada (substitua os valores das flags para os desejados):
`docker run loadtest --url=https://google.com --requests=100 --concurrency=10`
