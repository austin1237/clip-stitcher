package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

var taskArn, clusterArn, subnetId string

func init() {
	taskArn = os.Getenv("TASK_ARN")
	clusterArn = os.Getenv("CLUSTER_ARN")
	subnetId = os.Getenv("SUBNET_ID")
	if taskArn == "" {
		log.Fatal("TASK_ARN was not set as an env var.")
		os.Exit(1)
	}

	if clusterArn == "" {
		log.Fatal("CLUSTER_ARN was not set as an env var.")
		os.Exit(1)
	}

	if subnetId == "" {
		log.Fatal("VPC_ID was not set as an env var.")
		os.Exit(1)
	}
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (string, error) {
	launchType := "FARGATE"
	sess := session.Must(session.NewSession())
	ecsClient := ecs.New(sess)
	vpcConfig := ecs.AwsVpcConfiguration{
		Subnets: []*string{&subnetId},
	}
	networkConfig := ecs.NetworkConfiguration{
		AwsvpcConfiguration: &vpcConfig,
	}
	runTask := ecs.RunTaskInput{
		Cluster:              &clusterArn,
		TaskDefinition:       &taskArn,
		NetworkConfiguration: &networkConfig,
		LaunchType:           &launchType,
	}
	_, err := ecsClient.RunTask(&runTask)

	if err != nil {
		return fmt.Sprintln("An error occured " + err.Error()), nil
	}

	return fmt.Sprintf("ECS Task Ran"), nil
}
