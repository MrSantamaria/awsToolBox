package awsToolBox

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

// AWS list All the VPCs in the region
func ListVPCs() ([]*ec2.Vpc, error) {
	svc, err := createEc2Client()
	if err != nil {
		return nil, err
	}

	params := &ec2.DescribeVpcsInput{}

	resp, err := svc.DescribeVpcs(params)
	if err != nil {
		return nil, err
	}
	return resp.Vpcs, nil
}
