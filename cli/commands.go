package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func getStartingCode(codeSnippets []CodeSnippets) string {
	var code string

	for _, snippet := range codeSnippets {
		if snippet.Lang == "Go" {
			code = snippet.Code
			break
		}
	}

	return code

}

func generateFiles(question Question) {
	dirName := filepath.Join("problems/", question.QuestionFrontendId+"."+question.TitleSlug)
	fileName := question.TitleSlug + ".go"
	path := filepath.Join(dirName, fileName)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dirName, os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}

		codeSnippet := getStartingCode(question.CodeSnippets)

		err = os.WriteFile(path, []byte(codeSnippet), 0644)

		if err != nil {
			log.Fatal(err)
		}

		return
	}
	println("Files already created")
}

var ProblemCmd = &cobra.Command{
	Use:   "problem",
	Short: "Get leetcode daily",
	Run: func(cmd *cobra.Command, args []string) {
		println("Generating leetcode problem...", id)
		question := GetProblemByTitleSlug(id)

		generateFiles(question)
	},
}

var DailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Get leetcode daily",
	Run: func(cmd *cobra.Command, args []string) {
		println("Generating leetcode daily...")
		result := GetDailyProblem()

		generateFiles(result.ActiveDailyCodingChallengeQuestion.Question)

	},
}
