package sqs

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//func main() {

//	const (
//		QueueUrl    = "https://sqs.us-west-2.amazonaws.com/224086203907/rb_data_scale"
//		Region      = "us-west-2"
//		CredPath    = "/Users/admin/.aws/credentials"
//		CredProfile = "default"
//	)
//
//var a *string = "t"
//attributesOutput := map[string]*string{}
//svc := sqs.New(Session())
//svcout, err := svc.GetQueueAttributes(a)

//	fmt.Println(a, svcout, svc, err)

//	a := &sqs.GetQueueAttributesInput{
//		AttributeNames: []*string{aws.String("ApproximateNumberOfMessages")},
//		QueueUrl:       aws.String(QueueUrl),
//	fmt.Println(qcount())
//}

func Sess(Region string, configpath string, profile string) *session.Session {
	//s := Session.New()
	sess := aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewSharedCredentials(configpath, profile),
		MaxRetries:  aws.Int(5),
	}

	clinet := session.New(&sess)
	return clinet

}

func Qcount(QUrl string) int {

	input := &sqs.GetQueueAttributesInput{
		AttributeNames: []*string{aws.String("ApproximateNumberOfMessages")},
		QueueUrl:       aws.String(QUrl),
	}
	svc := sqs.New(Sess("us-west-2", "/Users/admin/.aws/config", "default"))
	svcout, err := svc.GetQueueAttributes(input)
	if err != nil {
		fmt.Println(err, "Failed to get messages in SQS")
	}
	messages, err := strconv.Atoi(*svcout.Attributes["ApproximateNumberOfMessages"])
	if err != nil {
		fmt.Println(err, "Failed to get number of messages in queue")
	}
	return messages
}
