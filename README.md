# unb-linguagens-de-programacao

Repositório dedicado à matéria de Linguagens de Programação

<div style="display: inline_block"><br>
  <img align="center" alt="Go" height="80" width="80" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg">
</div>

Este projeto será desenvolvido em Go. A ideia é explorar o conceito de concorrência numa aplicação desenvolvida em Go.

## Setup do ambiente

1) Go

Documentação: https://golang.org/doc/install
Código inicial de exemplo (main.go) retirado da documentação: https://golang.org/doc/tutorial/web-service-gin

2) Sqs (localstack)

> docker-compose up 

3) Aws CLI

https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-linux.html


### SQS

- aws --endpoint-url=http://localstack:4576 sqs list-queues
- aws --endpoint-url=http://localstack:4576 sqs create-queue --queue-name=filaTeste