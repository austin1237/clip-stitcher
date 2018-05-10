package producer

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type clipMessage struct {
	VideoLinks       []string `json:"videoLinks"`
	VideoDescription string   `json:"videoDescription"`
	ChannelName      string   `json:"channelName"`
}

type producerService struct {
	SnsClient *sns.SNS
	TopicArn  *string
}

func NewProducerService(endpoint string, name string) (producerService, error) {
	pService := producerService{}
	sess := session.Must(session.NewSession())
	snsClient := &sns.SNS{}
	if endpoint != "" {
		snsClient = sns.New(sess, aws.NewConfig().WithEndpoint(endpoint))
	} else {
		snsClient = sns.New(sess, aws.NewConfig())
	}

	topicConfig := &sns.CreateTopicInput{Name: &name}
	output, err := snsClient.CreateTopic(topicConfig)
	if err != nil {
		return pService, err
	}
	topicArn := output.TopicArn
	pService.SnsClient = snsClient
	pService.TopicArn = topicArn
	return pService, nil
}

func (pService producerService) SendMessage(videoLinks []string, videoDescription string, channelName string) error {
	cMessage := clipMessage{
		VideoLinks:       videoLinks,
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
	_, err = pService.SnsClient.Publish(message)
	if err != nil {
		return err
	}
	return nil
}
