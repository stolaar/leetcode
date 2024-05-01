package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var client = graphql.NewClient("https://leetcode.com/graphql")

type CodeSnippets struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
}

type Question struct {
	QuestionFrontendId string         `json:"questionFrontendId"`
	QuestionId         string         `json:"questionId"`
	TitleSlug          string         `json:"titleSlug"`
	CodeSnippets       []CodeSnippets `json:"codeSnippets"`
}

type Problem struct {
  Question Question `json:"question"`
}

type DailyResponse struct {
	ActiveDailyCodingChallengeQuestion struct {
		Question Question `json:"question"`
	} `json:"activeDailyCodingChallengeQuestion"`
}

func getProblemByTitleSlug(titleSlug string) Question {
	request := graphql.NewRequest(`
  query selectProblem($titleSlug: String!) {
    question(titleSlug: $titleSlug) {
      questionId
      questionFrontendId
      titleSlug
      codeSnippets {
          lang
          langSlug
          code
      }
    }
  }
  `)

	request.Var("titleSlug", titleSlug)

	ctx := context.Background()

	var response Problem

	if err := client.Run(ctx, request, &response); err != nil {
		log.Fatal(err)
	}

	return response.Question
}

func getDailyProblem() DailyResponse {
	request := graphql.NewRequest(`
    query getDailyProblem {
    activeDailyCodingChallengeQuestion {
        link
        question {
            questionId
            questionFrontendId
            title
            titleSlug
            content
            translatedTitle
            contributors {
                username
                profileUrl
                avatarUrl
            }
            topicTags {
                name
                slug
                translatedName
            }
            companyTagStats
            codeSnippets {
                lang
                langSlug
                code
            }
            status
            sampleTestCase
        }
    }
}
  `)

	ctx := context.Background()

	var response DailyResponse

	if err := client.Run(ctx, request, &response); err != nil {
		log.Fatal(err)
	}

	return response

}

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

var (
	cfgFile string
	id      string
	rootCmd = &cobra.Command{
		Use:   "leetcode",
		Short: "Generate leetcode starter code",
	}
	dailyCmd = &cobra.Command{
		Use:   "daily",
		Short: "Get leetcode daily",
		Run: func(cmd *cobra.Command, args []string) {
			println("Generating leetcode daily...")
			result := getDailyProblem()

			generateFiles(result.ActiveDailyCodingChallengeQuestion.Question)

		},
	}
	problemCmd = &cobra.Command{
		Use:   "problem",
		Short: "Get leetcode daily",
		Run: func(cmd *cobra.Command, args []string) {
			println("Generating leetcode problem...", id)
			question := getProblemByTitleSlug(id)

			generateFiles(question)
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	rootCmd.AddCommand(dailyCmd)
	rootCmd.AddCommand(problemCmd)

	problemCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "problem id")
	problemCmd.MarkPersistentFlagRequired("id")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".leetcode")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	rootCmd.Execute()
}

func main() {
	Execute()
}
