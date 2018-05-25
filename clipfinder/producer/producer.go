package producer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type clipMessage struct {
	VideoSlugs       []string `json:"videoSlugs"`
	VideoDescription string   `json:"videoDescription"`
	ChannelName      string   `json:"channelName"`
}

type producerService struct {
	SnsClient    *sns.SNS
	TopicArn     *string
	RetriesCount int
}

func NewProducerService(endpoint string, arn string) producerService {
	pService := producerService{}
	sess := session.Must(session.NewSession())
	snsClient := &sns.SNS{}
	if endpoint != "" {
		snsClient = sns.New(sess, aws.NewConfig().WithEndpoint(endpoint))
	} else {
		snsClient = sns.New(sess, aws.NewConfig())
	}

	pService.SnsClient = snsClient
	pService.TopicArn = &arn
	pService.RetriesCount = 30
	return pService
}

func dummyStub() error {
	return errors.New("Test error")
}

func (pService producerService) CheckSubscriptions(customEndPoint string) error {
	check := &sns.ListSubscriptionsByTopicInput{
		TopicArn: pService.TopicArn,
	}

	for retry := 1; retry <= pService.RetriesCount; retry++ {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		req, resp := pService.SnsClient.ListSubscriptionsByTopicRequest(check)
		req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
		err := req.Send()
		cancel()
		if err == nil && len(resp.Subscriptions) > 0 {
			return nil
		} else if retry == pService.RetriesCount {
			err = errors.New("Max retries hit trying to see subscriptions")
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
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

	message := &sns.PublishInput{
		TopicArn: pService.TopicArn,
		Message:  &messageBody,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for retry := 1; retry <= pService.RetriesCount; retry++ {
		req, _ := pService.SnsClient.PublishRequest(message)
		req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
		err := req.Send()
		if err == nil {
			return nil
		} else if retry == pService.RetriesCount {
			err = errors.New("Max retries reached trying to publish sns mesage")
			return err
		}
		fmt.Println("Error trying to publish to sns for " + channelName + " retrying")
		time.Sleep(1 * time.Second)
	}

	return nil
}
