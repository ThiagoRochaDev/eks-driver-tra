package cluster

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

type State struct {
	ClusterName     string
	SecurityGroupId string
	RoleArn         string
	Version         string
}

func CreateCluster(States *State) {

	eksclient := eks.New(session.New())

	clusterstate := &eks.CreateClusterInput{
		Name: aws.String(States.ClusterName),
		ResourcesVpcConfig: &eks.VpcConfigRequest{
			SecurityGroupIds: []*string{
				aws.String(States.SecurityGroupId),
			},
			SubnetIds: []*string{
				aws.String("subnet-3419e54c"),
				aws.String("subnet-44647d0f"),
				aws.String("subnet-4c78b711"),
				aws.String("subnet-fb4d0bd0"),
			},
		},
		RoleArn: aws.String(States.RoleArn),
	}

	result, err := eksclient.CreateCluster(clusterstate)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeResourceInUseException:
				fmt.Println(eks.ErrCodeResourceInUseException, aerr.Error())
			case eks.ErrCodeResourceLimitExceededException:
				fmt.Println(eks.ErrCodeResourceLimitExceededException, aerr.Error())
			case eks.ErrCodeInvalidParameterException:
				fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
			case eks.ErrCodeClientException:
				fmt.Println(eks.ErrCodeClientException, aerr.Error())
			case eks.ErrCodeServerException:
				fmt.Println(eks.ErrCodeServerException, aerr.Error())
			case eks.ErrCodeServiceUnavailableException:
				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
			case eks.ErrCodeUnsupportedAvailabilityZoneException:
				fmt.Println(eks.ErrCodeUnsupportedAvailabilityZoneException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {

			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}

func UpdateCluster() {

	fmt.Println("not implemented")
}

func ListClusters() {

	eksC := eks.New(session.New())
	state := &eks.ListClustersInput{}

	result, err := eksC.ListClusters(state)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeInvalidParameterException:
				fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
			case eks.ErrCodeClientException:
				fmt.Println(eks.ErrCodeClientException, aerr.Error())
			case eks.ErrCodeServerException:
				fmt.Println(eks.ErrCodeServerException, aerr.Error())
			case eks.ErrCodeServiceUnavailableException:
				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {

			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}
