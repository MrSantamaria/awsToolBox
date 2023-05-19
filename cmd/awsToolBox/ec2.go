/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package awsToolBox

import (
	"fmt"
	"os"

	"github.com/MrSantamaria/awsToolBox/pkg/awsToolBox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Functions for managing ec2 resources",
	Long:  `Functions for managing ec2 resources. This command will use the AWS CLI to login to AWS.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		//if no flags are set, print help
		if !viper.GetBool("ec2.snapshots") {
			cmd.Help()
			os.Exit(0)
		}

		res := awsToolBox.AWSSession.Ec2()
		if res != nil {

			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(ec2Cmd)

	pfs := ec2Cmd.PersistentFlags()
	pfs.BoolP("snapshots", "s", false, "Used to manage ec2 snapshots")

	viper.BindPFlag("ec2.snapshots", pfs.Lookup("snapshots"))
}
