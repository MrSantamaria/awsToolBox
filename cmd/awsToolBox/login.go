/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package awsToolBox

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to AWS",
	Long:  `login to AWS. This command will use the AWS CLI to login to AWS.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		res := awsToolBox.getSession(args[0], args[1], args[2])
		fmt.Println("aws login called")
	},
}

func init() {
	awsToolBoxCmd.AddCommand(loginCmd)
}
