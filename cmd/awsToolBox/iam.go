/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package awsToolBox

import (
	"fmt"

	"github.com/MrSantamaria/awsToolBox/pkg/awsToolBox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "Functions for managing IAMs",
	Long:  `Functions for managing IAMs. This command will use the AWS CLI to login to AWS.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		res := awsToolBox.AWSSession.Iam()
		if res != nil {

			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(iamCmd)

	pfs := iamCmd.PersistentFlags()
	pfs.BoolP("list", "l", false, "List Iams")

	viper.BindPFlag("iam.list", pfs.Lookup("list"))
}
