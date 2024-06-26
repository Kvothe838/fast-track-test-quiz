package memory

import (
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	"github.com/Kvothe838/fast-track-test-quiz/internal/services"
)

type repository struct {
	quiz                         models.Quiz
	currentSelectionByQuestionID map[int]models.ChoiceSelection
	quizSubmissions              []models.QuizSubmission
	submitted                    bool
}

func NewRepository() *repository {
	quiz := getFakeQuiz()
	return &repository{
		quiz:                         quiz,
		currentSelectionByQuestionID: make(map[int]models.ChoiceSelection),
		quizSubmissions:              getFakeQuizSubmissions(quiz.Questions),
	}
}

func getFakeQuiz() models.Quiz {
	return models.Quiz{
		Title:       "Fast Track Quiz",
		Description: "How much do you know about Fast Track? Let's see!",
		Questions: []models.Question{
			{
				ID:          1,
				Description: "Which industry is Fast Track revolutionising?",
				Choices: []models.Choice{
					{
						ID:          1,
						Description: "iGaming",
						IsCorrect:   true,
					},
					{
						ID:          2,
						Description: "iPhone",
						IsCorrect:   false,
					},
					{
						ID:          3,
						Description: "Board game",
						IsCorrect:   false,
					},
					{
						ID:          4,
						Description: "Spacecraft",
						IsCorrect:   false,
					},
				},
			},
			{
				ID:          2,
				Description: "Which are three Fast Track's promises?",
				Choices: []models.Choice{
					{
						ID:          1,
						Description: "Don't Settle, Work Smarter, Be Transparent",
						IsCorrect:   true,
					},
					{
						ID:          2,
						Description: "Care Deeply, Work Foolishly, Be Cutting Edge",
						IsCorrect:   false,
					},
					{
						ID:          3,
						Description: "Be Fearless, Hide Information, Don't Settle",
						IsCorrect:   false,
					},
					{
						ID:          4,
						Description: "Work Smarter, Be Coward, Don't Care",
						IsCorrect:   false,
					},
				},
			},
			{
				ID:          3,
				Description: "How can Fast Track help you?",
				Choices: []models.Choice{
					{
						ID:          1,
						Description: "By accessing consolidated, actionable data in one central location, enabling faster decision-making",
						IsCorrect:   true,
					},
					{
						ID:          2,
						Description: "By maximizing your efforts with customizable skateboards, covering race performance",
						IsCorrect:   false,
					},
					{
						ID:          3,
						Description: "By connecting cars and players with a single map, streamlining the race process",
						IsCorrect:   false,
					},
					{
						ID:          4,
						Description: "It can't help me, my games cannot be more powerful",
						IsCorrect:   false,
					},
				},
			},
			{
				ID:          4,
				Description: "Can BetConstruct be integrated to your CRM using Fast Track?",
				Choices: []models.Choice{
					{
						ID:          1,
						Description: "Yes",
						IsCorrect:   true,
					},
					{
						ID:          2,
						Description: "No",
						IsCorrect:   false,
					},
				},
			},
			{
				ID:          5,
				Description: "What is the name of comprehensive Fast Track's intelligence hub focused on data-driven innovation?",
				Choices: []models.Choice{
					{
						ID:          1,
						Description: "The Singularity Project",
						IsCorrect:   true,
					},
					{
						ID:          2,
						Description: "The Wild Project",
						IsCorrect:   false,
					},
					{
						ID:          3,
						Description: "The Singularity of Black Holes",
						IsCorrect:   false,
					},
					{
						ID:          4,
						Description: "Intelligence Hub",
						IsCorrect:   false,
					},
					{
						ID:          5,
						Description: "The Amazing Project-Man",
						IsCorrect:   false,
					},
				},
			},
		},
	}
}

func getFakeQuizSubmissions(questions []models.Question) []models.QuizSubmission {
	submissions := []models.QuizSubmission{
		{
			ID: 1,
			Selections: []models.ChoiceSelection{
				{
					QuestionID: 1,
					ChoiceID:   2,
				},
				{
					QuestionID: 2,
					ChoiceID:   2,
				},

				{
					QuestionID: 3,
					ChoiceID:   3,
				},
				{
					QuestionID: 4,
					ChoiceID:   1,
				},
				{
					QuestionID: 5,
					ChoiceID:   5,
				},
			},
		},
		{
			ID: 2,
			Selections: []models.ChoiceSelection{
				{
					QuestionID: 1,
					ChoiceID:   3,
				},
				{
					QuestionID: 2,
					ChoiceID:   1,
				},

				{
					QuestionID: 3,
					ChoiceID:   1,
				},
				{
					QuestionID: 4,
					ChoiceID:   2,
				},
				{
					QuestionID: 5,
					ChoiceID:   4,
				},
			},
		},
		{
			ID: 3,
			Selections: []models.ChoiceSelection{
				{
					QuestionID: 1,
					ChoiceID:   2,
				},
				{
					QuestionID: 2,
					ChoiceID:   1,
				},

				{
					QuestionID: 3,
					ChoiceID:   2,
				},
				{
					QuestionID: 4,
					ChoiceID:   1,
				},
				{
					QuestionID: 5,
					ChoiceID:   1,
				},
			},
		},
		{
			ID: 4,
			Selections: []models.ChoiceSelection{
				{
					QuestionID: 1,
					ChoiceID:   4,
				},
				{
					QuestionID: 2,
					ChoiceID:   4,
				},

				{
					QuestionID: 3,
					ChoiceID:   4,
				},
				{
					QuestionID: 4,
					ChoiceID:   2,
				},
				{
					QuestionID: 5,
					ChoiceID:   5,
				},
			},
		},
	}

	for index, submission := range submissions {
		submissions[index].HitsAmount = services.CalcHits(questions, submission.Selections)
	}

	return submissions
}
