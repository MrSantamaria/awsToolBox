/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package awsToolBox

import (
	"fmt"

	"github.com/MrSantamaria/awsToolBox/pkg/awsToolBox"
	"github.com/spf13/cobra"
)

var (
	awsAccessKey string
	awsSecretKey string
	awsRegion    string
	viperEnv     bool
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to AWS",
	Long:  `login to AWS. This command will use the AWS CLI to login to AWS.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := awsToolBox.AWSSession.InitialLogin(args[0], args[1], args[2])
		fmt.Println(res)
	},
}

func init() {
	loginCmd.Flags().StringVarP(&awsAccessKey, "accessKey", "a", "", "AWS Access Key")
	loginCmd.Flags().StringVarP(&awsSecretKey, "secretKey", "s", "", "AWS Secret Key")
	loginCmd.Flags().StringVarP(&awsRegion, "region", "r", "", "AWS Region")
	loginCmd.Flags().BoolVarP(&viperEnv, "viperEnv", "v", false, "Use Viper to get AWS credentials from environment variables")
	rootCmd.AddCommand(loginCmd)
}
