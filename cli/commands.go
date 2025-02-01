package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

type Language string

const (
	Golang     Language = "golang"
	Typescript Language = "typescript"
)

func MapStringToLanguage(lang string) (Language, error) {
	switch lang {
	case "go":
		return Golang, nil
	case "typescript":
		return Typescript, nil
	default:
		return "", errors.New("Language not supported")
	}
}

func (lang Language) String() string {
	switch lang {
	case Golang:
		return "Go"
	case Typescript:
		return "TypeScript"
	default:
		return ""
	}
}

func (lang Language) AddExtension(fileName string) string {
	switch lang {
	case Golang:
		return fileName + ".go"
	case Typescript:
		return fileName + ".ts"
	default:
		return fileName
	}
}

func getStartingCode(codeSnippets []CodeSnippets, lang Language) string {
	var code string

	for _, snippet := range codeSnippets {
		if snippet.Lang == lang.String() {
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

		err = os.WriteFile(path, []byte(content), 0o644)
		if err != nil {
			log.Fatal(err)
		}

		return
	}
	println("Files already created")
}

func generateFiles(question *Question, lang Language) {
	dirName := filepath.Join("problems/", question.QuestionFrontendId+"."+question.TitleSlug)
	fileName := lang.AddExtension("main")
	descriptionName := "description.txt"

	codePath := filepath.Join(dirName, fileName)
	descriptionPath := filepath.Join(dirName, descriptionName)

	snippet := getStartingCode(question.CodeSnippets, lang)
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

		lang, err := MapStringToLanguage(language)
		if err != nil {
			panic(err)
		}

		generateFiles(question, lang)
	},
}

var DailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Get leetcode daily",
	Run: func(cmd *cobra.Command, args []string) {
		println("Generating leetcode daily...")
		result := GetDailyProblem()

		lang, err := MapStringToLanguage(language)
		if err != nil {
			panic(err)
		}

		generateFiles(result.ActiveDailyCodingChallengeQuestion.Question, lang)
	},
}
