package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"

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

func removeTags(content string) string {
	r := regexp.MustCompile("(?im)(<.*?>)")

	return r.ReplaceAllString(content, "")
}

func writeProblemFile(path string, dirName string, content string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dirName, os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(path, []byte(content), 0644)

		if err != nil {
			log.Fatal(err)
		}

		return
	}
	println("Files already created")

}

func generateFiles(question Question) {
	dirName := filepath.Join("problems/", question.QuestionFrontendId+"."+question.TitleSlug)
	fileName := question.TitleSlug + ".go"
	descriptionName := "description.txt"

	codePath := filepath.Join(dirName, fileName)
	descriptionPath := filepath.Join(dirName, descriptionName)

	snippet := getStartingCode(question.CodeSnippets)
	description := removeTags(question.Content)

	writeProblemFile(codePath, dirName, snippet)
	writeProblemFile(descriptionPath, dirName, description)
}

var ProblemCmd = &cobra.Command{
	Use:   "problem",
	Short: "Get problem by title slug",
	Run: func(cmd *cobra.Command, args []string) {
		println("Generating leetcode problem...")
		question := GetProblemBySearchInput(id)

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
