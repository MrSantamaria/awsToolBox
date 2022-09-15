package awsToolBox

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/spf13/viper"
)

type awsSession struct {
	session *session.Session
	iam     *iam.IAM
	ec2     *ec2.EC2
	once    sync.Once
}

// AwsSession is the global AWS session for interacting with AWS.
var AWSSession awsSession

func (AWSSession *awsSession) InitialLogin() error {
	if !viper.IsSet("aws.awsAccessKey") || !viper.IsSet("aws.awsSecretAccessKey") || !viper.IsSet("aws.region") {
		return fmt.Errorf("AWS credentials not set")
	}

	err := AWSSession.getSession(viper.GetString("aws.awsAccessKey"), viper.GetString("aws.awsSecretAccessKey"), viper.GetString("aws.region"))
	if err != nil {
		return err
	}
	return nil
}

func (AWSSession *awsSession) getSession(awsAccessKey, awsSecretAccessKey, awsRegion string) error {
	var err error

	AWSSession.once.Do(func() {
		// Create the AWS session
		AWSSession.session, err = session.NewSession(aws.NewConfig().
			WithCredentials(credentials.NewStaticCredentials(awsAccessKey, awsSecretAccessKey, "")).
			WithRegion(awsRegion))

		AWSSession.iam = iam.New(AWSSession.session)
		AWSSession.ec2 = ec2.New(AWSSession.session)

		if err != nil {
			log.Printf("error initializing AWS session: %v", err)
		}
	})

	if AWSSession.session == nil {
		err = fmt.Errorf("unable to initialize AWS session")
	}

	return err
}
