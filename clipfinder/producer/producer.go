package producer

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type clipMessage struct {
	VideoSlugs       []string `json:"videoSlugs"`
	VideoDescription string   `json:"videoDescription"`
	ChannelName      string   `json:"channelName"`
}

type producerService struct {
	SqsClient  *sqs.SQS
	QueueURL   *string
	RetryCount int
}

func NewProducerService(queEndpoint string, queueURL string) producerService {
	producerService := producerService{}
	sess := session.Must(session.NewSession())
	sqsClient := &sqs.SQS{}
	if queEndpoint != "" {
		sqsClient = sqs.New(sess, aws.NewConfig().WithEndpoint(queEndpoint))
	} else {
		config := aws.NewConfig()
		sqsClient = sqs.New(sess, config)
	}

	producerService.QueueURL = &queueURL
	producerService.SqsClient = sqsClient
	producerService.RetryCount = 10
	return producerService
}

func (pService producerService) SendMessage(videoSlugs []string, videoDescription string, channelName string) error {
	cMessage := clipMessage{
		VideoSlugs:       videoSlugs,
		VideoDescription: videoDescription,
		ChannelName:      channelName,
	}
	messageBytes, err := json.Marshal(cMessage)
	if err != nil {
		return err
	}
	messageBody := string(messageBytes)

	message := &sqs.SendMessageInput{
		QueueUrl:    pService.QueueURL,
		MessageBody: &messageBody,
	}
	for retry := 1; retry <= pService.RetryCount; retry++ {
		_, err := pService.SqsClient.SendMessage(message)
		if err == nil {
			return nil
		} else if retry == pService.RetryCount {
			err = errors.New("Max retries reached trying to publish sns mesage")
			return err
		}
		fmt.Println("Err from SendMessage sqs " + err.Error())
		fmt.Println("Error trying to push messgage to sqs for " + channelName + " retrying")
		time.Sleep(1 * time.Second)
	}

	return nil
}
