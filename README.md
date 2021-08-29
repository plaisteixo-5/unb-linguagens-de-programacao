# unb-linguagens-de-programacao

Repositório dedicado à matéria de Linguagens de Programação

<div style="display: inline_block"><br>
  <img align="center" alt="Go" height="80" width="80" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg">
</div>

Este projeto será desenvolvido em Go. A ideia é explorar o conceito de concorrência numa aplicação desenvolvida em Go.

## Arquitetura geral do projeto

![](/assets/arquitetura_geral_v002.png)

## Proposta de trabalho

 Nosso consumer precisa resolver várias questões ao mesmo tempo e não pode realizar tudo de forma síncrona. Será necessário explorar os recursos da linguagem Golang
para lidar com várias situações de modo concorrente. Utilizaremos as Goroutines e channels como recursos da linguagem para lidar com os desafios que nossa aplicação tem.

## Setup do ambiente

1) Go

Documentação: https://golang.org/doc/install
<br>
Código inicial de exemplo (main.go) retirado da documentação: https://golang.org/doc/tutorial/web-service-gin

2) Sqs (localstack) + criação da fila (AWS CLI)

```
 docker-compose up 
```
 

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