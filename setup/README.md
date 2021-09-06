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