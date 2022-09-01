package awsToolBox

import "github.com/aws/aws-sdk-go/service/ec2"

func (a *awsSession) createEC2Client() (*ec2.EC2, error) {
	if a.session == nil {
		err := a.InitialLogin()
		if err != nil {
			return nil, err
		}
	}

	// Create the EC2 service client.
	svc := ec2.New(a.session)

	return svc, nil
}
