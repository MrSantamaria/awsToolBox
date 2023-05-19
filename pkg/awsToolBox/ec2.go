package awsToolBox

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (AwsSession *awsSession) Ec2() error {
	if AwsSession.session == nil {
		err := AwsSession.InitialLogin()
		if err != nil {
			return err
		}
	}

	AwsSession.DeleteOldSnapshots()
	//AwsSession.listSnapshots()

	return nil
}

func (AwsSession *awsSession) listSnapshots() error {
	// Describe all EBS snapshots
	resp, err := AwsSession.ec2.DescribeSnapshots(nil)
	if err != nil {
		return err
	}

	// Print table header
	fmt.Printf("%-20s %-20s %-20s %-20s\n", "SNAPSHOT ID", "START TIME", "VOLUME ID", "VOLUME TYPE")

	// Print table rows
	for _, snapshot := range resp.Snapshots {
		fmt.Printf("%-20s %-20s %-20s %-20s\n", *snapshot.SnapshotId, *snapshot.StartTime, *snapshot.VolumeId, *snapshot.VolumeSize)
	}

	return nil
}

func (AwsSession *awsSession) DeleteOldSnapshots() error {
	// Describe all EBS snapshots
	resp, err := AwsSession.ec2.DescribeSnapshots(nil)
	if err != nil {
		return err
	}

	// Iterate over snapshots and delete those older than 5 days
	for _, snapshot := range resp.Snapshots {
		if time.Since(*snapshot.StartTime) > 5*24*time.Hour && *snapshot.State == "completed" {
			fmt.Printf("Deleting snapshot %s (%s) created on %s\n", *snapshot.SnapshotId, *snapshot.VolumeId, *snapshot.StartTime)

			// Delete the snapshot
			_, err := AwsSession.ec2.DeleteSnapshot(&ec2.DeleteSnapshotInput{
				SnapshotId: aws.String(*snapshot.SnapshotId),
			})
			if err != nil {
				fmt.Printf("Error deleting snapshot %s: %s\n", *snapshot.SnapshotId, err)
			}

			fmt.Printf("Deleted snapshot %s\n", *snapshot.SnapshotId)
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}
