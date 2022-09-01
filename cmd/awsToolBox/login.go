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
	Run: func(cmd *cobra.Command, args []string) {
		res := awsToolBox.AWSSession.InitialLogin()
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
