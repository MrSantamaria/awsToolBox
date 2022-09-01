package awsToolBox

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/viper"
)

type awsSession struct {
	session *session.Session
	once    sync.Once
}

// AwsSession is the global AWS session for interacting with AWS.
var AWSSession awsSession

func (a *awsSession) InitialLogin() error {
	_, err := a.getSession(viper.GetString("aws.awsAccessKey"), viper.GetString("aws.awsSecretKey"), viper.GetString("aws.region"))
	return err
}

func (a *awsSession) getSession(awsAccessKey, awsSecretAccessKey, awsRegion string) (*session.Session, error) {
	var err error

	a.once.Do(func() {
		// Create the AWS session
		a.session, err = session.NewSession(aws.NewConfig().
			WithCredentials(credentials.NewStaticCredentials(awsAccessKey, awsSecretAccessKey, "")).
			WithRegion(awsRegion))

		if err != nil {
			log.Printf("error initializing AWS session: %v", err)
		}
	})

	if a.session == nil {
		err = fmt.Errorf("unable to initialize AWS session")
	}

	//print a.session
	fmt.Println(a.session.Config.Credentials)

	return a.session, err
}
