package nodegroup

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

type StateNode struct {
	ClusterName   string
	NodegroupName string
	RoleArn       string
}

func CreateNodeGroup(state *StateNode) {

	// session := session.New(&aws.Config{
	// 	Region:      aws.String("us-west-2"),
	// 	Credentials: credentials.NewEnvCredentials(),
	// })

	eksclient := eks.New(session.New())

	nodegroupstate := &eks.CreateNodegroupInput{
		ClusterName:   aws.String(state.ClusterName),
		NodegroupName: aws.String(state.NodegroupName),

		Subnets: []*string{
			aws.String("subnet-3419e54c"),
			aws.String("subnet-44647d0f"),
			aws.String("subnet-4c78b711"),
			aws.String("subnet-fb4d0bd0"),
		},

		NodeRole: aws.String(state.RoleArn),
	}

	result, err := eksclient.CreateNodegroup(nodegroupstate)

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

func UpdateNodeGroup() {

	eksclient := eks.New(session.New())

	nodegroupstate := &eks.UpdateNodegroupConfigInput{
		ClusterName:   aws.String("novoclust"),
		NodegroupName: aws.String("tessteglobo"),
		Labels: &eks.UpdateLabelsPayload{
			AddOrUpdateLabels: map[string]*string{
				"nodegroupName": aws.String("novogloboow"),
			},
		},

		ScalingConfig: &eks.NodegroupScalingConfig{
			DesiredSize: aws.Int64(4),
			MaxSize:     aws.Int64(4),
			MinSize:     aws.Int64(4),
		},
	}

	result, err := eksclient.UpdateNodegroupConfig(nodegroupstate)

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

func ListNodegroups() {

	svc := eks.New(session.New())
	state := &eks.ListNodegroupsInput{
		ClusterName: aws.String("novoclust"),
	}

	result, err := svc.ListNodegroups(state)
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
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}
