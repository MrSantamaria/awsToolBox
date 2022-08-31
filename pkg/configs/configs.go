package config

import (
	"github.com/spf13/viper"
)

func initViper() {

	// Grab AWS credentials from environment variables
	viper.BindEnv("aws.awsAccessKey", "AWS_ACCESS_KEY")
	viper.BindEnv("aws.awsSecretKey", "AWS_SECRET_KEY")
	viper.BindEnv("aws.region", "AWS_REGION")

}