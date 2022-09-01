package awsToolBox

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

func createEc2Client() (*ec2.EC2, error) {
	if AWSSession.session == nil {
		err := AWSSession.InitialLogin()
		if err != nil {
			return nil, err
		}
	}

	// Create the EC2 service client.
	svc := ec2.New(AWSSession.session)

	return svc, nil
}
