/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package awsToolBox

import (
	"fmt"

	"github.com/MrSantamaria/awsToolBox/pkg/awsToolBox"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Verify AWS credentials",
	Long:  `Login to AWS. This command will use the AWS CLI to login to AWS.`,
	Run: func(cmd *cobra.Command, args []string) {
		res := awsToolBox.AWSSession.InitialLogin()
		if res != nil {
			fmt.Println(res)
		} else {
			fmt.Println("Login successful")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
