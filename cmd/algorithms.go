package cmd

import (
	"fmt"

	"github.com/hoop33/jwtbot/model"
	"github.com/spf13/cobra"
)

// AlgorithmsCmd lists the supported algorithms
var AlgorithmsCmd = &cobra.Command{
	Use:   "algorithms",
	Short: "list available algorithms",
	Long:  "lists all the available algorithms you can use for signing the JWT",
	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range model.Algorithms {
			fmt.Println(a)
		}
	},
}

func init() {
	RootCmd.AddCommand(AlgorithmsCmd)
}
