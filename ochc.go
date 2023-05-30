/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ZhangSetSail/OChc/cmd"
	"os"
)

func main() {
	oCmd := cmd.NewCMD()
	err := oCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
