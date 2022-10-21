package awsToolBox

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var (
	wg sync.WaitGroup
)

func (AwsSession *awsSession) Vpcs() error {
	if AwsSession.session == nil {
		err := AwsSession.InitialLogin()
		if err != nil {
			return err
		}
	}

	input := &ec2.DescribeVpcsInput{}
	result, err := AwsSession.ec2.DescribeVpcs(input)
	if err != nil {
		return err
	}
	fmt.Println("VPCs found: ", len(result.Vpcs))

	//Print the names of the VPCs
vpcs:
	for _, vpc := range result.Vpcs {
		var vpcName string
		//If there exists a tag with the key "Name", then print the value

		for _, tag := range vpc.Tags {
			if *tag.Key == "Name" {
				vpcName = *tag.Value

				if viper.GetBool("vpc.list") {
					color.Set(color.FgGreen)
					log.Printf("VPC ID: %s, Name: %s\n", *vpc.VpcId, vpcName)
					color.Unset()
				}
				continue vpcs
			}
		}
		color.Set(color.FgYellow)
		log.Printf("VPC ID: %s, Name: %s\n", *vpc.VpcId, "No name found")
		color.Unset()

		if viper.GetBool("vpc.delete") {
			wg.Add(1)

			go func(vpcId string) {
				defer wg.Done()
				err := AwsSession.deleteVPC(vpcId)
				if err != nil {
					color.Set(color.FgRed)
					log.Println(err)
					color.Unset()
				}
			}(*vpc.VpcId)

			break vpcs

		}
	}
	wg.Wait()
	return nil
}

func (AwsSession *awsSession) deleteVPC(vpcId string) error {
	input := &ec2.DeleteVpcInput{
		VpcId: aws.String(vpcId),
	}

	_, err := AwsSession.ec2.DeleteVpc(input)
	//if err contains dependencyViolation, then delete the subnets
	if viper.GetBool("vpc.force") && strings.Contains(err.Error(), "DependencyViolation") {
		//Delete the subnets
		err := AwsSession.deleteDependencySubnet(vpcId)
		if err != nil {
			fmt.Println(err)
		}
		//Try deleting the VPC again
		// err = AwsSession.deleteVPC(vpcId)
		// if err != nil {
		// 	return err
		// }

	}
	if err != nil {
		return err
	}
	color.Set(color.FgGreen)
	log.Printf("VPC %s deleted\n", vpcId)
	color.Unset()

	return nil
}

func (AwsSession *awsSession) deleteDependencySubnet(vpcID string) error {

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

	//Delete the subnets
	for _, subnet := range result.Subnets {
		err := AwsSession.deleteVPCSubnet(vpcID, *subnet.SubnetId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (AwsSession *awsSession) deleteVPCSubnet(vpcID string, subnetId string) error {
	input := &ec2.DeleteSubnetInput{
		SubnetId: aws.String(subnetId),
	}

	_, err := AwsSession.ec2.DeleteSubnet(input)
	if err != nil {
		return err
	}
	color.Set(color.FgGreen)
	log.Printf("Subnet dependency %s for VPC %s was deleted. \n", subnetId, vpcID)
	color.Unset()

	return nil
}
