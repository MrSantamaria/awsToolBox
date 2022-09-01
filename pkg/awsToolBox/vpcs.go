package awsToolBox

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//Print out the VPCs in the account and region
func (a *awsSession) ListVpcs() error {
	if a.session == nil {
		return fmt.Errorf("AWS session not initialized")
	}

	// Create the EC2 service client.
	svc := ec2.New(a.session)

	// Get the list of VPCs
	result, err := svc.DescribeVpcs(nil)
	if err != nil {
		return err
	}

	// Print out the VPCs
	fmt.Println("VPCs:")
	for _, vpc := range result.Vpcs {
		fmt.Printf("%s (%s)\n", aws.StringValue(vpc.VpcId), aws.StringValue(vpc.State))
	}

	return nil
}
