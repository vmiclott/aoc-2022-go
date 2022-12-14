package cmd

import (
	"aoc/day11"

	"github.com/spf13/cobra"
)

var cmd11 = &cobra.Command{
	Use:   "11",
	Short: "Solve problem 11",
	Run: func(cmd *cobra.Command, args []string) {
		err := day11.Solve(inputFileName, debug)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	solveCmd.AddCommand(cmd11)
}
