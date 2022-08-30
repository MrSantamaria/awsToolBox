package awsToolBox

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsSession struct {
	session *session.Session
	once    sync.Once
}

// AwsSession is the global AWS session for interacting with AWS.
var AWSSession awsSession

func (a *awsSession) getSession(awsAccessKey string, awsSecretKey string, awsRegion string) (*session.Session, error) {
	var err error

	a.once.Do(func() {
		// We're using static credentials here so that we can use AWS credentials for cluster providers.
		// When we have more time, we should make this not metrics focused, as the intent of this library is to be purpose agnostic.
		a.session, err = session.NewSession(aws.NewConfig().
			WithCredentials(credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")).
			WithRegion(awsRegion))

		if err != nil {
			log.Printf("error initializing AWS session: %v", err)
		}
	})

	if a.session == nil {
		err = fmt.Errorf("unable to initialize AWS session")
	}

	return a.session, err
}
