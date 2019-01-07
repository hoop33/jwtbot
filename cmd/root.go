package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the root command for jwtbot
var RootCmd = &cobra.Command{
	Use:   "jwtbot",
	Short: "A CLI for creating JWTs",
	Long:  "jwtbot lets you create JWTs",
}

// Execute adds all child commands to the root command and sets flag appropriately
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
