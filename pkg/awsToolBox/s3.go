package awsToolBox

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (AwsSession *awsSession) S3() error {
	if AwsSession.session == nil {
		err := AwsSession.InitialLogin()
		if err != nil {
			return err
		}
	}

	//AwsSession.ListS3Buckets()

	count, err := AwsSession.GetS3BucketCount()
	if err != nil {
		return err
	}
	fmt.Printf("S3 Bucket Count: %d\n", count)

	AwsSession.DeleteAndCleanSpecificS3Buckets()

	return nil
}

// func (AwsSession *awsSession) GetS3BucketQuotaUsage() (int, int, error) {
// 	bucketCount, err := AwsSession.GetS3BucketCount()
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	quota, err := AwsSession.s3ControlClient.GetBucketQuota(&s3control.GetBucketQuotaInput{})
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	return bucketCount, int(*quota.Quota.BucketQuota), nil
// }

func (AwsSession *awsSession) ListS3Buckets() error {
	input := &s3.ListBucketsInput{}
	result, err := AwsSession.s3.ListBuckets(input)
	if err != nil {
		return err
	}

	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name)
	}

	return nil
}

func (AwsSession *awsSession) GetS3BucketCount() (int, error) {
	input := &s3.ListBucketsInput{}
	result, err := AwsSession.s3.ListBuckets(input)
	if err != nil {
		return 0, err
	}

	return len(result.Buckets), nil
}

func (AwsSession *awsSession) DeleteAndCleanSpecificS3Buckets() error {
	input := &s3.ListBucketsInput{}
	result, err := AwsSession.s3.ListBuckets(input)
	if err != nil {
		return err
	}

	oneMonthsAgo := time.Now().AddDate(0, -1, 0)

	for _, bucket := range result.Buckets {
		if bucket.CreationDate.Before(oneMonthsAgo) {
			// Clean up the contents of the bucket before attempting to delete
			err = AwsSession.cleanUpBucket(*bucket.Name)
			if err != nil {
				//return err
				fmt.Println("Error cleaning up bucket:", *bucket.Name)
				continue
			}

			_, err := AwsSession.s3.DeleteBucket(&s3.DeleteBucketInput{
				Bucket: bucket.Name,
			})
			if err != nil {
				return err
			}

			fmt.Println("Deleted S3 bucket:", *bucket.Name)
		}
	}

	return nil
}

// Function to clean up the contents of an S3 bucket
func (AwsSession *awsSession) cleanUpBucket(bucketName string) error {
	// Get all objects in the bucket
	// Get all objects in the bucket
	err := AwsSession.s3.ListObjectsV2Pages(&s3.ListObjectsV2Input{Bucket: &bucketName},
		func(page *s3.ListObjectsV2Output, lastPage bool) bool {
			// Delete all objects in the bucket
			if len(page.Contents) > 0 {
				deleteObjectsInput := &s3.DeleteObjectsInput{
					Bucket: &bucketName,
					Delete: &s3.Delete{
						Objects: make([]*s3.ObjectIdentifier, len(page.Contents)),
						Quiet:   aws.Bool(true),
					},
				}

				for i, object := range page.Contents {
					deleteObjectsInput.Delete.Objects[i] = &s3.ObjectIdentifier{
						Key: object.Key,
					}
				}

				_, err := AwsSession.s3.DeleteObjects(deleteObjectsInput)
				if err != nil {
					return false
				}
			}

			return true
		})

	if err != nil {
		return err
	}

	return nil
}
