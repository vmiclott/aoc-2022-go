package cmd

import (
	"aoc/day10"

	"github.com/spf13/cobra"
)

var cmd10 = &cobra.Command{
	Use:   "10",
	Short: "Solve problem 10",
	Run: func(cmd *cobra.Command, args []string) {
		err := day10.Solve(inputFileName, debug)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	solveCmd.AddCommand(cmd10)
}
