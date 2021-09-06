package functional

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

func GetMessages(sess *session.Session, queueURL string, timeout int64) (*sqs.ReceiveMessageOutput, error) {

    svc := sqs.New(sess)

    msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
        AttributeNames: []*string{
            aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
        },
        MessageAttributeNames: []*string{
            aws.String(sqs.QueueAttributeNameAll),
        },
        QueueUrl:            aws.String(queueURL),
        MaxNumberOfMessages: aws.Int64(10),
        VisibilityTimeout:   aws.Int64(timeout),
    })

    if err != nil {
        fmt.Println("[METHOD - GetMessages] Não foi possível buscar novas mensagens...algum erro ocorreu!")
        return nil, err
    }

    return msgResult, nil
}


func DeleteMessage(sess *session.Session, queueURL string, messageHandle string) error {

    svc := sqs.New(sess)

    _, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
        QueueUrl:      aws.String(queueURL),
        ReceiptHandle: aws.String(messageHandle),
    })

    if err != nil {
        fmt.Println("[METHOD - DeleteMessage] Não foi possível deletar essa mensagem de messageHandle = ", messageHandle)
        return err
    }else {
        fmt.Println("[METHOD - DeleteMessage] Mensagem deletada com sucesso! MessageHandle = ", messageHandle)
    }

    return nil
}

func sendMessagesToChannel(chn chan<- *sqs.Message, msgResult *sqs.ReceiveMessageOutput) {

      for _, message := range msgResult.Messages {

        fmt.Println("[POLL SQS] Enviando mensagens para o channel...")
        chn <- message

      }

}

func handleMessage(message sqs.Message, sess *session.Session, queueURL string, messageCounter int) {

    fmt.Println("[METHOD - handleMessage] MESSAGE NUMBER= ", messageCounter, " BEING SEND TO RESPECTIVE DESTINATION")
    SendByPOSTRequest(message)
    DeleteMessage(sess, queueURL, *message.ReceiptHandle)

}

func HandleAllReceivedMessages(chnMessages chan *sqs.Message, sess *session.Session) {

    messageCounter := 0
    for message := range chnMessages {

       messageCounter += 1

       fmt.Println("[METHOD - handleAllReceivedMessages] As mensagens estão sendo enviadas para seu destino final...")

       // Função executada de forma concorrente...
       go handleMessage(*message, sess, queueURL, messageCounter)

   }

}

func SendByPOSTRequest(message sqs.Message) {

    fmt.Println("\n")
    fmt.Println("NEW MESSAGE:")
    fmt.Println("#############################")
    fmt.Println("Message ID:     " + *message.MessageId)
    fmt.Println("Message Handle: " + *message.ReceiptHandle)
    fmt.Println("Message Body: " + *message.Body)
    fmt.Println("#############################")
    fmt.Println("\n")


}
