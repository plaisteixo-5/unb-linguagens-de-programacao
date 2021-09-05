# unb-linguagens-de-programacao

Repositório dedicado à matéria de Linguagens de Programação

<div style="display: inline_block"><br>
  <img align="center" alt="Go" height="80" width="80" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg">
</div>

Este projeto será desenvolvido em Go. A ideia é explorar o conceito de concorrência numa aplicação desenvolvida em Go.

## Contextualização

Um serviço produtor de mensagens (Producer) é responsável por receber retornos relativos à três Universidades. A empresa responsável pela produção de mensagens 
envia os três tipos: A, B e C. Entretanto, algumas mensagens estão vindo com formatos incorretos e precisam ser filtradas. Essas universidades contraram uma equipe de
desenvolvedores para solucionar o problema de distribuição de mensagens, após muitas reclamações de mensagens que não chegavam ao seu destino. Os devs identificaram
que muitas mensagens incorretas estavam sobrecarregando as aplicações e o serviço de fila da AWS (SQS) - o qual continuava tentando reenviar as mensagens incorretas.
Além disso, foi percebido que os fluxos de mensagens eram muito maiores para alguns tipos de mensagens apenas. A implementação original era totalmente síncrona
e sequencial, gerando diversos problemas no processamento dessas mensagens. Nesse projeto, tentamos solucionar a questão utilizando os recursos da linguagem Go.


Vantagens dessa implementação -> Caso exista um fluxo muito grande de mensagens de A, isso não gerará impacto nos fluxos de B e C, tendo em vista o funcionamento independente.

### Arquitetura geral do projeto

![](/assets/arquitetura_geral_v002.png)

## Funcionamento do Consumer

![](/assets/consumer-estrutura-v2.png)

### Proposta de trabalho

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