package main

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
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

func GetProblemByTitleSlug(titleSlug string) Question {
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

func GetDailyProblem() DailyResponse {
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



