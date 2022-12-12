package cmd

import (
	"aoc/day09"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use: "9",
	Run: func(cmd *cobra.Command, args []string) {
		err := day09.Solve(inputFileName, debug)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	solveCmd.AddCommand(cmd)
}
