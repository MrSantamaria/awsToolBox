/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package awsToolBox

import (
	"fmt"

	"github.com/MrSantamaria/awsToolBox/pkg/awsToolBox"
	"github.com/spf13/cobra"
)

var vpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "Functions for managing VPCs",
	Long:  `Functions for managing VPCs. This command will use the AWS CLI to login to AWS.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res := awsToolBox.AWSSession.ListVPCs()
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(vpcCmd)
}
