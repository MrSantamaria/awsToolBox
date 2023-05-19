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

var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Functions for managing s3",
	Long:  `Functions for managing s3. This command can be used to interact with s3.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		res := awsToolBox.AWSSession.S3()
		if res != nil {

			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)

	pfs := s3Cmd.PersistentFlags()
	pfs.BoolP("list", "l", false, "List s3 buckets")
	pfs.BoolP("delete", "d", false, "List s3 buckets")

	viper.BindPFlag("s3.list", pfs.Lookup("list"))
	viper.BindPFlag("s3.delete", pfs.Lookup("delete"))
}
