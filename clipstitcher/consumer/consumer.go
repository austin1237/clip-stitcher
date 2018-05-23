package consumer

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type snsMessage struct {
	Message string `json:"Message"`
}

type clipMessage struct {
	VideoLinks       []string `json:"videoLinks"`
	VideoDescription string   `json:"videoDescription"`
	ChannelName      string   `json:"channelName"`
	ReceiptHandle    *string
}

type consumerService struct {
	SqsClient *sqs.SQS
	QueueURL  *string
}

func NewConsumerService(queEndpoint string, queueURL string) (consumerService, error) {
	consumerService := consumerService{}
	sess := session.Must(session.NewSession())
	sqsClient := &sqs.SQS{}
	if queEndpoint != "" {
		sqsClient = sqs.New(sess, aws.NewConfig().WithEndpoint(queEndpoint))
	} else {
		sqsClient = sqs.New(sess, aws.NewConfig())
	}

	consumerService.QueueURL = &queueURL
	consumerService.SqsClient = sqsClient
	return consumerService, nil
}

func (cService consumerService) GetMessage() (clipMessage, error) {
	cMessage := clipMessage{}
	snsWrapper := snsMessage{}
	messageConfig := &sqs.ReceiveMessageInput{
		QueueUrl: cService.QueueURL,
	}

	messageOutput, err := cService.SqsClient.ReceiveMessage(messageConfig)
	if err != nil {
		return cMessage, err
	}
	if len(messageOutput.Messages) == 0 {
		err = errors.New("No messages found in sqs que")
		return cMessage, err
	}
	wrapperMessage := *messageOutput.Messages[0].Body
	err = json.Unmarshal([]byte(wrapperMessage), &snsWrapper)
	if err != nil {
		return cMessage, err
	}
	actualMessage := snsWrapper.Message
	err = json.Unmarshal([]byte(actualMessage), &cMessage)
	if err != nil {
		return cMessage, err
	}
	cMessage.ReceiptHandle = messageOutput.Messages[0].ReceiptHandle
	return cMessage, nil
}

func (cService consumerService) DeleteMessage(message clipMessage) error {
	deleteInput := &sqs.DeleteMessageInput{
		QueueUrl:      cService.QueueURL,
		ReceiptHandle: message.ReceiptHandle,
	}
	_, err := cService.SqsClient.DeleteMessage(deleteInput)
	if err != nil {
		return err
	}
	return nil
}
