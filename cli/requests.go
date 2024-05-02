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
	Content            string         `json:"content"`
	SampleTestCase     string         `json:"sampleTestCase"`
}

type Problem struct {
	Question Question `json:"question"`
}

type DailyResponse struct {
	ActiveDailyCodingChallengeQuestion struct {
		Question Question `json:"question"`
	} `json:"activeDailyCodingChallengeQuestion"`
}

type ProblemsList struct {
	ProblemsetQuestionList struct {
			Questions []Question `json:"questions"`
	} `json:"problemsetQuestionList"`
}

func GetProblemBySearchInput(input string) Question {
	request := graphql.NewRequest(`
  query problemsetQuestionList($categorySlug: String!, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
    problemsetQuestionList: questionList(categorySlug: $categorySlug, limit: $limit, skip: $skip, filters: $filters) {
      questions: data {
          questionId
          questionFrontendId
          titleSlug
          content
          sampleTestCase
          codeSnippets {
              lang
              langSlug
              code
          }
        }
      }
  }
  `)

	filters := make(map[string]string)
	filters["searchKeywords"] = input

	request.Var("categorySlug", "all-code-essentials")
	request.Var("limit", 1)
	request.Var("skip", 0)
	request.Var("filters", filters)

	ctx := context.Background()

	var response ProblemsList

	if err := client.Run(ctx, request, &response); err != nil {
		log.Fatal(err)
	}

	if len(response.ProblemsetQuestionList.Questions) > 0 {
		return response.ProblemsetQuestionList.Questions[0]
	}

	return Question{}
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
            sampleTestCase
            codeSnippets {
                lang
                langSlug
                code
            }
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
