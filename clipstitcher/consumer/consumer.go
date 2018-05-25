package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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
	SqsClient  *sqs.SQS
	QueueURL   *string
	RetryCount int
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
	consumerService.RetryCount = 12
	return consumerService, nil
}

func receiveMessageFromQue(cService consumerService) (*sqs.ReceiveMessageOutput, error) {
	waitTime := int64(20)
	resp := &sqs.ReceiveMessageOutput{}
	messageConfig := &sqs.ReceiveMessageInput{
		QueueUrl:        cService.QueueURL,
		WaitTimeSeconds: &waitTime,
	}

	for retry := 1; retry <= cService.RetryCount; retry++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		req, resp := cService.SqsClient.ReceiveMessageRequest(messageConfig)
		req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
		err := req.Send()
		cancel()
		if err == nil && len(resp.Messages) > 0 {
			return resp, err
		} else if retry == cService.RetryCount {
			if err != nil {
				fmt.Println(err.Error())
			}
			err = errors.New("Max retries reached trying to publish sns mesage")
			return resp, err
		}
		time.Sleep(1 * time.Second)
	}
	return resp, nil
}

func formatQueResponse(rawSqs *sqs.ReceiveMessageOutput) (clipMessage, error) {
	cMessage := clipMessage{}
	snsWrapper := snsMessage{}
	wrapperMessage := *rawSqs.Messages[0].Body
	err := json.Unmarshal([]byte(wrapperMessage), &snsWrapper)
	if err != nil {
		return cMessage, err
	}
	actualMessage := snsWrapper.Message
	err = json.Unmarshal([]byte(actualMessage), &cMessage)
	if err != nil {
		return cMessage, err
	}
	cMessage.ReceiptHandle = rawSqs.Messages[0].ReceiptHandle
	return cMessage, nil
}

func (cService consumerService) GetMessage() (clipMessage, error) {
	cMessage := clipMessage{}
	resp, err := receiveMessageFromQue(cService)
	if err != nil {
		return cMessage, err
	}
	cMessage, err = formatQueResponse(resp)
	if err != nil {
		return cMessage, err
	}
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
