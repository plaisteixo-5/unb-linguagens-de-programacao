# Universidade de Brasília 2021
## *Linguagens de Programação*

Repositório dedicado à matéria de Linguagens de Programação

<div style="display: inline-block"><br>
  <img alt="Go" height="150" width="150" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg">
</div>

Este projeto será desenvolvido em Go. A ideia é explorar o conceito de concorrência numa aplicação desenvolvida em Go.

## *Grupo de trabalho*

Felipe Fontelene <br>
William Coelho <br>
Marcelo Amorim <br>
João Victor Alves <br>
Gabriel Angelo Alves

## *Referência Bibliográfica*:

- https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sqs-example-receive-message.html
- https://medium.com/trainingcenter/goroutines-e-go-channels-f019784d6855
- https://medium.com/trendyol-tech/concurrency-and-channels-in-go-bbc4dea75286
- https://marcioghiraldelli.medium.com/elegant-use-of-golang-channels-with-aws-sqs-dad20cd59f34
- https://medium.com/@tilaklodha/concurrency-and-parallelism-in-golang-5333e9a4ba64

## Setup do ambiente para executar as aplicações

[*Fazer setup*](./setup/README.md)

## Contextualização

### a) Geral

Três universidades se reuniram para fazer uma campanha de incentivo à leitura. Para isso, fizeram uma parceria com a Biblioteca Central da Cidade X. O que foi
acordado é que a Biblioteca enviará uma mensagem às Universidades A, B e C notificando os eventos de EMPRESTIMO, DEVOLUÇÃO e ATRASO. Cada uma dessas mensagens
também deverão ser enviadas a um Comitê Central (composto por membros das três universidades), o qual ficará responsável pela conferência da contagem
feita por cada uma das Universidades participantes e servirá como recurso reserva caso o sistema de uma das universidades fique offline.


### b) Situação problema

As universidades possuem muitos alunos e o fluxo de mensagens pode ser intenso. Para isso, utilizaremos um serviço de fila da Amazon Web Services (SQS). A fila SQS 
é escalável, performática e segura. Assim garantimos que a competição tenha uma contagem fidedigna.

## Construção da solução

### a) De início, desenvolveremos a aplicação de uma Biblioteca.

Essa aplicação será responsável apenas por receber três tipos de mensagens e enviar para o Producer.

### b) Em seguida, serão desenvolvidas duas aplicações para interagir com a fila SQS: o Producer e o Consumer.

Producer: será responsável por receber as mensagens da Biblioteca e enviar pra fila.
Consumer: será responsável por receber as mensagens da fila e direcionar para as rotas corretas.

### c) Por fim, as APIs das Universidades e do Comitê Central. 

api-universidade-a: recebe a mensagem e faz a persistência no banco de dados.
api-universidade-b: recebe a mensagem e faz a persistência no banco de dados.
api-universidade-c: recebe a mensagem e faz a persistência no banco de dados.

## Pontuação

Empréstimo: 5 pontos
Devolução: 10 pontos
Atraso para devolução: -5 pontos

## Premiação

A Universidade que conseguir incentivar o maior número de estudantes a ler, ganhará um ano de recursos ilimitados nos serviços da AWS (Amazon Web Services), com
direito a cursos direcionados à computação em nuvem.

## Arquitetura geral do projeto

![](/assets/arquitetura_geral_v004.png)

## Funcionamento do Consumer

![](/assets/consumer-estrutura-v3.png)

## Funcionamento do Producer

![](/assets/producer-estrutura-v1.png)

## Funcionamento das APIs das Universidades



