package awsToolBox

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/fatih/color"
)

func (AwsSession *awsSession) ListVPCs() ([]*ec2.Vpc, error) {
	if AwsSession.session == nil {
		err := AwsSession.InitialLogin()
		if err != nil {
			return nil, err
		}
	}

	input := &ec2.DescribeVpcsInput{}
	result, err := AwsSession.ec2.DescribeVpcs(input)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("VPCs found: ", len(result.Vpcs))
	}

	//Print the names of the VPCs
	for _, vpc := range result.Vpcs {
		var vpcName string
		for _, tag := range vpc.Tags {
			if *tag.Key == "Name" {
				vpcName = *tag.Value

				color.Set(color.FgGreen)
				fmt.Printf("VPC ID: %s, Name: %s\n", *vpc.VpcId, vpcName)
				color.Unset()
				continue
			}
		}
		color.Set(color.FgYellow)
		fmt.Printf("VPC ID: %s, Name: %s\n", *vpc.VpcId, "No name found")
		color.Unset()

		// //Attempt to delete the VPC
		// err := AwsSession.deleteVPC(*vpc.VpcId)
		// if err != nil {
		// 	color.Set(color.FgRed)
		// 	fmt.Println(err)
		// 	color.Unset()
		// }
	}

	return result.Vpcs, nil
}

func (AwsSession *awsSession) deleteVPC(vpcID string) error {
	if AwsSession.session == nil {
		err := AwsSession.InitialLogin()
		if err != nil {
			return err
		}
	}

	input := &ec2.DeleteVpcInput{
		VpcId: &vpcID,
	}
	_, err := AwsSession.ec2.DeleteVpc(input)
	//if err contains dependencyViolation, then delete the subnets
	// if strings.Contains(err.Error(), "DependencyViolation") {
	// 	//Delete the subnets
	// 	err := AwsSession.deleteDependencySubnet(vpcID)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	//Try deleting the VPC again
	// 	err = AwsSession.deleteVPC(vpcID)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	if err != nil {
		return err
	}

	return nil
}

func (AwsSession *awsSession) deleteDependencySubnet(vpcID string) error {
	if AwsSession.session == nil {
		err := AwsSession.InitialLogin()
		if err != nil {
			return err
		}
	}

	//Get the subnets based on the VPC ID
	input := &ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []*string{aws.String(vpcID)},
			},
		},
	}

	result, err := AwsSession.ec2.DescribeSubnets(input)
	if err != nil {
		return err
	}

	//Print the names of the subnets
	for _, subnet := range result.Subnets {
		var subnetName string
		for _, tag := range subnet.Tags {
			if *tag.Key == "Name" {
				subnetName = *tag.Value

				fmt.Printf("Subnet ID: %s, Name: %s\n", *subnet.SubnetId, subnetName)
			}
		}

	}

	return nil

}
