/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package awsToolBox

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var vesrion = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "awsToolBox",
	Version: vesrion,
	Short:   "awsToolBox is a CLI tool for managing AWS resources",
	Long: `awsToolBox is a CLI tool for managing AWS resources.
			This CLI is a work in progress and will be updated frequently.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aws login called")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
