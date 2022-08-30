/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package awsToolBox

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// awsCmd represents the aws command
var awsToolBoxCmd = &cobra.Command{
	Use:   "awsToolBox",
	Short: "awsToolBox is a CLI tool for managing AWS resources",
	Long: `awsToolBox is a CLI tool for managing AWS resources.
			This CLI is a work in progress and will be updated frequently.`,
	Args: cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aws login called")
	},
}

func Execute() {
	if err := awsToolBoxCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
