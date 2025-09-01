package service

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"quiz.com/quiz/internal/collection"
	"quiz.com/quiz/internal/entity"
)

type QuizService struct {
	quizCollection *collection.QuizCollection
}


func (s QuizService) CreateQuiz(name string, questions []entity.QuizQuestion) error {
	quiz := entity.Quiz{
		Id:        primitive.NewObjectID(),
		Name:      name,
		Questions: questions,
	}
	return s.quizCollection.InsertQuiz(quiz)
}
func (s QuizService) DeleteQuizById(id primitive.ObjectID) error {
	return s.quizCollection.DeleteQuizById(id)
}


func Quiz(quizCollection *collection.QuizCollection) *QuizService {
	return &QuizService{
		quizCollection: quizCollection,
	}
}



func (s QuizService) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	return s.quizCollection.GetQuizById(id)
}

func (s QuizService) UpdateQuiz(id primitive.ObjectID, name string, questions []entity.QuizQuestion) error {
	quiz, err := s.quizCollection.GetQuizById(id)
	if err != nil {
		return err
	}

	if quiz == nil {
		return errors.New("quiz not found")
	}

	quiz.Name = name
	quiz.Questions = questions
	return s.quizCollection.UpdateQuiz(*quiz)
}

func (s QuizService) GetQuizzes() ([]entity.Quiz, error) {
	return s.quizCollection.GetQuizzes()
}
