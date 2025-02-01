package main

import (
	"github.com/spf13/cobra"
)

var (
	id       string
	language string
	rootCmd  = &cobra.Command{
		Use:   "leetcode",
		Short: "Generate leetcode starter code",
	}
)

func init() {
	rootCmd.AddCommand(DailyCmd)
	rootCmd.AddCommand(ProblemCmd)

	DailyCmd.PersistentFlags().StringVarP(&language, "language", "l", "go", "Language")

	ProblemCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "problem id")
	ProblemCmd.MarkPersistentFlagRequired("id")
}

func main() {
	rootCmd.Execute()
}
