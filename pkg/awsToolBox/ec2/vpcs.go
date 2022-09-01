package awsToolBox

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

//Print out the VPCs in the account and region
func (a *awsSession) ListVpcs() error {
	if a.session == nil {
		err := a.InitialLogin()
		if err != nil {
			return err
		}
	}

	// Create the EC2 service client.
	svc := ec2.New(a.session)

	// Get the list of VPCs
	result, err := svc.DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err != nil {
		return err

	}

}

// Get all of the VPCs configured in the environment
func getAllVPCs(ec2client *ec2.EC2) ([]*ec2.Vpc, error) {
	//Get all of the VPCs
	vpcs, err := ec2client.DescribeVpcs(&ec2.DescribeVpcsInput{})

	//If we had an error, return it
	if err != nil {
		return []*ec2.Vpc{}, err
	}

	//Otherwise, return all of our VPCs
	return vpcs.Vpcs, nil
}
