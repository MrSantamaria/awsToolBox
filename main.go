/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/MrSantamaria/awsToolBox/cmd/awsToolBox"
	"github.com/MrSantamaria/awsToolBox/pkg/configs"
)

func main() {
	configs.InitViper()
	awsToolBox.Execute()
}
