package consumer

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type clipMessage struct {
	VideoLinks       []string `json:"videoLinks"`
	VideoDescription string   `json:"videoDescription"`
	ChannelName      string   `json:"channelName "`
	ReceiptHandle    *string
}

type consumerService struct {
	SqsClient *sqs.SQS
	QueueURL  *string
}

func NewConsumerService(queEndpoint string, queueName string) (consumerService, error) {
	consumerService := consumerService{}
	sess := session.Must(session.NewSession())
	sqsClient := &sqs.SQS{}
	if queEndpoint != "" {
		sqsClient = sqs.New(sess, aws.NewConfig().WithEndpoint(queEndpoint))
	} else {
		sqsClient = sqs.New(sess, aws.NewConfig())
	}

	queConfig := &sqs.CreateQueueInput{QueueName: &queueName}
	output, err := sqsClient.CreateQueue(queConfig)
	if err != nil {
		return consumerService, err
	}
	consumerService.QueueURL = output.QueueUrl
	consumerService.SqsClient = sqsClient
	return consumerService, nil
}

func (cService consumerService) GetMessage() (clipMessage, error) {
	cMessage := clipMessage{}
	messageConfig := &sqs.ReceiveMessageInput{
		QueueUrl: cService.QueueURL,
	}

	messageOutput, err := cService.SqsClient.ReceiveMessage(messageConfig)
	if err != nil {
		return cMessage, err
	}
	fmt.Printf("%+v\n", messageOutput.Messages[0])
	if len(messageOutput.Messages) == 0 {
		err = errors.New("No messages found in sqs que")
		return cMessage, err
	}
	cMessage.ReceiptHandle = messageOutput.Messages[0].ReceiptHandle
	message := *messageOutput.Messages[0].Body
	err = json.Unmarshal([]byte(message), &cMessage)
	if err != nil {
		return cMessage, err
	}
	fmt.Println("Message recieved from sqs" + cMessage.ChannelName)
	return cMessage, nil
}

func (cService consumerService) DeleteMessage(message clipMessage) {
	deleteInput := &sqs.DeleteMessageInput{
		QueueUrl:      cService.QueueURL,
		ReceiptHandle: message.ReceiptHandle,
	}
	_, err := cService.SqsClient.DeleteMessage(deleteInput)
	if err != nil {
		fmt.Println("err is" + err.Error())
	}
	fmt.Println("message deleted from sqs")
}
