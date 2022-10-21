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

var vpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "Functions for managing VPCs",
	Long:  `Functions for managing VPCs. This command will use the AWS CLI to login to AWS.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		//if no flags are set, print help
		if !viper.GetBool("vpc.list") && !viper.GetBool("vpc.delete") && !viper.GetBool("vpc.force") {
			cmd.Help()
			os.Exit(0)
		}

		res := awsToolBox.AWSSession.Vpcs()
		if res != nil {

			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(vpcCmd)

	pfs := vpcCmd.PersistentFlags()
	pfs.BoolP("list", "l", false, "List VPCs")
	pfs.BoolP("delete", "d", false, "Delete VPCs")
	pfs.BoolP("force", "f", false, "Attempts to force delete VPCs by deleting all dependencies")

	viper.BindPFlag("vpc.list", pfs.Lookup("list"))
	viper.BindPFlag("vpc.delete", pfs.Lookup("delete"))
	viper.BindPFlag("vpc.force", pfs.Lookup("force"))
}
