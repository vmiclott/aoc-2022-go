package cmd

import (
	"github.com/spf13/cobra"
)

var inputFileName string
var debug bool

var solveCmd = &cobra.Command{
	Use: "solve",
}

func init() {
	rootCmd.AddCommand(solveCmd)
	solveCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable to print debugging messages")
	solveCmd.PersistentFlags().StringVarP(&inputFileName, "input", "i", "", "input file")
	solveCmd.MarkPersistentFlagRequired("input")
}
