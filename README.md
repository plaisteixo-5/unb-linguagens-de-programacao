# Universidade de Brasília
# Linguagens de Programação

Repositório dedicado à matéria de Linguagens de Programação

<div style="display: inline_block"><br>
  <img align="center" alt="Go" height="80" width="80" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg">
</div>

Este projeto será desenvolvido em Go. A ideia é explorar o conceito de concorrência numa aplicação desenvolvida em Go.

## Contextualização

#### Geral

Três universidades se reuniram para fazer uma campanha de incentivo à leitura. Para isso, fizeram uma parceria com a Biblioteca Central da Cidade X. O que foi
acordado é que a Biblioteca enviará uma mensagem às Universidades A, B e C notificando os eventos de EMPRESTIMO, DEVOLUÇÃO e ATRASO. Cada uma dessas mensagens
também deverão ser enviadas a um Comitê Central (composto por membros das três universidades), o qual ficará responsável pela conferência da contagem
feita por cada uma das Universidades participantes e servirá como recurso reserva caso o sistema de uma das universidades fique offline.


#### Situação problema

As universidades possuem muitos alunos e o fluxo de mensagens pode ser intenso. Para isso, utilizaremos um serviço de fila da Amazon Web Services (SQS). A fila SQS 
é escalável, performática e segura. Assim garantimos que a competição tenha uma contagem fidedigna.

### Construção da solução

##### De início, desenvolveremos a aplicação de uma Biblioteca.

Essa aplicação será responsável apenas por receber três tipos de mensagens e enviar para o Producer.

##### Em seguida, serão desenvolvidas duas aplicações para interagir com a fila SQS: o Producer e o Consumer.

Producer: será responsável por receber as mensagens da Biblioteca e enviar pra fila.
Consumer: será responsável por receber as mensagens da fila e direcionar para as rotas corretas.

##### Por fim, as APIs das Universidades e do Comitê Central. 

api-universidade-a: recebe a mensagem e faz a persistência no banco de dados.
api-universidade-b: recebe a mensagem e faz a persistência no banco de dados.
api-universidade-c: recebe a mensagem e faz a persistência no banco de dados.

#### Pontuação

Empréstimo: 5 pontos
Devolução: 10 pontos
Atraso para devolução: -5 pontos

#### Premiação

A Universidade que conseguir incentivar o maior número de estudantes a ler, ganhará um ano de recursos ilimitados nos serviços da AWS (Amazon Web Services), com
direito a cursos direcionados à computação em nuvem.

### Arquitetura geral do projeto

![](/assets/arquitetura_geral_v004.png)

## Funcionamento do Consumer

![](/assets/consumer-estrutura-v2.png)


## Setup do ambiente

1) Go

Documentação: https://golang.org/doc/install
<br>
Código inicial de exemplo (main.go) retirado da documentação: https://golang.org/doc/tutorial/web-service-gin

2) Sqs (localstack) + criação da fila (AWS CLI)

```
 docker-compose up 
```


### Exemplo de uso das aplicações

1) executar o seguinte comando na pasta raíz do projeto:

> docker compose up 

2) definir as variáveis de ambiente na pasta do consumer e do producer:

```
 export AWS_ACCESS_KEY_ID=teste &&
 export AWS_SECRET_ACCESS_KEY=teste &&
 export AWS_SESSION_TOKEN=teste &&
 export AWS_DEFAULT_REGION=sa-east-1 &&
 export AWS_DEFAULT_OUTPUT=json
```

3) definir a variável de ambiente do GO (https://golang.org/doc/install):

> export PATH=$PATH:/usr/local/go/bin

4) Nesse momento a fila já está pronta, precisamos executar nossas aplicações:
4.1) Na pasta /consumer:
   
> go get .

<br/>

> go run .
 
obs.: talvez seja necessário executar o comando go mod tidy, caso alteremos as dependências do projeto

5) Nesse momento a fila já está pronta, precisamos executar nossas aplicações:
5.1) Na pasta /producer:

> go get .

<br/>

> go run .

obs.: talvez seja necessário executar o comando go mod tidy, caso alteremos as dependências do projeto

6) Em seguida, produza mensagens na pasta /producer (deixe a aplicação gerando mensagens lá por algum tempo):

> watch -n 5 go run .

7) Depois podemos acompanhar o funcionamento do consumer (aplicação foco do trabalho)


### SQS

#### Introdução

Obs.: o ponto de partida para configurar o producer e o consumer foi a documentação da AWS:

Consumer:

```
https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/sqs/ReceiveMessage/ReceiveMessage.go
```

Producer:

```
https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/sqs/SendMessage/SendMessage.go
```

#### Comandos úteis

##### Criar uma fila

```
 aws --endpoint-url=http://localstack:4556 sqs create-queue --queue-name=fila_trabalho_lp
```


##### Listar as filas

```
aws --endpoint-url=http://localhost:4566 sqs list-queues
```

##### Listar mensagens

```
 aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url http://localhost:4566/000000000000/fila_trabalho_lp
```

##### Enviar mensagens

```
 aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url http://localhost:4566/000000000000/fila_trabalho_lp --message-body "Teste"
```

##### Variáveis de ambiente (linux)

```
 export AWS_ACCESS_KEY_ID=teste &&
 export AWS_SECRET_ACCESS_KEY=teste &&
 export AWS_SESSION_TOKEN=teste &&
 export AWS_DEFAULT_REGION=sa-east-1 &&
 export AWS_DEFAULT_OUTPUT=json
```

Referências:

- https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sqs-example-receive-message.html
- https://medium.com/trainingcenter/goroutines-e-go-channels-f019784d6855
- https://medium.com/trendyol-tech/concurrency-and-channels-in-go-bbc4dea75286
- https://marcioghiraldelli.medium.com/elegant-use-of-golang-channels-with-aws-sqs-dad20cd59f34
- https://medium.com/@tilaklodha/concurrency-and-parallelism-in-golang-5333e9a4ba64