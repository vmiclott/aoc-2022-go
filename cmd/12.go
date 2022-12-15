package cmd

import (
	"aoc/day12"

	"github.com/spf13/cobra"
)

var cmd12 = &cobra.Command{
	Use:   "12",
	Short: "Solve problem 12",
	Run: func(cmd *cobra.Command, args []string) {
		err := day12.Solve(inputFileName, debug)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	solveCmd.AddCommand(cmd12)
}
